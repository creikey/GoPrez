package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/faiface/pixel"
)

// Serve runs on a loop prompting the user for images to serve
func Serve(picChannel chan *pixel.Sprite) {
	screenReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter a path for an image to load: ")
		toLoad, err := screenReader.ReadString('\n')
		toLoad = toLoad[:len(toLoad)-1]
		if err != nil {
			panic(err)
		}
		toShow, err := getSprite(toLoad)
		if err != nil {
			fmt.Println("ERROR LOADING IMAGE")
			fmt.Println("ERROR: ")
			fmt.Println(err)
			continue
		}
		fmt.Println(fmt.Sprintf("Successfully loaded image from %s!", toLoad))
		picChannel <- toShow
	}
}
