package services

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"time"
	"url-shortener/db"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// generates a random short url
func GenerateShortURL(length int) string {
	// rand.Seed(time.Now().UnixNano())
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[rng.Intn(len(charset))]
	}
	return string(shortURL)
}

func GenerateAndStoreURL(originalURL string, customURL string) (string, error) {
	db := db.InitDB()
	defer db.Close() // closes connection after function ends

	var shortURL string

	if customURL != "" {
		shortURL = customURL
	} else {
		for {
			shortURL = GenerateShortURL(5)
			if !IsURLTaken(shortURL) {
				break
			}
		}
	}

	insertQuery := `
	INSERT INTO urls(original_url, short_url) 
	VALUES (?, ?)`

	_, err := db.Exec(insertQuery, originalURL, shortURL)

	if err != nil {
		log.Println("Error while writing to database:", err)
		return "", err
	}
	return shortURL, nil
}

func IsURLTaken(url string) bool {
	db := db.InitDB()
	defer db.Close()

	query := `
	SELECT EXISTS(SELECT 1 FROM urls WHERE short_url = ?)
	`
	var exists bool

	err := db.QueryRow(query, url).Scan(&exists)
	if err != nil {
		log.Println("Error while trying to check if URL is taken:", err)
	}
	return exists
}

func CheckURLExists(originalURL string) (string, error) {
	fmt.Println("inside CheckURLexists")
	db := db.InitDB()
	defer db.Close()

	var shortURL string

	getQuery := `
	SELECT short_url FROM urls WHERE original_url = ?
	`
	err := db.QueryRow(getQuery, originalURL).Scan(&shortURL)

	if err != nil && err != sql.ErrNoRows {
		log.Println("Error while checking for originalURL in database:", err)
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", err
	}

	return shortURL, nil
}

func GetOriginalURL(shortURL string) (string, error) {
	db := db.InitDB()
	defer db.Close()
	var originalURL string

	getQuery := `
	SELECT original_url FROM urls WHERE short_url = ?
	`
	err := db.QueryRow(getQuery, shortURL).Scan(&originalURL)

	if err != nil {
		log.Println("Error while trying to fetch original URL:", err)
		return "", err
	}
	return originalURL, nil
}

func CheckSpecialCharacter(url string) int8 {
	if len(url) > 10 {
		return 1
	}
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]*$`, url)
	if !match {
		return 2
	}
	return 3
}
