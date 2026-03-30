package services

import (
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
	today := time.Now().Format("2006-01-02")

	// Load previous word
	date, word, err := s.repo.LoadCurrentWord()
	if err == nil && date == today {
		return word, nil
	}

	// Load all words
	words, err := s.repo.LoadWords()
	if err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "No more words left", nil // no more words left
	}

	idx := rand.Intn(len(words))
	word = words[idx]

	words = append(words[:idx], words[idx+1:]...)

	if err := s.repo.SaveWords(words); err != nil {
		return "", err
	}
	if err := s.repo.SaveCurrentWord(word); err != nil {
		return "", err
	}

	return word, nil
}
