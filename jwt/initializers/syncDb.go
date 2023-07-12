package initializers

import "github.com/naqet/learning-go/jwt/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{});
}
