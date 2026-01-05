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

type RepoItem struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Size        int    `json:"size"`
	HtmlUrl     string `json:"html_url"`
	DownloadUrl string `json:"download_url"`
}

func getFormattedSize(size int) (formatted string) {
	if size < 1024 {
		formatted = fmt.Sprintf("%d B", size)
	} else {
		kb := size / 1024
		if kb < 1024 {
			formatted = fmt.Sprintf("%d KiB", kb)
		} else {
			mb := size / 1024
			if mb < 1024 {
				formatted = fmt.Sprintf("%d MiB", mb)
			} else {
				gb := size / 1024
				formatted = fmt.Sprintf("%d GiB", gb)
			}
		}
	}
	return
}

func getListItems(items []RepoItem) string {
	var listItems string
	for _, item := range items {
		listItem := fmt.Sprintf("<div class='flex justify-between items-center w-full gap-8'><span>%s</span><span class='flex gap-4'><span>%s</span><a href='%s'>Github</a><a href='%s'>Download</a></span></div>", item.Name, getFormattedSize(item.Size), item.HtmlUrl, item.DownloadUrl)
		listItems = listItems + "\n" + listItem
	}
	return listItems
}

func main() {
	var url string = "https://api.github.com/repos/woof-os/woof-pacman-repo/contents/x86_64"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var items []RepoItem
	err = json.Unmarshal(body, &items)
	if err != nil {
		panic(err)
	}

	listItems := getListItems(items)

	contentBytes, err := os.ReadFile("template.html")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert the byte slice to a string
	fileContent := string(contentBytes)
	finalContent := strings.Replace(fileContent, "WOOFCONTENTS", listItems, 1)
	fmt.Println(finalContent)
}
