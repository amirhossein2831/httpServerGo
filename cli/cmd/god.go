package cmd

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	firstName string
	lastName  string
	email     string
	password  string
)

var godCmd = &cobra.Command{
	Use:   "god",
	Short: "create god user",
	Run: func(cmd *cobra.Command, args []string) {
		createGod(firstName, lastName, email, password)
	},
}

func init() {
	rootCmd.AddCommand(godCmd)

	godCmd.Flags().StringVarP(&firstName, "firstName", "f", "admin", "first name of the god user")
	godCmd.Flags().StringVarP(&lastName, "lastName", "l", "admin", "last name of the god user")
	godCmd.Flags().StringVarP(&email, "email", "e", "admin@gmail.com", "email of the god user")
	godCmd.Flags().StringVarP(&password, "password", "p", "admin1234", "password of the god user")
}

func createGod(firstName, lastName, email, password string) {
	var roles []model.Role
	var userNum int64
	DB.GetInstance().GetDb().Model(&model.User{}).Where("email = ?", email).Count(&userNum)
	if userNum > 0 {
		log.Fatalf("This email is already used")
	}

	result := DB.GetInstance().GetDb().Find(&roles)
	if result.Error != nil {
		log.Fatalf("Error retrieving roles: %v", result.Error)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("can't hashed the password")
	}
	user := model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
		Roles:     roles,
	}
	err = DB.GetInstance().GetDb().Create(&user).Error
	if err != nil {
		log.Fatalf("can't create user: %v", err)
	}

	fmt.Printf("user created successfully\n")
}
