// TODO - TUI Inventory Management System

/*
	Functionalities:
		1. Add Item
		2. Remove Item
		3. Update Item
		4. View Inventory
		5. Search Item
		6. Save/Load Inventory to/from File

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "github.com/rivo/tview"
)

// Define an item structure that will hold the stock information
type Item struct {
	Name string `json:"name"`
	Stock int `json:"stock"`
}

var (
	inventory = []Item{}
	inventoryFile = "inventory.json"
)


// 1 - Load Inventory Function
func loadInventory() {
	if _,err := os.Stat(inventoryFile);
	err == nil {
		// Checking if file exists
		data, err := os.ReadFile(inventoryFile)

		if err != nil {
			log.Fatal("Error reading inventory file:", err)
		}

		json.Unmarshal(data, &inventory	)
	}
}

// 2 - Save Inventory Function
func saveInventory() {
	data, err := json.MarshalIndent(inventory, "", "  ")

	if err != nil {
		log.Fatal("Error saving inventroy:", err)
	}

	os.WriteFile(inventoryFile, data, 0644)

	if err != nil {
		log.Fatal("Error writing inventory file:", err)
	}
}


// 3 - Delete Item Function
func deleteItem(index int) {
	if index < 0 || index >= len(inventory) {
		fmt.Println("Invalid item index")
		return 
	}

	inventory = append(inventory[:index], inventory[index+1:]...)

	saveInventory()

}