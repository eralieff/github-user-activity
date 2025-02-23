package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Action  string     `json:"action"`
		Commits []struct{} `json:"commits"`
		RefType string     `json:"ref_type"`
	} `json:"payload"`
}

func fetchGithubEvents(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("no recent activity found")
	}

	var events []Event
	if err = json.Unmarshal(body, &events); err != nil {
		return nil, err
	}

	return events, nil
}

func formatEvent(event Event) string {
	switch event.Type {
	case "PushEvent":
		return fmt.Sprintf("Pushed %d commit(s) to %s", len(event.Payload.Commits), event.Repo.Name)
	case "IssuesEvent":
		return fmt.Sprintf("%s an issue in %s", strings.ToUpper(string(event.Payload.Action[0]))+event.Payload.Action[1:], event.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s", event.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("Forked %s", event.Repo.Name)
	case "CreateEvent":
		return fmt.Sprintf("Created %s in %s", event.Payload.RefType, event.Repo.Name)
	default:
		return fmt.Sprintf("%s in %s", strings.TrimSuffix(event.Type, "Event"), event.Repo.Name)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <github-username>")
	}

	username := os.Args[1]
	events, err := fetchGithubEvents(username)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		fmt.Println("-", formatEvent(event))
	}
}
