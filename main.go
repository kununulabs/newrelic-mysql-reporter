package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kununulabs/newrelic-mysql-reporter/mysql"
	"github.com/kununulabs/newrelic-mysql-reporter/yaml"
	"github.com/newrelic/newrelic-telemetry-sdk-go/telemetry"
)

// GetURL returns the insights api url based on account id and region
func GetURL(region, account string) string {
	return strings.ToLower(fmt.Sprintf(
		"https://insights-collector.%s01.nr-data.net/v1/accounts/%s/events",
		region,
		account,
	))
}

func main() {
	config, err := yaml.New(os.Args[1])
	if err != nil {
		panic(err)
	}

	if len(config.Metrics) < 1 {
		panic("No metrics provided. Aborting.")
	}

	mysqlConnection, err := mysql.GetConnection(
		os.Getenv("DATABASE_URL"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
	)
	if err != nil {
		panic(err)
	}

	defer mysqlConnection.Close()

	harvester, err := telemetry.NewHarvester(
		telemetry.ConfigAPIKey(os.Getenv("NEW_RELIC_INSIGHTS_INSERT_KEY")),
		telemetry.ConfigBasicAuditLogger(os.Stdout),
		telemetry.ConfigBasicDebugLogger(os.Stdout),
		telemetry.ConfigBasicErrorLogger(os.Stdout),
		telemetry.ConfigEventsURLOverride(GetURL(
			os.Getenv("NEW_RELIC_REGION"),
			os.Getenv("NEW_RELIC_ACCOUNT_ID"),
		)),
		telemetry.ConfigHarvestPeriod(0),
	)
	if err != nil {
		panic(err)
	}

	for _, metric := range config.Metrics {
		result := 0

		if err := mysqlConnection.QueryRow(metric.Query).Scan(&result); err != nil {
			log.Printf("%s: %s\n", metric.Name, err.Error())
			continue
		}

		log.Printf("%s: %d\n", metric.Name, result)

		config.Attributes["value"] = result

		harvester.RecordEvent(telemetry.Event{
			EventType:  metric.Name,
			Timestamp:  time.Now(),
			Attributes: config.Attributes,
		})
	}

	harvester.HarvestNow(context.Background())
}
