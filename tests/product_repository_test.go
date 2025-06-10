package tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/product"
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	ctx := context.Background()
	cnf, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	db, err := database.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := product.NewStore(db)
	products, err := store.GetAllProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(products), 0)
	for _, product := range products {
		fmt.Println(product)
	}
}
