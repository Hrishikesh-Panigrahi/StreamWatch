package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/Hrishikesh-Panigrahi/StreamWatch/render"
	"github.com/gin-gonic/gin"
)

func CreateResolution(inputPath, outputPath, resolution string) error {
	var ffmpegCmd *exec.Cmd
	switch resolution {
	case "480p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=854:480", outputPath)
	case "720p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1280:720", outputPath)
	case "1080p":
		ffmpegCmd = exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1920:1080", outputPath)
	default:
		return fmt.Errorf("unsupported resolution")
	}

	if err := ffmpegCmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}

func CreateMasterPlaylist(c *gin.Context, folderPath string) error {
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "playlist.txt", "-c", "copy", folderPath+"/master.m3u8")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}

func CreatePlaylist(c *gin.Context, folderPath string) error {
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "playlist.txt", "-c", "copy", folderPath+"/playlist.m3u8")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}

func CreateThumbnail(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ss", "00:00:01.000", "-vframes", "1", outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}

func GenerateMasterPlaylist(c *gin.Context, folderPath string, originalVideoPath string) error {
	masterPlaylist := folderPath + "/master.m3u8"
	cmd := exec.Command("ffmpeg", "-i", originalVideoPath,
		// HLS output options
		"-f", "hls", "-hls_time", "4", "-hls_playlist_type", "vod",
		"-hls_segment_filename", folderPath+"/segment_%03d.ts",
		masterPlaylist,
	)

	var stderrOutput bytes.Buffer
	cmd.Stderr = &stderrOutput

	// Run FFmpeg command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running FFmpeg command: ", err)
		fmt.Println("FFmpeg stderr: ", stderrOutput.String())
		render.RenderError(c, http.StatusInternalServerError, "Failed to create video")
		return fmt.Errorf("ffmpeg error: %v", err)
	}
	return nil
}
