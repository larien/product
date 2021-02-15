package product_test

import (
	"log"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/larien/product/product/drivers/database"
	"github.com/larien/product/product/drivers/database/tests"
	"github.com/larien/product/product/entity"
	productRepository "github.com/larien/product/product/repository/product"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func (suite *ProductTestSuite) TestGet_Success() {
	s := productRepository.New(suite.DB.Connection)
	product := &entity.Product{
		ID:           faker.UUIDDigit(),
		PriceInCents: faker.UnixTime(),
		Title:        faker.Name(),
		Description:  faker.Sentence(),
	}
	suite.is.Nil(s.Create(product))
	obtainedProduct, err := s.Get(product.ID)
	suite.is.NoError(err)
	suite.is.Equal(product.PriceInCents, obtainedProduct.PriceInCents)
	suite.is.Equal(product.Title, obtainedProduct.Title)
	suite.is.Equal(product.Description, obtainedProduct.Description)
}

func (suite *ProductTestSuite) TestGet_Success_NotFound() {
	s := productRepository.New(suite.DB.Connection)
	product, err := s.Get(uuid.New().String())
	suite.is.Nil(product)
	suite.is.NoError(err)
}

func (suite *ProductTestSuite) TestGet_Failure() {
	suite.is.Nil(suite.DB.Connection.Close())
	s := productRepository.New(suite.DB.Connection)
	product, err := s.Get(uuid.New().String())
	suite.is.Nil(product)
	suite.is.EqualError(err, database.ErrClosed.Error())
}

type ProductTestSuite struct {
	suite.Suite
	is require.Assertions
	DB *tests.Database
}

func (suite *ProductTestSuite) SetupSuite() {
	suite.DB = tests.NewDB(suite.T())
}

func (suite *ProductTestSuite) SetupTest() {
	suite.is.Nil(suite.Up())
}

func (suite *ProductTestSuite) TearDownTest() {
	suite.is.Nil(suite.Down())
}

// TestSuite won't run if -short flag is provided
func TestSuite(t *testing.T) {
	if testing.Short() {
		return
	}
	suite.Run(t, new(ProductTestSuite))
}

// Up opens a new connection if it's not open and creates a new table for this layer.
func (suite *ProductTestSuite) Up() error {
	// check if connection is open
	if err := suite.DB.Connection.Ping(); err != nil {
		// restablish the connection if it's closed
		db, err := tests.DefaultConnection(suite.DB.Port)
		suite.is.Nil(err)
		suite.DB.Connection = db
	}

	schema := `
	CREATE TABLE IF NOT EXISTS products
	(
		id uuid PRIMARY KEY,
		price_in_cents INTEGER NOT NULL,
		title VARCHAR(100) NOT NULL,
		description VARCHAR(500),
		created_at TIMESTAMP DEFAULT(now()),
		updated_at TIMESTAMP DEFAULT(now()),
		deleted_at TIMESTAMP
	);`

	_, err := suite.DB.Connection.Exec(schema)
	if err != nil {
		log.Fatalf("%v: couldn't create table: %v", suite.Suite.T().Name(), err)
	}
	return nil
}

// Down checks if the database connection is open before dropping the tables
// related to this layer. If the connection is already closed, no dropping is done.
func (suite *ProductTestSuite) Down() error {
	// checks if the connectio is open
	if err := suite.DB.Connection.Ping(); err != nil {
		// if it's not open, don't drop anything
		return nil
	}
	_, err := suite.DB.Connection.Exec("DROP TABLE products")
	if err != nil {
		log.Fatalf("%v: couldn't drop table: %v", suite.Suite.T().Name(), err)
	}
	return nil
}
