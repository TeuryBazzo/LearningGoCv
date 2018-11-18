package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	var NameImg string = os.Args[1]

	img := gocv.IMRead(NameImg, 1)
	defer img.Close()

	if img.Empty() {
		fmt.Println("img vazia")
	}

	windowIn := gocv.NewWindow("sample 5 - input")
	defer windowIn.Close()

	windowOut := gocv.NewWindow("sample 5 - output")
	windowOut.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowNormal)
	defer windowOut.Close()

	windowIn.IMShow(img)

	newMath := gocv.NewMat()

	chanels := img.Channels()
	fmt.Println("Number chanels ", chanels)

	size := img.Size()
	for _, s := range size {
		fmt.Println("size input", s)
	}

	arr := image.Point{
		X: size[0] / 2,
		Y: size[1] / 2,
	}

	gocv.PyrDown(img, &newMath, arr, gocv.BorderWrap)

	windowOut.IMShow(newMath)

	size = newMath.Size()
	for _, s := range size {
		fmt.Println("size output", s)
	}

	windowOut.WaitKey(0)
	windowIn.WaitKey(0)
}
