package feeds

import (
	"context"
	"log"
	"strings"

	"github.com/MondaleFelix/Jobbot/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

type PublicFeed interface {
	Connect()                                 // Connects to the feed
	Name() string                             // Name of connected feed
	Limit() int                               // Cycle of post for the feed
	SavePost(post *models.Post) (bool, error) // Saves into DB
	GetDocument(url string) *goquery.Document // Parse Post
}

type BaseFeed struct {
	limit int
	name  string
	posts *models.PostsHandler
}

func NewBaseFeed(name string) *BaseFeed {
	return &BaseFeed{
		limit: 10,
		name:  name,
		posts: models.NewPostsHandler(),
	}
}

func (f *BaseFeed) Name() string {
	return f.name
}

func (f *BaseFeed) Limit() int {
	return f.limit
}

func (f *BaseFeed) SavePost(post *models.Post) (bool, error) {
	c, err := f.posts.GetPostsCount(post.Name, post.Path)
	if err != nil {
		return false, err
	}

	if c == 1 {
		return false, nil
	}

	return true, f.posts.SavePost(post)
}

func (f *BaseFeed) GetDocument(url string) *goquery.Document {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var doc *goquery.Document
	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)

			if err != nil {
				return err
			}

			doc, err = goquery.NewDocumentFromReader(strings.NewReader(res))
			if err != nil {
				return err
			}

			return nil

		}),
	}

	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	return doc
}
