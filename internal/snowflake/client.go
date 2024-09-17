package snowflake

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"

	sf "github.com/snowflakedb/gosnowflake"
)

type Client interface {
	GetDb() (*sql.DB, error)
	GetTenantDb(tenantName string, username string, password string, role string) (*sql.DB, error)
	GetTenantDbWithPrivateKey(tenantName string, username string, role string, privateKey string) (*sql.DB, error)
	CloseConnection(db *sql.DB) error
}

const DriverName string = "snowflake"

type snowflake struct {
	config *Config
}

func NewClient(config *Config) Client {
	return &snowflake{config: config}
}

func (s *snowflake) GetDb() (*sql.DB, error) {
	if s.config.PrivateKey != "" {
		return s.getDBWithPrivateKey()
	}
	cfg := &sf.Config{
		Account:   s.config.Account,
		User:      s.config.Username,
		Password:  s.config.Password,
		Warehouse: s.config.Warehouse,
		Database:  s.config.Database,
		Schema:    s.config.Schema,
		Role:      s.config.Role,
	}
	dsn, err := sf.DSN(cfg)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(DriverName, dsn)
	return db, err
}

func (s *snowflake) GetTenantDb(tenantName string, username string, password string, role string) (*sql.DB, error) {
	cfg := &sf.Config{
		Account:   s.config.Account,
		User:      username,
		Password:  password,
		Warehouse: s.config.Warehouse,
		Database:  tenantName,
		Schema:    s.config.Schema,
		Role:      role,
	}
	dsn, err := sf.DSN(cfg)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(DriverName, dsn)
	return db, err
}

func (s *snowflake) CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *snowflake) getDBWithPrivateKey() (*sql.DB, error) {
	privateKey, err := getPrivateKey(s.config.PrivateKey)
	if err != nil {
		return nil, err
	}
	cfg := &sf.Config{
		Account:       s.config.Account,
		User:          s.config.Username,
		PrivateKey:    privateKey,
		Authenticator: sf.AuthTypeJwt,
		Warehouse:     s.config.Warehouse,
		Database:      s.config.Database,
		Schema:        s.config.Schema,
		Role:          s.config.Role,
	}
	dsn, err := sf.DSN(cfg)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(DriverName, dsn)
	return db, err
}

func (s *snowflake) GetTenantDbWithPrivateKey(tenantName string, username string, role string, pemString string) (*sql.DB, error) {
	privateKey, err := getPrivateKey(pemString)
	if err != nil {
		return nil, err
	}
	cfg := &sf.Config{
		Account:       s.config.Account,
		User:          username,
		PrivateKey:    privateKey,
		Authenticator: sf.AuthTypeJwt,
		Warehouse:     s.config.Warehouse,
		Database:      tenantName,
		Schema:        s.config.Schema,
		Role:          role,
	}
	dsn, err := sf.DSN(cfg)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(DriverName, dsn)
	return db, err
}

func getPrivateKey(pemString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key := parseResult.(*rsa.PrivateKey)
	return key, nil
}
