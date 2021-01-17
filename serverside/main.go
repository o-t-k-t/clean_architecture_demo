package main

import (
	"log"

	"github.com/pkg/errors"

	"github.com/TechDepa/c_tool/infrastructures"
)

func main() {
	err := infrastructures.SetupRouter().Run()
	log.Fatal(errors.WithMessagef(err, "ルーティングが終了した"))
}
