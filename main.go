package main

import (
	"fmt"
	"contact-go/config"
	"contact-go/handler"
	"contact-go/helper"
	"contact-go/repository"
	"os"
)

func main() {
	config := config.LoadConfig()

	var contactRepo repository.ContactRepositorier
	switch config.Storage {
	case "json":
		contactRepo = repository.NewContactJsonRepository()
	default:
		contactRepo = repository.NewContactRepository()
	}
	
	contactHandler := handler.NewContactHandler(contactRepo)
	Menu(contactHandler)
}

func Menu(contactHandler handler.ContactHandlerInterface) {
	fmt.Println("\nSelect menu")
	fmt.Println("1. List contact")
	fmt.Println("2. Add contact")
	fmt.Println("3. Update contach")
	fmt.Println("4. Delete contach")
	fmt.Println("5. Exit")

	var choose int
	fmt.Print("Select menu : \n")
	fmt.Scanln(&choose)

	switch choose {
	case 1 :
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 2 :
		contactHandler.Add()
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 3 :
		contactHandler.List()
		contactHandler.Update()
		fmt.Printf("------------Updated Datas------------")
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 4 :
		contactHandler.List()
		contactHandler.Delete()
		fmt.Printf("------------Updated Datas------------")
		contactHandler.List()
		helper.ClearScreeen()
		Menu(contactHandler)
	case 5:
		os.Exit(1)
	}
}