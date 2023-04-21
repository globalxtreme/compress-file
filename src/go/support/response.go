package support

import (
	"encoding/json"
	"fmt"
	"os"
)

func Success(message string, destination string) {

	result := map[string]interface{}{
		"success":     true,
		"destination": destination,
		"message":     message,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		Error(err.Error())
	}
	fmt.Println(string(jsonResult))
	os.Exit(1)

}

func Error(message string) {

	result := map[string]interface{}{
		"success": false,
		"message": message,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonResult))
	os.Exit(1)

}
