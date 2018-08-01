package main

import (
	"github.com/sausozluk/sozluk-etl/helpers"
	"github.com/sausozluk/sozluk-etl/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
	"fmt"
)

type SozlukData struct {
	entries []models.Entry
	topics  []models.Topic
}

type Config struct {
	Server   string `yaml:"server"`
	Database string `yaml:"database"`
}

func main() {
	config := GetConfig()
	s := helpers.SozlukDB{Server: config.Server, Database: config.Database}
	db := s.InitMongo()
	sozlukData := SozlukData{}
	err := FetchFromMongoDB(db, &sozlukData)

	if err != nil {
		log.Fatal(err)
	} else {
		GoParallel(&sozlukData)
	}
}

func GoParallel(sozlukData *SozlukData) {
	entryLen := len((*sozlukData).entries)

	var wg sync.WaitGroup
	wg.Add(entryLen)

	for i := 0; i < entryLen; i++ {
		go ProcessData(i, sozlukData, &wg)
	}

	wg.Wait()
}

func ProcessData(index int, sozlukData *SozlukData, wg *sync.WaitGroup) {
	defer wg.Done()
	entry := sozlukData.entries[index]
	fmt.Printf("Index: %v, EntryID: %v\n", index, entry.EntryID)
}

func FetchFromMongoDB(db *mgo.Database, data *SozlukData) error {
	entries := db.C("entries")
	topics := db.C("topics")

	fetchEntriesErr := entries.Find(bson.M{}).All(&data.entries)
	fetchTopicsErr := topics.Find(bson.M{}).All(&data.topics)

	if fetchEntriesErr != nil {
		return fetchEntriesErr
	}

	if fetchTopicsErr != nil {
		return fetchTopicsErr
	}

	return nil
}

func GetConfig() Config {
	config := Config{}
	env := GetSozlukEnv()

	yamlFile, err := ioutil.ReadFile("configs/" + env + ".yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}

func GetSozlukEnv() string {
	env := os.Getenv("SOZLUK_ENV")

	if env == "" {
		env = "local"
	}

	return env
}
