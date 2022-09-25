package main

import (

	// For sql, postgres driver
	"github.com/shunta0213/go_test/infrastructures"
)

func main() {
	infrastructures.Router.Run()
}
