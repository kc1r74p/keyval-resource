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
	versions := []models.InResponse{}
	versions = append(versions, models.InResponse{Version: "Init=1"})
	json.NewEncoder(os.Stdout).Encode(versions)
}
