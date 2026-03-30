package services

import (
	"fmt"
	"math/rand"
	"time"

	"Quortle/internal/repository"
)

type WordService struct {
	repo *repository.WordRepository
}

func NewWordService(r *repository.WordRepository) *WordService {
	rand.Seed(time.Now().UnixNano())
	return &WordService{repo: r}
}

// Get the word for today
func (s *WordService) GetWordOfTheDay() (string, error) {
	words, err := s.repo.LoadWords()
	if err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", fmt.Errorf("no words available")
	}

	days := int(time.Now().Unix() / 86400)
	word := words[days%len(words)]

	return word, nil
}
