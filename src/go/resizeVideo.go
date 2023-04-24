package main

import (
	"compress-file/support"
	"encoding/json"
	"fmt"
	_ "fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	jsonData := os.Args[1]
	var params support.ParamVideo
	if err := json.Unmarshal([]byte(jsonData), &params); err != nil {
		panic(err)
	}

	if len(params.Destination) == 0 {
		support.Error("Please set your destination file!!")
	} else {
		_, err := os.Stat(params.Destination)
		if err != nil {
			if os.IsNotExist(err) {
				os.Mkdir(params.Destination, 0777)
			} else {
				support.Error("Directory invalid!!")
			}
		}
	}

	outputFilePath := params.Destination + "output_" + params.Filename
	maxFileSizeMB := 10
	if params.MaxSize > 0 {
		maxFileSizeMB = params.MaxSize
	}

	bitrate, err := getBitrate(params.FilePath)
	if err != nil {
		support.Error(err.Error())
	}

	targetBitrate := bitrate / 2
	for {
		cmd := exec.Command("ffmpeg", "-i", params.FilePath, "-b:v", strconv.Itoa(targetBitrate)+"k", "-maxrate", strconv.Itoa(targetBitrate)+"k", "-bufsize", strconv.Itoa(targetBitrate*2)+"k", "-threads", "0", "-vf", "scale=-2:720", "-c:a", "copy", "-y", outputFilePath)
		err := cmd.Run()
		if err != nil {
			support.Error(err.Error())
		}

		file, err := os.Stat(outputFilePath)
		if err != nil {
			support.Error(err.Error())
		}
		fileSizeMB := file.Size() / 1024 / 1024

		if fileSizeMB <= int64(maxFileSizeMB) {
			break
		}

		targetBitrate = targetBitrate / 2
	}

	err = os.Rename(outputFilePath, params.Destination+params.Filename)
	if err != nil {
		support.Error(fmt.Sprintf("Error: %s", err.Error()))
	}

	support.Success("Success", params.Destination+params.Filename)
}

func getBitrate(filePath string) (int, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=bit_rate", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	bitrate, err := strconv.Atoi(string(out[:len(out)-1]))
	if err != nil {
		return 0, err
	}

	return bitrate / 1000, nil
}
