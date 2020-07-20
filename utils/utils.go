package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client *redis.Client

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Person is DB schema
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

//Initiate starts cache
func Initiate() {
	var (
		host     = getEnv("REDIS_HOST", "localhost")
		port     = string(getEnv("REDIS_PORT", "6379"))
		password = getEnv("REDIS_PASSWORD", "")
	)

	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	output, err := client.Ping().Result()
	fmt.Println(output)
	if err != nil {
		log.Fatal(err)
	}
}

//StoreInCache stores in cache
func StoreInCache(key string, data *Person) {
	b, _ := json.Marshal(&data)
	client.Set(key, string(b), 0)
}

//FetchFromCache get data from cache
func FetchFromCache(key string, offset int, limit int) []Person {
	data := client.Keys(key)
	var people []Person
	entries, _ := data.Result()
	for index, id := range entries {
		if index >= offset && index < (offset+limit) {
			data := client.Get(id)
			entry, _ := data.Result()
			person := new(Person)
			json.Unmarshal([]byte(entry), person)
			people = append(people, *person)
		}
	}
	return people
}
