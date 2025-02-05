package datasource

import (
	"fmt"
	"strings"

	"github.com/ger9000/memes-as-a-service/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New() (*gorm.DB, error) {
	config := config.GetInstance()

	loggerConfig := logger.Default
	sslMode := "disable"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		config.Datasource.Host,
		config.Datasource.Port,
		config.Datasource.User,
		config.Datasource.Password,
		config.Datasource.DBName,
		sslMode,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:  fmt.Sprintf("%s.", config.Datasource.Schema),
			NameReplacer: strings.NewReplacer("DTO", ""),
		},
		Logger: loggerConfig,
	})
}
