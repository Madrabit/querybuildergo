package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/manager"
	"testing"
)

func TestGetDailyReport(t *testing.T) {
	cnf, err := config.Load()
	require.NoError(t, err)
	db, err := database.NewMssqlStorage(cnf.DB)
	require.NoError(t, err)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal("error closing db")
		}
	}()
	store := manager.NewStore(db)
	tx, err := store.BeginTransaction()
	require.NoError(t, err)
	report, err := store.GetDailyReport(tx, "Бартенева", "2024-11-08", "2024-11-09")
	require.NoError(t, err)
	require.Greater(t, len(report), 0)
	for _, r := range report {
		fmt.Println(r)
	}
}
