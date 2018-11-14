package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

//reler exemplo do livro
func AlterVideo(tracker *gocv.Trackbar, initialPos float64, video *gocv.VideoCapture) {
	for {
		//action alter trackbar
		if initialPos != float64(tracker.GetPos()) {
			initialPos := float64(tracker.GetPos())
			video.Set(gocv.VideoCapturePosFrames, initialPos)
		}
	}
}

func main() {
	var NameImg string = os.Args[1]

	video, error := gocv.VideoCaptureFile(NameImg)
	defer video.Close()

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	window := gocv.NewWindow("sample")
	defer window.Close()

	countFrame := int(video.Get(gocv.VideoCaptureFrameWidth))

	tracker := window.CreateTrackbar("slide", countFrame)

	initialPos := float64(tracker.GetPos())

	go AlterVideo(tracker, initialPos, video)

	mat := gocv.NewMat()
	defer mat.Close()

	for {
		if ok := video.Read(&mat); !ok {
			fmt.Println("Não foi possivel ler o arquivo")
			return
		}

		window.IMShow(mat)
		window.WaitKey(1000)

		//i := video.Get(gocv.VideoCapturePosFrames)
		//tracker.SetPos(int(i))
	}

	window.WaitKey(0)
}
