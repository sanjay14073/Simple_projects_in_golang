package main

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

//Roast lc profiles
func main(){
		ctx := context.Background()

		client, err := genai.NewClient(ctx, option.WithAPIKey("YOUR_API_KEY_HERE"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer client.Close()
		
		model := client.GenerativeModel("gemini-1.5-flash")
		resp, err := model.GenerateContent(ctx, genai.Text("Your Question"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		printResponse(resp)
		
}
func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}