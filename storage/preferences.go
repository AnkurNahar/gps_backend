package storage

import (
	"context"
	"encoding/json"
	"log"
	"github.com/go-redis/redis/v8"
	models "gps_backend/models"
)


var ctx = context.Background()
var redisClient *redis.Client

func InitRedis() { //redis connection
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

func GetPreferences(userID string) (models.Preferences, error) {
	val, err := redisClient.Get(ctx, userID).Result()
	if err == redis.Nil {
		//default preferences
		return models.Preferences{
			SortBy:       		"name",
			HiddenDeviceIds:	[]string{},
			UserDeviceIcons:	map[string]string{},
		}, nil
	} else if err != nil {
		return models.Preferences{}, err
	}

	var preferences models.Preferences
	err = json.Unmarshal([]byte(val), &preferences)
	return preferences, err
}

func SavePreferences(userID string, preferences models.Preferences) error {
	data, err := json.Marshal(preferences)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, userID, data, 0).Err()
}