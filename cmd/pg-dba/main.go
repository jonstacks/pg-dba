package main

import (
	"fmt"
	"os"

	"github.com/jonstacks/pg-dba/pkg/config"
	"github.com/jonstacks/pg-dba/pkg/dba"
)

func fatal(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	admin := dba.New(config.DBConnectionString())
	fatal(admin.Run())
}
