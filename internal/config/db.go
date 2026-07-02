package config
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func ConnectDB(env *Config) (*gorm.DB) {
	db, err := gorm.Open(postgres.Open(env.DatabaseURL), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
