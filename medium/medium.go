package medium

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pchchv/reader/models"
)

type feedReturn struct {
	Payload struct {
		Collection struct {
			Slug string `json:"slug"`
		}
		Posts []struct {
			Title string `json:"title"`
			Date  int    `json:"updatedAt"`
			URL   string `json:"uniqueSlug"`
		}
	}
}

// Main function for catching stories
func Stories() []models.Article {
	feeds := []string{"message", "the-launchism"}
	var stories []models.Article
	for _, content := range feeds {
		dat := feedReturn{}
		err := json.Unmarshal(getFeed(content), &dat)
		if err != nil {
			log.Fatal(err)
		}
		for _, post := range dat.Payload.Posts {
			stories = append(stories, models.Article{
				URL:    "https://medium.com/" + dat.Payload.Collection.Slug + "/" + post.URL,
				Title:  post.Title,
				Author: "",
				Date:   post.Date,
				Source: "Medium",
			})
		}
	}
	return stories
}

func getFeed(feed string) []byte {
	req, err := http.Get("https://medium.com/" + feed + "/latest?format=json")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	return []byte(string(body)[16:])
}
