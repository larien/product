
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
