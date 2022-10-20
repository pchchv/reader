package consumer

import (
	"github.com/pchchv/reader/hackernews"
	"github.com/pchchv/reader/models"
	"github.com/pchchv/reader/reddit"
)

// Implement merging stories for different packages and returning to the main package
func Stories() (stories []models.Article) {
	for _, x := range hackernews.Stories() {
		stories = append(stories, x)
	}
	// TODO: Implement retrieval of articles from medium
	for _, x := range reddit.Stories() {
		stories = append(stories, x)
	}
	return stories
}
