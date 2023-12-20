package main

import (
	"log"

	"github.com/mowshon/moviego"
)

func main() {
	vid, _ := moviego.Load("video.mp4")
	vid.ResizeByWidth(100).Output("video-resized-by-width.mp4").Run()
	vid.ResizeByHeight(100).Output("video-resized-by-height.mp4").Run()
	vid.Resize(50, 50).Output("video-resized.mp4").Run()

	err := vid.SubClip(2, 5).Output("video-clip.mp4").Run()
	if err != nil {
		log.Fatal(err)
	}

	firstVid, _ := moviego.Load("video-clip.mp4")
	secondVid, _ := moviego.Load("video-clip.mp4")

	concVid, err := moviego.Concat([]moviego.Video{
		firstVid,
		secondVid,
		vid.SubClip(1, 2),
		vid.SubClip(1, 2).FadeIn(0, 0.5).FadeOut(0.5),
	})
	if err != nil {
		log.Fatal(err)
	}
	renderErr := concVid.Output("video-conc.mp4").Run()
	if renderErr != nil {
		log.Fatal(renderErr)
	}

	vid.Screenshot(2, "video-screen.png")
}
