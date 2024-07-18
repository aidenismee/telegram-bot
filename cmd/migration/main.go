package main

import (
	"fmt"

	"github.com/nekizz/telegram-bot/internal/migration"
)

func main() {
	if err := migration.Run(); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}
}
