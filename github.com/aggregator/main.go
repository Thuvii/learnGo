package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Thuvii/aggregator/internal/configure"
	"github.com/Thuvii/aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configure.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	stateMain := state{stateConfig: &config}
	//get and create sql from url
	db, err := sql.Open("postgres", stateMain.stateConfig.DbURL)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	dbQueries := database.New(db)

	stateMain.db = dbQueries

	//create and register new command handler
	commandsHandler := commands{handler: make(map[string]func(*state, command) error)}

	commandsHandler.register("login", handlerLogin)
	commandsHandler.register("register", handlerRegister)
	commandsHandler.register("reset", handlerReset)
	commandsHandler.register("users", handlerGet)
	commandsHandler.register("agg", handleAgg)
	commandsHandler.register("addfeed", middlewareLoggedIn(handleAddFeed))
	commandsHandler.register("feeds", handlListFeeds)
	commandsHandler.register("follow", middlewareLoggedIn(handleFollow))
	commandsHandler.register("following", middlewareLoggedIn(handleListFollows))
	commandsHandler.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commandsHandler.register("browse", middlewareLoggedIn(handlerBrowse))
	if len(os.Args) < 2 {
		fmt.Println("usage: gator <command> [args...]")
		os.Exit(1)
	}

	username := os.Args[2:]

	cmd := command{Name: os.Args[1],
		Args: username,
	}
	err = commandsHandler.run(&stateMain, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
