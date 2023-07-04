package repositoryutil_test

// import (
// 	"database/sql"
// 	"testing"

// 	// "github.com/DATA-DOG/go-sqlmock"
// 	// "github.com/JainyMartins/goweb/internal/repository"
// 	// "github.com/JainyMartins/goweb/internal/repository/repositoryutil"
// 	"github.com/stretchr/testify/assert"
// )

//Vai falar depois como fazer
// func Test_sql_Repository_Salvar_Mock(t *testing.T) {
// 	t.Run("deve buscar um produto com o id informado", func(t *testing.T) {
// 		db, mock := SetupMock(t)
// 		defer db.Close()
// 		productId := 1

// 		//Falamos quais são as colunas
// 		columns := []string{"id", "name", "type", "count", "price"}
// 		rows := sqlmock.NewRows(columns)
// 		rows.AddRow(productId, "", "", "", "") //add linha como se fosse resultado

// 		//Falamos o que vamos passar e o que esperamos
// 		mock.ExpectExec("INSERT INTO products").WithArgs(sql.Named("id", productId),sql.Named("name", "batata")).WillReturnResult(sqlmock.NewResult(1, 1))

// 		mock.ExpectQuery("SELECT .* FROM products").WithArgs(sql.Named("id", productId),sql.Named("name", "batata")).WillReturnRows(rows) //interagindo com tabela

// 		//Criamos nosso repository com o db criado no MOCK
// 		repository := repository.NewMySqlRepository(db)

// 		//Produto que iremos inserir primeiro para depois buscar
// 		product := repositoryutil.Product{
// 			ID: productId,
// 			Name: "batata",
// 		}

// 		//Verificamos se não há produtos na base de dados com esse id, fazemos asserções de não ter erros na interação e verificamos se retorno é nil
// 		getResult, err := repository.Get(productId)
// 		assert.NoError(t, err)
// 		assert.Nil(t, getResult)

// 		//Aqui estamos inserindo o produto
// 		_, err = repository.Salvar(product)
// 		assert.NoError(t, err)

// 		//Aqui buscamos o produto que inserimos
// 		getResult, err = repository.Get(productId)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, getResult)
// 		assert.Equal(t, product.ID, getResult.ID)
// 		assert.NoError(t, mock.ExpectationsWereMet())
// 	})

// }

// func SetupMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
// 	t.Helper()
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	return db, mock
// }

