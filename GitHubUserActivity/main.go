package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

func main() {

	//get username from user
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username github:")
	username, _ := read.ReadString('\n')
	username = username[:len(username)-1] // remove newline

	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch data: %s\n", resp.Status)
		return
	}
	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("\nRecent GitHub activity for %s:\n", username)
	for i, event := range events {
		if i >= 5 {
			break
		}
		fmt.Printf("- %s on %s\n", event.Type, event.Repo.Name)
	}

}
