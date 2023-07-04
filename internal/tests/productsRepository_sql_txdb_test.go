package repositoryutil_test

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/JainyMartins/goweb/internal/repository"
	"github.com/JainyMartins/goweb/internal/repository/repositoryutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	t.Run("deve salvar e testar", func(t *testing.T) {
		db := InitDatabase(t)
		defer db.Close()

		repository := repository.NewMySqlRepository(db)
	
		invalidProductId := 2
		product := repositoryutil.Product{
			Name: "batata",
			Category: "verduras",
			Count: 10,
			Price: 4.0,
		}
	
		result, err := repository.Salvar(product)
		log.Println(result)
		assert.NoError(t, err)

		getResult, err := repository.Get(invalidProductId)
		assert.NoError(t, err)
		assert.Equal(t, repositoryutil.Product{}, getResult)

		getResult, err = repository.Get(result.ID)
		assert.NoError(t, err)
		assert.NotNil(t, getResult)
		assert.Equal(t, product.Name, getResult.Name)
	})
}

func TestGetOneWithContext(t *testing.T){
	db := InitDatabase(t)

	id := 9
	product := repositoryutil.Product {
		Name: "",
	}

	myRepo := repository.NewMySqlRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//Caso quisermos simular uma demora na execucao de instrucao do banco, usamos o time.Sleep
	productResult, err := myRepo.GetOneWithContext(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productResult.Name)
}

func InitDatabase(t *testing.T) *sql.DB{
	t.Helper()
	txdb.Register("txdb", "mysql", "root:@/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	return db
}