package utils

import (
	"fmt"
	"os/exec"
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
