package modules

import (
	"base_app_go/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zebresel-com/mongodm"
)

func MongodbConn() *mongodm.Connection {
	file, err := ioutil.ReadFile("storages/files/locals.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal JSON to map
	var localMap map[string]map[string]string
	json.Unmarshal(file, &localMap)
	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{config.App.AppConfig.String("host")},
		DatabaseName:  config.App.AppConfig.String("database"),
		Locals:        localMap["en-US"],
	}
	// Connect to database
	db, err := mongodm.Connect(dbConfig)

	// Check for error
	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}

	return db
}
