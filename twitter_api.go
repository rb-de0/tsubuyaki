package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type Config struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func NewApi() *anaconda.TwitterApi {
	config, err := loadConfig()

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(-1)
	}

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessTokenSecret)
	return api
}

func loadConfig() (Config, error) {
	fp, err := os.Open("./config.json")

	if err != nil {
		fp.Close()
		return Config{}, errors.New("Error, could not find a config file.")
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanned := ""

	for scanner.Scan() {
		scanned += scanner.Text() + "\n"
	}

	var mapped Config
	mappingErr := json.Unmarshal([]byte(scanned), &mapped)

	if mappingErr != nil {
		return Config{}, errors.New("Error, could not find a config file.")
	}

	return mapped, nil
}
