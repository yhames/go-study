package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	fileDir = "recorded"
)

func main() {
	http.HandleFunc("/", serveStaticHTML)
	http.HandleFunc("/upload", uploadHandler)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveStaticHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("video")
	if err != nil {
		log.Println("Failed to get form file:", err)
		return
	}
	defer file.Close()

	// The reason for using a temp file is to ensure that the file is fully written before processing it.
	// and retain the original file name format.
	now := time.Now().Format("2006-01-02_15-04-05")
	tempFile, err := os.CreateTemp(fileDir, fmt.Sprintf("%s_*.webm", now))
	if err != nil {
		log.Println("Failed to create temp file:", err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		log.Println("Failed to copy file:", err)
		return
	}

	// Convert the uploaded WebM file to MP4
	outputFile := fmt.Sprintf("%s/%s.mp4", fileDir, now)
	err = convertWebmToMp4(tempFile.Name(), outputFile) // Convert the uploaded WebM file to MP4
	if err != nil {
		log.Println("Failed to convert video:", err)
		return
	}

	// Serve the converted MP4 file
	http.ServeFile(w, r, outputFile)
}

/**
 * ffmpeg는 비디오 파일을 변환하는 데 사용되는 CLI 도구
 * 반드시 ffmpeg가 시스템에 설치되어 있어야 하며, PATH에 추가되어 있어야 합니다.
 */
func convertWebmToMp4(inputPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vcodec", "libx264", "-acodec", "aac", "-strict", "experimental", outputPath)
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to run ffmpeg command:", err)
		return err
	}
	return nil
}
