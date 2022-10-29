package data

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	gorm.Model
	UUID      string
	Topic     string
	UserId    int
}

type Post struct {
	gorm.Model
	UUID      string
	Body      string
	UserId    int
	ThreadId  int
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	var c int64
	Db.Model(&Post{}).Where("thread_id = ?",thread.ID).Count(&c)

	return int(c)
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	posts = make([]Post,0)
	Db.Where("thread_id = ?", thread.ID).Find(&posts)

	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	uuid := createUUID()

	thread := &Thread{
		UUID: uuid,
		Topic: topic,
		UserId: int(user.ID),
	}

	// 创建记录
	Db.Create(thread)

	conv.CreatedAt = thread.CreatedAt
	conv.UserId = int(user.ID)
	conv.Topic = topic
	conv.UUID = uuid
	conv.ID = thread.ID

	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	uuid := createUUID()
	currTime := time.Now()

	p := &Post{
		UUID: uuid,
		Body: body,
		UserId: int(user.ID),
		ThreadId: int(conv.ID),
	}

	// 创建记录
	Db.Create(p)

	post.UUID = uuid
	post.CreatedAt = currTime
	post.Body = body
	post.UserId = int(user.ID)
	post.ThreadId = int(conv.ID)
	post.ID = p.ID

	return
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	Db.Find(&threads)

	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	Db.First(&conv, "uuid = ?", uuid)

	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	Db.First(&user, thread.UserId)

	return
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
	Db.First(&user, post.UserId)

	return
}
