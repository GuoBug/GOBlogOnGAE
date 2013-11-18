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

func SaveTopic(w http.ResponseWriter, r *http.Request, topic *Topic) {
	c := appengine.NewContext(r)

	log.Println(topic.Id)

	k := datastore.NewKey(c, "Topic", "", topic.Id, nil)

	log.Println(k)

	_, err := datastore.Put(c, k, topic)

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

func GetAllTopic(w http.ResponseWriter, r *http.Request) []template.FuncMap {
	var err error
	var i int64
	var topic Topic

	data := []template.FuncMap{}

	log.Println("in all")
	for i = 0; i < 100; i++ {
		if err != nil {
			break
		}
		topic, err = GetTopic(w, r, i)

		data = append(data, template.FuncMap{"Topic": topic})
	}

	return data
}

/* 存数据库的原型函数
func SaveTopic(w http.ResponseWriter, r *http.Request, topic *Topic) {
	c := appengine.NewContext(r)

	log.Println(topic.TopicNum)

	k := datastore.NewKey(c, "Topic", "stringID", topic.TopicNum, nil)

	log.Println(k)

	_, err := datastore.Put(c, k, topic)

	if err != nil {
		log.Fatal(err)
		return
	}
}
*/

/* 取数据原型函数 *
func GetTopic(w http.ResponseWriter, r *http.Request, i int64) (Topic, error) {

	var topic Topic

	c := appengine.NewContext(r)

	log.Printf("I =[%v]", i)

	k := datastore.NewKey(c, "Topic", "stringID", i, nil)

	log.Println(k)

	err := datastore.Get(c, k, &topic)

	if err != nil {
		log.Fatal(err)
		return topic, err
	}

	return topic, err
}
*/
