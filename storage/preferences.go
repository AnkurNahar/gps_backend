package storage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	models "gps_backend/models"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr, password string, db int) (*RedisStorage, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis")
	return &RedisStorage{
		client: client,
		ctx:    ctx,
	}, nil
}

func (r *RedisStorage) GetPreferences(userID string) (models.Preferences, error) {
	val, err := r.client.Get(r.ctx, userID).Result()
	if err == redis.Nil {
		//default preferences
		return models.Preferences{
			SortBy:       		"display_name",
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

func (r *RedisStorage) SavePreferences(userID string, preferences models.Preferences) error {
	//log.Println(preferences)
	data, err := json.Marshal(preferences)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, userID, data, 0).Err()
}