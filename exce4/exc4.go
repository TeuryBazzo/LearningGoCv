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

	windowIn := gocv.NewWindow("sample 4 - input")
	defer windowIn.Close()

	windowOut := gocv.NewWindow("sample 4 - output")
	defer windowOut.Close()

	windowIn.IMShow(img)

	newMath := gocv.NewMat()

	arr := image.Point{
		X: 50,
		Y: 50,
	}

	gocv.Blur(img, &newMath, arr)

	windowOut.IMShow(newMath)

	windowOut.WaitKey(0)
	windowIn.WaitKey(0)
}
