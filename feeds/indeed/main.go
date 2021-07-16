package indeed

import (
	"fmt"
	"log"
	"strings"

	"github.com/MondaleFelix/Jobbot/feeds"
	"github.com/MondaleFelix/Jobbot/models"
	"github.com/PuerkitoBio/goquery"
)

type PublicFeedConfig struct {
	url  string
	host string
}

type PublicFeed struct {
	*feeds.BaseFeed
	config *PublicFeedConfig
}

func NewPublicFeed(name string) *PublicFeed {
	config := &PublicFeedConfig{}
	config.host = "https://www.indeed.com"
	return &PublicFeed{
		config:   config,
		BaseFeed: feeds.NewBaseFeed(name),
	}
}

func (feed *PublicFeed) Connect() {
	counter := 0
	doc := feed.GetDocument(fmt.Sprintf("%s/obs?q=golang&sort=date&fromage1&start", feed.config.host))
	doc.Find("td#resultsCol .jobsearch-SerpJobCard").Each(func(i int, s *goquery.Selection) {
		if counter < feed.Limit() {
			id, exists := s.Attr("data-jk")
			if exists {
				path := fmt.Sprintf("viewjob?jk=%s", id)
				href := fmt.Sprintf("/%s%s", feed.config.host, path)
				job := feed.GetDocument(href)

				title := job.Find(".jobsearch-JobInfoHeader-title-job").Text()
				salary := job.Find(".jobsearch-JobMetadataHeader-item").Text()
				position := job.Find(".jobsearch-DeskstopStickyContainer-subtitle").Children().Last().Text()
				company := job.Find(".jobsearch-DeskstopStickyContainer-subtitle").Children().First().Children().First().Text()

				apply, exists := job.Find("#applyButtonLinkContainer a").Attr("href")
				if exists {
					post := &models.Post{
						Path:     path,
						Name:     feed.Name(),
						Host:     feed.config.host,
						Title:    strings.TrimSpace(title),
						Apply:    strings.TrimSpace(apply),
						Company:  strings.TrimSpace(company),
						Salary:   strings.TrimSpace(salary),
						Position: strings.TrimSpace(position),
					}
					saved, err := feed.SavePost(post)
					if err != nil {
						log.Fatal(err)
					}
					if saved {
						log.Println(fmt.Sprintf("Post:%v saved successfully ", post))
						counter++
					}
				}
			}
		}
	})
}
