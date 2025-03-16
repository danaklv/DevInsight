package config

import (
	"AiTaskManager/internal/models"
	"fmt"
	"log"
)

func CreateTables() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrate User: ", err)
	}

}

func InsertUser(user *models.User) error {
	id := DB.Create(user)
	fmt.Println(id)
	return nil
}
