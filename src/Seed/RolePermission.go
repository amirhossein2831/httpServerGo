package Seed

import (
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"log"
)

func CreatePermission() {
	permissions := []*model.Permission{
		{Name: "create-user", Description: "Creating user record"},
		{Name: "read-user", Description: "Reading user record"},
		{Name: "update-user", Description: "Updating user record"},
		{Name: "delete-user", Description: "Deleting user record"},

		{Name: "create-movie", Description: "Creating movie record"},
		{Name: "read-movie", Description: "Reading movie record"},
		{Name: "update-movie", Description: "Updating movie record"},
		{Name: "delete-movie", Description: "Deleting movie record"},

		{Name: "create-book", Description: "Creating book record"},
		{Name: "read-book", Description: "Reading book record"},
		{Name: "update-book", Description: "Updating book record"},
		{Name: "delete-book", Description: "Deleting book record"},
	}

	for _, p := range permissions {
		result := DB.GetInstance().GetDb().Create(p)
		if result.Error != nil {
			log.Fatalf("Error creating permission: %v", result.Error)
		}
	}
}

func CreateRole() {
	var userPermission []model.Permission
	result := DB.GetInstance().GetDb().Where("name LIKE ?", "%user").Find(&userPermission)
	if result.Error != nil {
		log.Fatalf("Error retrieving permissions: %v", result.Error)
	}

	var moviePermission []model.Permission
	result = DB.GetInstance().GetDb().Where("name LIKE ?", "%movie").Find(&moviePermission)
	if result.Error != nil {
		log.Fatalf("Error retrieving permissions: %v", result.Error)
	}
	var bookPermission []model.Permission
	result = DB.GetInstance().GetDb().Where("name LIKE ?", "%book").Find(&bookPermission)
	if result.Error != nil {
		log.Fatalf("Error retrieving permissions: %v", result.Error)
	}
	roles := []*model.Role{
		{Name: "user", Description: "Read and mutate the user", Permissions: userPermission},
		{Name: "movie", Description: "Read and mutate the movie", Permissions: moviePermission},
		{Name: "book", Description: "Read and mutate the book", Permissions: bookPermission},
	}

	for _, role := range roles {
		result := DB.GetInstance().GetDb().Create(role)
		if result.Error != nil {
			log.Fatalf("error while creatign role: %v", result.Error)
		}
	}

}
