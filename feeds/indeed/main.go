package indeed

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"github.com/MondaleFelix/Jobbot/feeds"
	"github.com/MondaleFelix/Jobbot/models"
	"strings"

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
	config.host = "https://www.indeed.com/"
	return &PublicFeed{
		config: config,
		BaseFeed: feeds.NewBaseFeed(name),
	}
}

func (feed *PublicFeed) Connect() {
	counter := 0
	doc := feed.GetDocument(fmt.Sprintf(format: "%s/obs?q=golang&sort=date&fromage1&start", feed.config.host))
	doc.Find(selector: "td#resultsCol .jobsearch-SerpJobCard").Each(func(i int, s *goquery.Selection){
		if counter < feed.Limit() {
			id, exists := s.Attr(attrName: "data-jk")
			if exists {
				path := fmt.Sprintf(format: "viewjob?jk=%s", id)
				href := fmt.Sprintf(format: "/%s%s", feed.config.host, path)
				job := feed.GetDocument(href)

				title := job.Find(selector: ".jobsearch-JobInfoHeader-title-job").Text()
				salary := job.Find(selector: ".jobsearch-JobMetadataHeader-item").Text()
				position := job.Find(selector: ".jobsearch-DeskstopStickyContainer-subtitle").Children().Last().Text()
				company := job.Find(selector: ".jobsearch-DeskstopStickyContainer-subtitle").Children().First().Children().First().Text()

				apply, exists := job.Find(selector: "#applyButtonLinkContainer a").Attr(attrName: "href")
				if exists {
					saved, err := feed.SavePost(&models.Post{
						Path: path,
						Name: feed.Name(),
						Host: feed.config.host,
						Title: strings.TrimSpace(title), 
						Apply: strings.TrimSpace(apply), 
						Company: strings.TrimSpace(company),
						Salary: strings.TrimSpace(salary),
						Postion: strings.TrimSpace(position)

					})

					if err != nil {
						log.Fatal(err)
					}
					if saved {
						counter ++
					}
				}
			}
		}
	})
}