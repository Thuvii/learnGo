package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Thuvii/aggregator/internal/database"
	"github.com/google/uuid"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Println("------------------------------")
	fmt.Printf(" * UserCreatedFeed:    %v\n", user.Name)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * Url:    %v\n", feed.Url)
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

// aggregator handler

func handleAgg(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		return fmt.Errorf("need a time duration: 1h, 1m, 1s, ...")
	}
	duration := cmd.Args[0]

	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeDuration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)

		go func() {
			end := time.Now().Add(timeDuration)
			for time.Now().Before(end) {
				remaining := time.Until(end).Round(time.Second)
				fmt.Printf("\rNext fetch in: %s   ", remaining)
				time.Sleep(time.Second)
			}
			fmt.Print("\r")
		}()
		fmt.Printf("Collecting feeds every %s\n", timeDuration)
	}
	return nil

}

// addfeed handler
func handleAddFeed(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) != 2 {
		return fmt.Errorf("Addfeed require 2 args!")
	}

	feedname := cmd.Args[0]
	feedurl := cmd.Args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:            uuid.New(),
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Valid: false},
		Name:          feedname,
		Url:           feedurl,
		UserID:        user.ID,
	})
	if err != nil {
		return err
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	printFeed(feed, user)
	fmt.Println("Following: ")
	fmt.Printf(" * FeedName:    %v\n", follow.FeedName)
	fmt.Println("------------------------------")
	return nil
}

//feeds handler

func handlListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("cannot get user by id, %v", err)
		}
		printFeed(feed, user)

	}
	return nil
}

// handler create/add follow

func handleFollow(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		return fmt.Errorf("need 1 feed url")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	fmt.Println("Following: ")
	fmt.Printf(" * Username:    %v\n", follow.UserName)
	fmt.Printf(" * FeedName:    %v\n", follow.FeedName)
	return nil
}

// handler list feed followed

func handleListFollows(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf(" * Username:    %v\n", user.Name)
	fmt.Println("------------------------------")
	for i, follow := range follows {
		fmt.Printf("%v. FeedName:    %v\n", i+1, follow.FeedName)
	}
	fmt.Println("------------------------------")
	return nil
}

// unfollow
func handlerUnfollow(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) > 1 {
		return fmt.Errorf("need 1 feed url")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	var err error
	limit := 0
	if cmd.Args == nil || len(cmd.Args) != 1 {
		limit = 2
	} else {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
	}

	data, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}
	for _, post := range data {
		fmt.Println(post.Title)
	}
	return nil
}
