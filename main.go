package main

import (
	"github.com/MondaleFelix/Jobbot/feeds"
	"github.com/MondaleFelix/Jobbot/feeds/indeed"
)

// "github.com/joho/godotenv"

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func createFeed(feedName string) feeds.PublicFeed {
	switch feedName {
	case "indeed":
		return indeed.NewPublicFeed(feedName)
	}
	return nil
}

func parseData() {
	go createFeed("indeed").Connect()
}

func broadcastData() {

}

func main() {
	// Parses specific site or all sites
	// Start each parsing of site in the separate goroutune
	// Run broadcast parsed jobs from site

	parseData()
	broadcastData()

}
