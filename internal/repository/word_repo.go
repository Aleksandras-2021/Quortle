package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
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

// Save remaining words back to words.txt
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

// Save today's word
func (r *WordRepository) SaveCurrentWord(word string) error {
	data := CurrentWord{
		Date: time.Now().Format("2006-01-02"),
		Word: word,
	}

	file, err := os.Create("current_word.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

// Load today's word
func (r *WordRepository) LoadCurrentWord() (string, string, error) {
	file, err := os.Open("current_word.json")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var data CurrentWord
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return "", "", err
	}

	return data.Date, data.Word, nil
}
