package repository

import (
	"bufio"
	"os"
)

type WordRepository struct {
	FilePath string
}

type CurrentWord struct {
	Date string `json:"date"`
	Word string `json:"word"`
}

// Load all words from words.txt
func (r *WordRepository) LoadWords() ([]string, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func (r *WordRepository) SaveWords(words []string) error {
	file, err := os.Create(r.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, w := range words {
		_, err := file.WriteString(w + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
