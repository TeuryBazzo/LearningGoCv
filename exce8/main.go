package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

//Writing to an avi file
func main() {
	video, error := gocv.VideoCaptureFile("video.mp4")
	defer video.Close()

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	width := video.Get(gocv.VideoCaptureFrameWidth)
	height := video.Get(gocv.VideoCaptureFrameHeight)
	fps := video.Get(gocv.VideoCaptureFPS)

	writer, error := gocv.VideoWriterFile("video_output.avi", "MJPG", fps, int(width), int(height), false)
	defer writer.Close()

	if error != nil {
		fmt.Println(error.Error())
		return
	}

	mat := gocv.NewMat()
	defer mat.Close()

	for {
		if ok := video.Read(&mat); !ok {
			fmt.Println("Não foi possivel ler o arquivo")
			return
		}
		error = writer.Write(mat)
		if error != nil {
			fmt.Println(error.Error())
		}
	}
}
