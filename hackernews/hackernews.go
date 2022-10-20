package hackernews

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/pchchv/reader/models"
)

type article struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Date   int    `json:"time"`
	Author string `json:"by"`
	Source string
}

const urlBase = "https://hacker-news.firebaseio.com/v0/"

// Main function for catching stories
func Stories() []models.Article {
	topIDs := topStories(30)
	var stories []models.Article
	var wg sync.WaitGroup
	for _, id := range topIDs {
		wg.Add(1)
		go func(i int) {
			stories = append(stories, models.Article(getStory(i)))
			wg.Done()
		}(id)
	}
	wg.Wait()
	return stories
}

// Implement the retrieval of history content
func getStory(id int) article {
	r, err := http.Get(urlBase + "item/" + strconv.Itoa(id) + ".json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data article
	err = json.Unmarshal(body, &data)
	data.Source = "Hackernews"
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Returns n best stories
func topStories(numStories int) []int {
	resp, err := http.Get(urlBase + "topstories.json?print=pretty")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var topArticles []int
	err = json.Unmarshal(body, &topArticles)
	if err != nil {
		log.Fatal(err)
	}
	return topArticles[:numStories]
}
