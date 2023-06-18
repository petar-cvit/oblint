package internal

import (
	"context"
	"encoding/json"
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
)

func NewStorage() (Storage, error) {
	host := os.Getenv("REDIS_HOST")

	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return Storage{}, err
	}

	if err := client.FlushDB(context.Background()).Err(); err != nil {
		return Storage{}, err
	}

	return Storage{
		client: client,
	}, nil
}

func (s Storage) Random() error {
	random := rand.Int()

	return s.client.Set(context.Background(), strconv.Itoa(random), random, 0).Err()
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
