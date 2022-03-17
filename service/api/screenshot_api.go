package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ThumbnailRequest struct {
	Url string `json:"url"`
}

type ScreenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

type ScreenshotAPIResponse struct {
	Screenshot string `json:"screenshot"`
}

func ScreenShotApi(decodedRequest ThumbnailRequest) ScreenshotAPIResponse {
	// Create a struct with the parameters needed to call the ScreenshotAPI.
	apiRequest := ScreenshotAPIRequest{
		Token:          "TOKEN_HERE",
		Url:            decodedRequest.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	// Create a HTTP request.
	req, err := http.NewRequest("POST", "https://shot.screenshotapi.net/screenshot", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request.
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	// Tell Go to close the response at the end of the function.
	defer response.Body.Close()

	var apiResponse ScreenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)
	return apiResponse

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error ::: ", err)
		return
	}
}
