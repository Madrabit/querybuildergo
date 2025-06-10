package tests

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/employee"
	"testing"
)

func TestGetEmplByProducts(t *testing.T) {
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
	store := employee.NewStore(db)
	products := []string{"IRB-моделирование для профессионалов"}
	empl, err := store.FindByProducts(ctx, products)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(empl), 0)
	for _, e := range empl {
		fmt.Println(e)
	}
}
