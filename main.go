package main

import (
	"image"
	"os"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	port = "5440"
)

func getSprite(path string) (*pixel.Sprite, error) {
	toRead, err := os.Open(path)
	if err != nil {
		return &pixel.Sprite{}, err
	}
	defer toRead.Close()
	goImg, _, err := image.Decode(toRead)
	if err != nil {
		return &pixel.Sprite{}, err
	}
	img := pixel.PictureDataFromImage(goImg)
	// sprite := pixel.NewSprite(pic, pic.Bounds())
	return pixel.NewSprite(img, img.Bounds()), nil
}

const (
	winWidth  = 1920
	winHeight = 1080
)

func runGui(picChannel chan *pixel.Sprite) {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, winWidth, winHeight),
		VSync:       true,
		Undecorated: true,
	})

	if err != nil {
		panic(err)
	}
	myPic, err := getSprite("booster_landing.jpg")
	go func() {
		for {
			myPic = <-picChannel
		}
	}()
	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))
		myPic.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	transferChan := make(chan *pixel.Sprite)
	go func() {
		myPic, err := getSprite("portal_lock.jpg")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
		transferChan <- myPic
	}()
	pixelgl.Run(func() {
		runGui(transferChan)
	})
}
