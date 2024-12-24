package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type payload struct {
	Theme    string `json:"theme"`
	Username string `json:"username"`
}

func main() {
	fmt.Println("Please enter your GitHub ID:")
	var x string
	fmt.Scanf("%s", &x)

	data := payload{
		Theme:    "orange",
		Username: x,
	}

	// Encode the data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while encoding data:", err)
		return
	}
	//Using github unwrapped's url for getting videos(Personal project).
	res, err := http.Post("https://githubunwrapped.com/api/render", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error while getting the video:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading the response body:", err)
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error while unmarshaling the response:", err)
		return
	}

	fmt.Println("Your Recap available here ", responseData["url"])
}
