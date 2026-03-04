package main

import (
	"encoding/json"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	var keys []string
	for key, _ := range pages {

		keys = append(keys, key)
	}
	sort.Strings(keys)

	var values []PageData
	for _, value := range keys {
		values = append(values, pages[value])
	}

	data, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		return err
	}

	os.WriteFile(filename, data, 0644)
	return nil
}
