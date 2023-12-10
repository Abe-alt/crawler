package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/crawl", crawlHandler)
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

func crawlHandler(w http.ResponseWriter, r *http.Request) {
	// Get the website URL from the query parameters
	websiteURL := r.URL.Query().Get("websiteUrl")

	// Check if the URL is provided
	if websiteURL == "" {
		http.Error(w, "Website URL is required", http.StatusBadRequest)
		return
	}

	// Execute the JavaScript code using Node.js
	cmd := exec.Command("node", "--no-warnings", "main.js", websiteURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing crawl script: %s\n", err) //////
		fmt.Printf("Script Output:\n%s\n", output)            ////////
		http.Error(w, fmt.Sprintf("Error executing crawl script: %s", err), http.StatusInternalServerError)
		return
	}

	// Parse the output as JSON
	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		fmt.Printf("Error decoding ////: %s\n", err) //////
		fmt.Printf("Script Output:\n%s\n", output)   ////////
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write the JSON response
	json.NewEncoder(w).Encode(result)
}
