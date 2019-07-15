package dasorm

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	_ "github.com/snowflakedb/gosnowflake"
)

func connectSnowflake(creds *Config) (*Connection, error) {
	connectionURL := fmt.Sprintf("%s:%s@%s/%s", creds.User, creds.Password, creds.Host, creds.Database)
	db, err := sqlx.Connect("snowflake", connectionURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Connection{
		DB:      db,
		Dialect: &snowflake{},
	}, nil
}

type snowflake struct{}

func (s *snowflake) Name() string {
	return "snowflake"
}

func (s *snowflake) TranslateSQL(sql string) string {
	return sql
}

func (s *snowflake) Create(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericCreate(db, model), "snowflake create")
}

func (s *snowflake) CreateMany(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericCreateMany(db, model), "snowflake create")
}

func (s *snowflake) Update(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericUpdate(db, model), "snowflake update")
}

func (s *snowflake) Destroy(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericDestroy(db, model), "snowflake destroy")
}

func (s *snowflake) DestroyMany(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericDestroyMany(db, model), "snowflake destroy many")
}

func (s *snowflake) SelectOne(db *sqlx.DB, model *Model, query Query) error {
	return errors.Wrap(genericSelectOne(db, model, query), "snowflake select one")
}

func (s *snowflake) SelectMany(db *sqlx.DB, models *Model, query Query) error {
	return errors.Wrap(genericSelectMany(db, models, query), "snowflake select many")
}

func (s *snowflake) SQLView(db *sqlx.DB, model *Model, format map[string]string) error {
	return errors.Wrap(genericSQLView(db, model, format), "snowflake sql view")
}

func (s *snowflake) CreateUpdate(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericCreateUpdate(db, model), "snowflake create update")
}
func (s *snowflake) CreateManyTemp(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericCreateManyTemp(db, model), "snowflake create many temp")
}

func (s *snowflake) CreateManyUpdate(db *sqlx.DB, model *Model) error {
	return errors.Wrap(genericCreateManyUpdate(db, model), "snowflake create update many")
}
