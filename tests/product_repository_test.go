package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/product"
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	cnf, err := config.Load()
	require.NoError(t, err)
	db, err := database.NewMssqlStorage(cnf.DB)
	require.NoError(t, err)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("error closing db")
		}
	}(db)
	require.NoError(t, err)
	store := product.NewStore(db)
	tx, err := store.BeginTransaction()
	require.NoError(t, err)
	products, err := store.GetAllProductsTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(products), 0)
	for _, p := range products {
		fmt.Println(p)
	}
}
