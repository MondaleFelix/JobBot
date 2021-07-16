package feeds

type PublicFeed interface {
	Connect()                                 // Connects to the feed
	Name() string                             // Name of connected feed
	Limit() int                               // Cycle of post for the feed
	SavePost(post *models.Post) (bool, error) // Saves into DB
	GetDocument(url string) *goquery.Document // Parse Post
}
