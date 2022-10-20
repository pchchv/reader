package reddit

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pchchv/reader/models"
)

type result struct {
	Articles []tmpArticle `xml:"entry"`
}

type link struct {
	Link string `xml:"href,attr"`
}

type tmpArticle struct {
	Author  string `xml:"author>name"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Link    link   `xml:"link"`
}

// Main function for catching stories
func Stories() []models.Article {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.reddit.com/r/inthenews/.rss", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res result
	err = xml.Unmarshal(body, &res)
	return joinArticles(res.Articles)
}

// Generates a list of articles
func joinArticles(tmpArticles []tmpArticle) (articles []models.Article) {
	for _, article := range tmpArticles {
		articles = append(articles, models.Article{
			URL:    article.Link.Link,
			Title:  article.Title,
			Author: article.Author[3:],
			Date:   123,
			Source: "Reddit",
		})
	}
	return articles
}
