package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr, password string) (*RedisClient, error) {
	opt := &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	}

	client := redis.NewClient(opt)

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}

func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}

func (rc *RedisClient) SetUser(user User) error {
	json, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return rc.Client.Set(context.Background(), user.Username, json, 0).Err()
}

func (rc *RedisClient) GetUser(username string) (*User, error) {
	val, err := rc.Client.Get(context.Background(), username).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	user := &User{}

	if err := json.Unmarshal([]byte(val), user); err != nil {
		return nil, err
	}

	return user, nil
}

func (rc *RedisClient) DeleteUser(username string) error {
	return rc.Client.Del(context.Background(), username).Err()
}
