package main
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap() error{
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil{
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil{
		return err
	}

	var 
}