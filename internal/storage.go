package internal

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"strconv"

	"example.com/oblint/internal/models"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	client *redis.Client
}

const (
	history = "history"
	ongoing = "ongoing"
	forum   = "forum"
)

func NewStorage() (Storage, error) {
	host := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return Storage{}, err
	}

	//if err := client.FlushAll(context.Background()).Err(); err != nil {
	//	return Storage{}, err
	//}

	return Storage{
		client: client,
	}, nil
}

func (s Storage) Random() error {
	random := rand.Int()

	return s.client.Set(context.Background(), strconv.Itoa(random), random, 0).Err()
}

func (s Storage) Clear() {
	if err := s.client.Del(context.Background(), history).Err(); err != nil {
		panic(err)
	}

	if err := s.client.Del(context.Background(), ongoing).Err(); err != nil {
		panic(err)
	}

	if err := s.client.Del(context.Background(), forum).Err(); err != nil {
		panic(err)
	}
}

func (s Storage) SaveToHistory(hw models.HistoryHomework) error {
	data, err := json.Marshal(hw)
	if err != nil {
		return err
	}

	return s.client.HSet(context.Background(), history, hw.ID, data).Err()
}

func (s Storage) GetHistory() ([]models.HistoryHomework, error) {
	res, err := s.client.HGetAll(context.Background(), history).Result()
	if err != nil {
		return nil, err
	}

	out := make([]models.HistoryHomework, 0)

	for _, hw := range res {
		var homework models.HistoryHomework

		if err := json.Unmarshal([]byte(hw), &homework); err != nil {
			return nil, err
		}

		out = append(out, homework)
	}

	return out, nil
}

func (s Storage) GetHistoryByID(ID string) (models.HistoryHomework, error) {
	res, err := s.client.HGet(context.Background(), history, ID).Result()
	if err != nil {
		return models.HistoryHomework{}, err
	}

	var homework models.HistoryHomework
	if err := json.Unmarshal([]byte(res), &homework); err != nil {
		return models.HistoryHomework{}, err
	}

	return homework, nil
}

func (s Storage) SaveToHomeworks(hw models.Homework) error {
	data, err := json.Marshal(hw)
	if err != nil {
		return err
	}

	return s.client.HSet(context.Background(), ongoing, hw.ID, data).Err()
}

func (s Storage) DeleteFromHomeworks(hw models.Homework) error {
	return s.client.HDel(context.Background(), ongoing, hw.ID).Err()
}

func (s Storage) GetHomeworks() ([]models.Homework, error) {
	res, err := s.client.HGetAll(context.Background(), ongoing).Result()
	if err != nil {
		return nil, err
	}

	out := make([]models.Homework, 0)

	for _, hw := range res {
		var homework models.Homework

		if err := json.Unmarshal([]byte(hw), &homework); err != nil {
			return nil, err
		}

		out = append(out, homework)
	}

	return out, nil
}

func (s Storage) GetHomeworkByID(ID string) (models.Homework, error) {
	res, err := s.client.HGet(context.Background(), ongoing, ID).Result()
	if err != nil {
		return models.Homework{}, err
	}

	var homework models.Homework
	if err := json.Unmarshal([]byte(res), &homework); err != nil {
		return models.Homework{}, err
	}

	return homework, nil
}

func (s Storage) GetForum() ([]models.Message, error) {
	data, err := s.client.LRange(context.Background(), forum, 0, -1).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return []models.Message{}, nil
		}
	}

	out := make([]models.Message, 0)

	for _, dataMsg := range data {
		var msg models.Message
		if err := json.Unmarshal([]byte(dataMsg), &msg); err != nil {
			return []models.Message{}, err
		}

		out = append(out, msg)
	}

	return out, nil
}

func (s Storage) AddMessage(msg models.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return s.client.LPush(context.Background(), forum, data).Err()
}
