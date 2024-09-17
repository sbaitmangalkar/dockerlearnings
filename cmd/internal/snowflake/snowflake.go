package snowflake

import "database/sql"

type Config struct {
	Account                 string
	Username                string
	Password                string
	PrivateKey              string
	Database                string
	Warehouse               string
	Schema                  string
	Role                    string
	PrivateKeySecretKeyName string
}

type SnowflakeClient interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type snowflakeDB struct {
	client *sql.DB
}

func NewSnowflakeClient(cfg *Config) (SnowflakeClient, error) {
	db, err := NewClient(cfg).GetDb()
	if err != nil {
		return nil, err
	}

	return &snowflakeDB{client: db}, nil
}

func (c *snowflakeDB) Query(query string, args ...any) (*sql.Rows, error) {
	return c.client.Query(query)
}
