package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Thuvii/aggregator/internal/database"
	"github.com/google/uuid"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

// Login
func handlerLogin(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		return fmt.Errorf("Login handler expects a single argument, the username.")
	}

	username := cmd.Args[0]
	if _, err := s.db.GetUser(context.Background(), username); err != nil {
		fmt.Println("This user does not exist")
		os.Exit(1)
	}
	err := s.stateConfig.SetUser(username)
	if err != nil {
		return fmt.Errorf("Cannot set usenname, error: %v", err)
	}
	fmt.Println("The user has been set.")
	return nil
}

// Register
func handlerRegister(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		return fmt.Errorf("Login handler expects a single argument, the username.")
	}
	username := cmd.Args[0]
	if _, err := s.db.GetUser(context.Background(), username); err == nil {
		fmt.Println("This user already exist")
		os.Exit(1)
	}
	if user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}); err != nil {
		return fmt.Errorf("Can't create user. Error: %v", err)
	} else {
		fmt.Println("The user has been register.")
		printUser(user)
	}

	err := s.stateConfig.SetUser(username)
	if err != nil {
		return fmt.Errorf("Cannot set usenname, error: %v", err)
	}
	return nil
}

// handle reset database
func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("Can not reset database. Error: %v", err)
	}
	fmt.Println("Reset database successful!")
	return nil
}

// handler print all users
func handlerGet(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if s.stateConfig.CurrentUserName == user.Name {
			fmt.Printf(" * Name:    %v (current)\n", user.Name)
		} else {

			fmt.Printf(" * Name:    %v\n", user.Name)
		}
	}
	return nil
}
