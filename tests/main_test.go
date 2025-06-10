package tests

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Загрузим .env перед всеми тестами
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env not found or failed to load")
	}
	os.Exit(m.Run())
}
