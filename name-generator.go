package main

import (
	"github.com/goombaio/namegenerator"
	"github.com/o1egl/govatar"
	"log"
	"os"
	"time"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	err := govatar.GenerateFileForUsername(govatar.MALE, name, "image/avatar.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("image/avatar.jpg")
	if err != nil {
		return
	}
}
