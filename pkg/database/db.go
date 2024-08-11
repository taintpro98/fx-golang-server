package database

import (
	"fmt"
	"fx-golang-server/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func PostgresqlDatabaseProvider(cnf *config.Config) *gorm.DB {
	db, err := NewPostgresqlDatabase(cnf.Database)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to db")
	}
	return db
}

func NewPostgresqlDatabase(databaseCnf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := GetDatabaseDSN(databaseCnf)
	customLogger := NewCustomLogger(dbLoggerConfig{
		ignoreRecordNotFoundError: false,
	})
	customLogger.logLevel = logger.Info
	client, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		},
	), &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", databaseCnf.Schema),
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = client.DB()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetDatabaseDSN(DBConf config.DatabaseConfig) string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s TimeZone=%s",
		DBConf.Host, DBConf.Port, DBConf.Username, DBConf.DatabaseName, "UTC",
	)

	if DBConf.SSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%s", DBConf.SSLMode)
	}

	if DBConf.Password != "" {
		dsn += fmt.Sprintf(" password=%s", DBConf.Password)
	}
	return dsn
}
