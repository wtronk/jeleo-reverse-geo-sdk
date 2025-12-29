package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	jeleogo "github.com/jeleo/reverse-geo-go"
)

func main() {
	apiKey := os.Getenv("JELEO_API_KEY")
	if apiKey == "" {
		log.Fatal("Set JELEO_API_KEY environment variable")
	}
	client, err := jeleogo.NewClient(apiKey, "")
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}
	var result map[string]any
	if err := client.ReverseGeo(48.239879, 9.216766, &result); err != nil {
		log.Fatalf("reverse geo failed: %v", err)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))
}
