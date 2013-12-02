package models

import (
	"appengine"
	"appengine/datastore"

	"log"
	"net/http"
	"text/template"
	"time"
)

type Topic struct {
	Id           int64
	Uid          int64
	Title        string
	Category     string
	Content      string
	Attachment   string
	Created      time.Time
	Updated      time.Time
	Views        int64
	Author       string
	ReplyTime    time.Time
	ReplyCount   int64
	ReplyLastUId int64
}

type Category struct {
	Id              int64
	Title           string
	TopicCount      int64
	TopicLastUserId int64
}

func SaveTopic(w http.ResponseWriter, r *http.Request, topic *Topic) {

	log.Println("save topic")

	c := appengine.NewContext(r)

	log.Println(topic.Id)

	k := datastore.NewKey(c, "Topic", "", topic.Id, nil)

	log.Println(k)

	_, err := datastore.Put(c, k, topic)

	if err != nil {
		log.Println(k)
	}

	return
}

func SaveCategroy(w http.ResponseWriter, r *http.Request, category *Category) {

	c := appengine.NewContext(r)

	/*存在 不再插 */
	_, err := GetCategory(w, r, category.Title)
	if err != nil {
		log.Printf("存在了，不再插入 %v", err)
		return
	}

	k := datastore.NewKey(c, "Category", category.Title, 0, nil)

	log.Printf("save category %v ***** %v", category, k)

	_, err = datastore.Put(c, k, category)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetTopic(w http.ResponseWriter, r *http.Request, i int64) (Topic, error) {

	var topic Topic

	c := appengine.NewContext(r)

	log.Printf("I =[%v]", i)

	k := datastore.NewKey(c, "Topic", "", i, nil)

	log.Println(k)

	err := datastore.Get(c, k, &topic)

	log.Println(i, err)

	return topic, err
}

func GetCategory(w http.ResponseWriter, r *http.Request, title string) (Category, error) {

	var category Category

	c := appengine.NewContext(r)

	k := datastore.NewKey(c, "Category", title, 0, nil)

	err := datastore.Get(c, k, &category)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(category, k)

	return category, err
}
func GetAllTopic(w http.ResponseWriter, r *http.Request) ([]template.FuncMap, int64) {
	var err error
	var i int64
	var topic Topic

	data := []template.FuncMap{}

	log.Println("in all")
	for i = 1; i < 100; i++ {
		topic, err = GetTopic(w, r, i)

		log.Println(err, topic)
		if err != nil {
			break
		}
		data = append(data, template.FuncMap{"Topic": topic})
	}

	return data, i
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) ([]template.FuncMap, int64) {
	var i int64
	var category []Category

	q := datastore.NewQuery("Category")

	c := appengine.NewContext(r)

	data := []template.FuncMap{}

	_, err := q.GetAll(c, &category)
	if err != nil {
		log.Fatal(err)
		return nil, 0
	}

	min, err := q.Count(c)
	if err != nil {
		log.Fatal(err)
		return nil, 0
	}

	log.Printf("look here ! count [%d] %v", min, q)

	for i = 0; i < int64(min); i++ {
		log.Println(category[i])
		data = append(data, template.FuncMap{"Category": category[i]})
	}

	return data, i

}
