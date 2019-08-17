package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/SWCE/keyval-resource/models"
)

func main() {
	var request models.CheckRequest
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err.Error())
		os.Exit(1)
	}

	var keys = map[string]string{}
	keys["Init"] = "1"
	//keys = append(keys, {"Init=1"})
	json.NewEncoder(os.Stdout).Encode(models.InResponse{
		Version:  keys,
	})

	//versions := []models.OutResponse{}
	//var keys []string
	//keys = append(keys, "Init=1")
	//versions = append(versions, models.OutResponse{Version: keys})
	//json.NewEncoder(os.Stdout).Encode(versions)
}
