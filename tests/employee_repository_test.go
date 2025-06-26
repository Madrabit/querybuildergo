package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/employee"
	"testing"
)

func TestGetEmplByProducts(t *testing.T) {
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
	store := employee.NewStore(db)
	tx, err := store.BeginTransaction()
	require.NoError(t, err)
	products := []string{"IRB-моделирование для профессионалов"}
	empl, err := store.FindByProducts(tx, products)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(empl), 0)
	for _, e := range empl {
		fmt.Println(e)
	}
}
