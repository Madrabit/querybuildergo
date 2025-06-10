package tests

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/manager"
	"testing"
)

func TestGetDailyReport(t *testing.T) {
	ctx := context.Background()
	cnf, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	db, err := database.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	store := manager.NewStore(db)
	report, err := store.GetDailyReport(ctx, "Бартенева", "2024-11-08", "2024-11-09")
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(report), 0)
	for _, r := range report {
		fmt.Println(r)
	}
}
