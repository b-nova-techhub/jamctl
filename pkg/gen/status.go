package gen

import (
	"fmt"
	"sort"
	"time"
)

type Status struct {
	TotalPages        int    `json:"totalPages"`
	PublishedPages    int    `json:"publishedPages"`
	UnpublishedPages  int    `json:"unpublishedPages"`
	LastPublishedPage string `json:"lastPublishedPage"`
	LastGeneratedAt   string `json:"lastGeneratedAt"`
}

var CurrentStatus *Status

func SetStatus() {
	CurrentStatus = new(Status)
	CurrentStatus.TotalPages = len(GeneratedPages)
	CurrentStatus.PublishedPages = getPublished(GeneratedPages)
	CurrentStatus.UnpublishedPages = getUnpublished(GeneratedPages)
	CurrentStatus.LastPublishedPage = getLastPublished(GeneratedPages)
	CurrentStatus.LastGeneratedAt = time.Now().String()
	fmt.Printf("Generate Status: %+v\n", CurrentStatus)
}

func getPublished(pages []StaticPage) (count int) {
	count = 0
	for _, page := range pages {
		if page.IsPublished == "true" {
			count++
		}
	}
	return count
}

func getUnpublished(pages []StaticPage) (count int) {
	count = 0
	for _, page := range pages {
		if page.IsPublished == "false" {
			count++
		}
	}
	return count
}

func getLastPublished(pages []StaticPage) string {
	dates := make([]string, 0)
	for _, page := range pages {
		dates = append(dates, page.PublishDate)
	}
	sort.Strings(dates)
	return dates[len(dates)-1]
}
