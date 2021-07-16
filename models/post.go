package models

type Post struct {

	ID bson.ObjectId 'bson:"_id,omitempty"'
	Path string
	Name string
	Host string
	Title string
	Salary string
	Position string
	Company string
	Processed bool
	Created time.Time
	Updated time.Time
}

type PostHandler struct {
	posts *mgo.Collection
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		posts: database.CreateConn().Posts,
	}
}

func (h *PostHandler) FindPosts(limit int)([]*Post, error){
	var ps []*Post
	return ps, h.posts.Find(bson.M{
		"processed" : false,
	}).Limit(limit).Sort(feild: : "-created").All(&ps)

}

func (h *PostHandler) Processed(ps []*Post]) error {
	bulk := h.posts.Bulk()
	for _, p := range ps {
		bulk.UpdateAll(bson.M{"_id":p.ID}, bson.M{
			"$set" :bson.M{"processed": true, "updated": time.Now()},
		})
	}
	
}

func (h *PostHandler) GetPostCount(name, path string) (int, error) {
	return h.posts.Find(bson.M{
		"name": name,
		"path": path,
	}).Count()
}

func (h *PostHandler) SavePost(post *Post) error {

	post.Created = time.Now()
	post.Processed = false
	return h.posts.Insert(post)
}

