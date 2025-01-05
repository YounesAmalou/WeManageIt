package main

import (
	"WeManageIt/internal/commands"

	_ "github.com/lib/pq"
)

func main() {
	commands.Execute()
}
