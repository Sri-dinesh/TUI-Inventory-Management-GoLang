# TUI Inventory Management

A small, terminal-based inventory manager written in Go. It provides a keyboard-driven text UI (TUI) for viewing and managing products stored in a simple JSON file (`inventory.json`). The project is focused on being lightweight, easy to run, and easy to extend.

---

## Features

- List, add, edit, and delete inventory items
- Search and sort items from the TUI
- Persist data to `inventory.json`
- Small, dependency-free Go codebase for easy modification

## Quick Start

1. Clone the repo and open the project directory:

   ```sh
   git clone https://github.com/Sri-dinesh/TUI-Inventory-Management-GoLang.git
   cd tui-inventory-management
   ```

2. Build and run:

   ```sh
   go build -o tui-inventory
   ./tui-inventory
   # or
   go run main.go
   ```

3. The app reads and writes `inventory.json` in the project root. A minimal example entry:

   ```json
   [
     {
       "id": 1,
       "name": "Widget",
       "qty": 10,
       "price": 2.5,
       "location": "A1"
     }
   ]
   ```

## Development & Testing

- Run: `go run main.go`
