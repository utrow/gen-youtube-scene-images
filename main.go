package main

import (
	"bytes"
	"fmt"
	"github.com/kkdai/youtube"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ErrFFmpegProcess = "failed ffmpeg process"
)

func main() {
	fmt.Print("[gen-youtube-scene-images]\n\n")

	// 動画URLの設定
	var url string
	fmt.Print("YouTube Video URL: ")
	_, err := fmt.Scan(&url)
	if err != nil {
		fmt.Println("input error")
	}

	// 動画の取得
	filename, err := videoDownload(url)
	if err != nil {
		fmt.Println("err:", err)
	}

	// FFmpegで処理
	err = genSceneImages(*filename)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func videoDownload(url string) (filename *string, err error) {
	y := youtube.NewYoutube(true)

	if err = y.DecodeURL(url); err != nil {
		return nil, err
	}

	videoDir := "./videos/"
	videoPath := filepath.Join(videoDir, y.VideoID+".mp4")
	fmt.Println("video:", videoPath)

	if fileExist(videoPath) {
		fmt.Println("video:", "Already exist.")
		return &videoPath, nil
	}

	if err := y.StartDownload(videoPath); err != nil {
		return nil, err
	}
	fmt.Println("video:", "Download complete!")

	return &videoPath, nil
}

func genSceneImages(filename string) error {
	// FFmpegのパラメータ定義
	cmd := exec.Command(
		"ffmpeg",
		"-i", filename,
		"-vf", "select=gt(scene\\,0.1),scale=960:540,showinfo",
		"-q:v", "0",
		"-vsync", "vfr", "output/%04d.jpg",
	)
	fmt.Println("exec:", cmd.Args)

	// 出力先を設定する
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// コマンドの実行
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())

		return errors.New(ErrFFmpegProcess)
	}

	// showinfoの情報を出力
	err = ioutil.WriteFile("output/ffout.log", stderr.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("FFmpeg:", "Process complete!")

	return nil
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
