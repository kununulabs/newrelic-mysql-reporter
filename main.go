package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	newrelic "github.com/newrelic/go-insights/client"
)

const DEFAULT_CONFIG_FILE = "/config.yaml"

type Metric struct {
	Comment string `yaml:"comment"`
	Name    string `yaml:"name"`
	Query   string `yaml:"query"`
}

type Config struct {
	Metrics []Metric `yaml:"metrics"`
}

type Event struct {
	Type  string `json:"eventType"`
	Value int    `json:"eventValue"`
}

func main() {
	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) < 1 {
		configFile = DEFAULT_CONFIG_FILE
	}

	configYaml, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	config := Config{}

	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		panic(err)
	}

	mysqlURL := os.Getenv("DATABASE_URL")
	mysqlUsername := os.Getenv("DATABASE_USERNAME")
	mysqlPassword := os.Getenv("DATABASE_PASSWORD")

	if len(mysqlUsername) > 0 || len(mysqlPassword) > 0 {
		mysqlURL = fmt.Sprintf("%s:%s@%s", mysqlUsername, mysqlPassword, mysqlURL)
	}

	mysqlConn, err := sql.Open("mysql", mysqlURL)
	if err != nil {
		panic(err)
	}

	defer mysqlConn.Close()

	err = mysqlConn.Ping()
	if err != nil {
		panic(err)
	}

	nrClient := newrelic.NewInsertClient(os.Getenv("NR_API_KEY"), os.Getenv("NR_ACCOUNT_ID"))

	nrCustomURL := os.Getenv("NR_API_URL")
	if len(nrCustomURL) > 0 {
		nrClient.UseCustomURL(nrCustomURL)
	}

	// TODO: https://github.com/newrelic/go-insights/issues/16
	//err = nrClient.Validate()
	//if err != nil {
	//	panic(err)
	//}

	// make sure everything is sent before the program exits
	defer nrClient.Flush()

	// TODO: Don't panic in this loop, give other metrics a chance to check?
	for _, metric := range config.Metrics {
		event := Event{Type: metric.Name}

		err = mysqlConn.QueryRow(metric.Query).Scan(&event.Value)
		if err != nil {
			panic(err)
		}

		err = nrClient.PostEvent(event)
		if err != nil {
			panic(err)
		}
	}
}
