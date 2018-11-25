package main

import (
	"fmt"
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

	windowIn := gocv.NewWindow("sample 7 - input")
	defer windowIn.Close()

	windowOut := gocv.NewWindow("sample 7 - output")
	windowOut.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowNormal)
	defer windowOut.Close()

	windowIn.IMShow(img)

	newMath := gocv.NewMat()

	//canny algorithmn
	gocv.Canny(img, &newMath, 200, 300)

	windowOut.IMShow(newMath)

	windowOut.WaitKey(0)
	windowIn.WaitKey(0)
}
