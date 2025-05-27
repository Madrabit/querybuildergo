package employee

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/storage"
	"testing"
)

func TestCreateExl(t *testing.T) {
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
	products := []string{"IRB-моделирование для профессионалов"}
	empl, err := store.findByProducts(ctx, products)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(empl), 0)
	CreateExl(empl)
}
