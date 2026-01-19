package main

import (
	"fmt"
	"log"

	"todo-manager/utils"
)

func main() {
	//admin password
	password := "Admin@123"

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	fmt.Println("===== ADMIN PASSWORD HASH =====")
	fmt.Println(hashedPassword)
}
