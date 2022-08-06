package main

import (
	"github.com/set2002satoshi/8-4/infrastructure"
)

func main() {
	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}