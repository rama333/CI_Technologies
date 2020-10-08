package main

import (
	"CI_Technologies/cmd/ci_ffmpeg"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	//test project cams repository
	numcpu := runtime.NumCPU()

	fmt.Println(numcpu)
	runtime.GOMAXPROCS(4)

	cams := []string{}

	cams = append(cams, "rtsp://streamer..")

	for i := 0; i < len(cams); i++ {
		go runSaveStreamOfCamera(cams[i], strconv.Itoa(i), 5)
		runtime.Gosched()
	}
	fmt.Scanln()
}

func runSaveStreamOfCamera(addresStream string, outputSave string, time2 int) {
	//dahua := "rtsp://streamer.tattelecom.ru/bd.kazan.ul..serova.46.podezd..vhodnaya.gruppa.-7880d3dab8"

	err := os.Mkdir("/Users/ramilramilev/go/src/CI_Technologies/archive/camera_"+outputSave, 0777)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i >= 0; i++ {

		fmt.Println(strconv.Itoa(i))

		output := "/Users/ramilramilev/go/src/CI_Technologies/archive/camera_" + outputSave + "/stream" + strconv.Itoa(i) + "_outbut.mp4"

		f := ci_ffmpeg.Get(&ci_ffmpeg.Config{
			FFMPEG: "/usr/local/bin/ffmpeg",
			Copy:   true,
			Audio:  false,
			Time:   10,
			Size:   int64(math.MaxInt64),
		})

		go func() {
			cmd, out, err := f.SaveVideo(addresStream, output, "stream -"+outputSave)

			log.Println(outputSave+" Command Used:", cmd)
			log.Println(outputSave+" Command Output:", out)
			log.Println(outputSave+" Command Output:", err)

			if err != nil {
				log.Fatalln(err)
			}
			runtime.Gosched()
		}()

		time.Sleep(10 * time.Second)

		//log.Println("Saved file from", dahua, "to", output)

		//s := ci_ffmpeg.Get(&ci_ffmpeg.Config{
		//})

		//go func() {
		//	cmd1, out1, err1 := s.MergeVideo("/Users/ramilramilev/go/src/CI_Technologies/archive/camera_"+outputSave+"/stream"+outputSave+"_input.m4v",
		//		"/Users/ramilramilev/go/src/CI_Technologies/archive/camera_"+outputSave+"/stream"+outputSave+"_output.m4v",
		//		"/Users/ramilramilev/go/src/CI_Technologies/archive/camera_"+outputSave+"/stream"+outputSave)
		//
		//	log.Println(outputSave+" Command Used:", cmd1)
		//	log.Println(outputSave+" Command Output:", out1)
		//	log.Println(outputSave+" Command Output:", err1)
		//
		//	if err1 != nil {
		//		log.Fatalln(err)
		//	}
		//}()
	}
}
