package main

import (
	"fmt"
	"contact-go/handler"
	"contact-go/helper"
	"contact-go/repository"
)

var stopped  bool = false

func Menu() {
	contactRepo := repository.NewContactRepository()
	contactHandler := handler.NewContactHandler(contactRepo)

	fmt.Println("\nSelect menu")
	fmt.Println("1. List contact")
	fmt.Println("2. Add contact")
	fmt.Println("3. Delete contach")
	fmt.Println("4. Update contach")
	fmt.Println("5. Exit")

	var choose int
	fmt.Print("Select menu : \n")
	fmt.Scanln(&choose)

	switch choose {
	case 1 :
		contactHandler.List()
		helper.ClearScreeen()
		Menu()
	case 2 :
		contactHandler.Add()
		contactHandler.List()
		helper.ClearScreeen()
		Menu()
	case 3 :
		contactHandler.List()
		contactHandler.Delete()
		fmt.Printf("-----Updated Datas-----")
		contactHandler.List()
		helper.ClearScreeen()
		Menu()
	case 4 :
		contactHandler.List()
		contactHandler.Update()
		fmt.Printf("-----Updated Datas-----")
		contactHandler.List()
		helper.ClearScreeen()
		Menu()
	case 5:
		stopped = true
		return
	}
}

func main() {
	for !stopped {
		Menu()
	}
}