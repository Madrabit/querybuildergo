package employee

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"log"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"testing"
)

func TestCreateExl(t *testing.T) {
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
	store := NewStore(db)
	products := []string{"IRB-моделирование для профессионалов"}
	empl, err := store.FindByProducts(ctx, products)
	if err != nil {
		log.Fatal(err)
	}
	require.Greater(t, len(empl), 0)
	err = CreateExl(empl)
	require.NoError(t, err)
}
