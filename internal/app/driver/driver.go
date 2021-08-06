package driver

import (
	"github.com/dtrwi/linebot/config"
	"github.com/jmoiron/sqlx"
)

type driver struct {
	Config config.Provider
}

func NewDriver(config config.Provider) *driver {
	return &driver{
		Config: config,
	}
}

func (d *driver) GetPostgreSQLConn() (*sqlx.DB, error) {
	return NewPostgreSQL(PostgreSQLOption{
		URL:         d.Config.GetString("postgres.url"),
		MaxIdleConn: d.Config.GetInt("postgres.max_idle_connections"),
		MaxOpenConn: d.Config.GetInt("postgres.max_open_connections"),
	})
}
