package manager

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"querybuilder/internal/config"
	"querybuilder/internal/storage"
	"testing"
)

func TestMain(m *testing.M) {
	// Загрузим .env перед всеми тестами
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(".env not found or failed to load")
	}
	os.Exit(m.Run())
}

func TestGetDailyReport(t *testing.T) {
	ctx := context.Background()
	cnf, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	db, err := storage.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := NewStore(db)
	report, err := store.GetDailyReport(ctx, "Бартенева", "2024-11-08", "2024-11-09")
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(report), 0)
	for _, r := range report {
		fmt.Println(r)
	}
}
