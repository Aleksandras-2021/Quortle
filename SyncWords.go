package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"Quortle/internal/models"
)

func syncWordsToDB(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", filePath, err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read %s: %v", filePath, err)
	}

	var dbCount int64
	if err := DB.Model(&models.Word{}).Count(&dbCount).Error; err != nil {
		log.Fatalf("Failed to count words in DB: %v", err)
	}

	if dbCount != int64(len(words)) {
		log.Printf("Syncing words to DB (%d words)...", len(words))
		if err := DB.Exec("TRUNCATE TABLE quortle.words").Error; err != nil {
			log.Fatalf("Failed to truncate words table: %v", err)
		}

		var wordRecords []models.Word
		for _, w := range words {
			wordRecords = append(wordRecords, models.Word{Word: w})
		}

		if err := DB.Create(&wordRecords).Error; err != nil {
			log.Fatalf("Failed to insert words into DB: %v", err)
		}
		log.Println("Words synced successfully!")
	} else {
		log.Println("Words table already up to date.")
	}
}
