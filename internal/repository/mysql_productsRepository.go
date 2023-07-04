package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/JainyMartins/goweb/internal/repository/repositoryutil"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySqlRepository(dbConn *sql.DB) repositoryutil.Repository {
	return &mysqlRepository{
		db: dbConn,
	}
}

const (
	GetProduct         = "SELECT id, name, type, count, price FROM products WHERE id = ?"
	GetAllProducts     = "SELECT id, name, type, count, price FROM products"
	SaveProduct        = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	UpdateProduct      = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	UpdateProductName = "UPDATE products SET name = ? WHERE id = ?"
	UpdateProductPrice = "UPDATE products SET price = ? WHERE id = ?"
	DeleteProduct      = "DELETE FROM products WHERE id = ?"
	GetFullData = "SELECT products.id, products.name, products.count, products.type, products.price, warehouses.name, warehouses.adress " +
		"FROM products INNER JOIN warehouses ON products.id_warehouse = warehouses.id " +
		"WHERE products.id = ?"
)

//JDBC - ORM JAVA
//GORM - ORM DE GoLang - Ele coloca muita coisa para debaixo dos panos.
// main -> routes -> controller <--> service <--> repository <--> db
//sqlc é uma biblioteca que conseguimos colocar todos os models, schemas, queries, constantes

func (r *mysqlRepository) GetFullData(id int) repositoryutil.ProductFullDataResponse {
	var productFullData repositoryutil.ProductFullDataResponse
	rows, err := r.db.Query(GetFullData, id)
	if err != nil {
		log.Println(err)
		return productFullData
	}
	for rows.Next() {
		if err := rows.Scan(&productFullData.ID, &productFullData.Name, &productFullData.Count, &productFullData.Category, &productFullData.Price, &productFullData.Warehouse,
			&productFullData.WarehouseAdress); err != nil {
			log.Fatal(err)
			return productFullData
		}
	}
	return productFullData
}

func (r *mysqlRepository) Salvar(product repositoryutil.Product) (repositoryutil.Product, error) {
	// o banco é iniciado
	stmt, err := r.db.Prepare(SaveProduct) // monta o  SQL
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // a instrução fecha quando termina. Se eles permanecerem abertos, o consumo de memória é gerado

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Category, product.Count, product.Price) // retorna um sql.Result ou um error
	if err != nil {
		return repositoryutil.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // do sql.Result retornado na execução obtemos o Id inserido
	product.ID = int(insertedId)

	return product, nil
}

func (r *mysqlRepository) Get(id int) (repositoryutil.Product, error) {
	var product repositoryutil.Product

	rows, err := r.db.Query(GetProduct, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return product, err
		}
	}
	return product, nil
}

func (r *mysqlRepository) GetOneWithContext(ctx context.Context, id int) (repositoryutil.Product, error) {
	var product repositoryutil.Product
	//Caso quisermos simular timeout, podemos utilizar 
	//getQuery := "SELECT SLEEP(30) FROM DUAL where 0 < ?" ou time.Sleep(time.Second * 2)
	rows, err := r.db.QueryContext(ctx, GetProduct, id)
	if err != nil {
		log.Println(err)
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Count, &product.Price); err != nil {
			log.Println(err.Error())
			return product, err
		}
	}
	return product, nil
}

func (r *mysqlRepository) Update(product repositoryutil.Product) (repositoryutil.Product, error) {
	stmt, err := r.db.Prepare(UpdateProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Category, product.Count, product.Price, product.ID)
	if err != nil {
		return repositoryutil.Product{}, err
	}
	return product, nil
}

func (r *mysqlRepository) Delete(id int) error {
	stmt, err := r.db.Prepare(DeleteProduct)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *mysqlRepository) GetAll() ([]repositoryutil.Product, error) {
	var products []repositoryutil.Product

	rows, err := r.db.Query(GetAllProducts)

	if err != nil {
		log.Println(err)
		return products, err
	}

	for rows.Next() {
		// id, name, type, count, price
		var product repositoryutil.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Count, &product.Price)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *mysqlRepository) UpdateName(id int, name string) (repositoryutil.Product, error) {
	var product repositoryutil.Product

	stmt, err := r.db.Prepare(UpdateProductName)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	product, err = r.Get(id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	return product, nil
}

func (r *mysqlRepository) UpdatePrice(id int, price float64) (repositoryutil.Product, error) {
	var product repositoryutil.Product

	stmt, err := r.db.Prepare(UpdateProductPrice)	
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(price, id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	product, err = r.Get(id)
	if err != nil {
		return repositoryutil.Product{}, err
	}

	return product, nil
}
