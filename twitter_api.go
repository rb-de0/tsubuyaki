package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kardianos/osext"
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
	path, extErr := osext.Executable()

	if extErr != nil {
		return Config{}, errors.New("Error, could not find a excutable path.")
	}

	splitPath := strings.Split(path, "/")
	splitPath = splitPath[0 : len(splitPath)-1]

	excutablePath := ""
	for index := range splitPath {
		excutablePath += splitPath[index] + "/"
	}

	fp, err := os.Open(excutablePath + "config.json")

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
