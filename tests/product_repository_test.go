package tests

import (
	"fmt"
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
	defer db.Close()
	// Начинаем транзакцию

	require.NoError(t, err)
	// Откатим транзакцию в конце, чтобы не оставлять изменения
	store := product.NewStore(db)
	tx, err := store.BeginTransaction()
	products, err := store.GetAllProductsTx(tx)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(products), 0)
	for _, product := range products {
		fmt.Println(product)
	}
}
