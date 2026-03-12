package main

import (
	"fmt"

	"github.com/Thuvii/aggregator/internal/configure"
	"github.com/Thuvii/aggregator/internal/database"
)

type state struct {
	db          *database.Queries
	stateConfig *configure.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handler map[string](func(*state, command) error)
}

// command handler functions
func (c *commands) run(s *state, cmd command) error {
	value, exist := c.handler[cmd.Name]
	if !exist {
		return fmt.Errorf("Command does not have handler function")
	}
	return value(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handler[name] = f
}
