package main

import (
    "fmt"
    "log"
    "net/url"
    "os"
    "github.com/ChimeraCoder/anaconda"
    "github.com/joho/godotenv"
)

func loadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func main() {
	loadEnv()

	api := getTwitterApi()

	v := url.Values{}
	v.Set("count", "30")

	searchResult, _ := api.GetSearch("#Cityzens開発記録", v)
	for _, tweet := range searchResult.Statuses {
			fmt.Println(tweet.Text)
	}
}