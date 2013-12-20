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

	log.Printf("After put %v", err)

	if err != nil {
		log.Println(k)
	}

	http.Redirect(w, r, "/topic", 303)

	return
}

func SaveCategroy(w http.ResponseWriter, r *http.Request, category *Category) {

	c := appengine.NewContext(r)
	/* 先设置 key */
	k := datastore.NewKey(c, "Category", category.Title, 0, nil)

	log.Printf("save category %v ***** %v", category, k)

	/*存在 不再插 */
	log.Printf("title [%s]", category.Title)
	tmpCategory, err := GetCategory(w, r, category.Title)

	if err == nil {
		log.Printf("存在了，不再插入 %v", err)
		category.TopicCount = tmpCategory.TopicCount + 1
	} else {
		category.TopicCount = 1
	}

	_, err = datastore.Put(c, k, category)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func UpdateCategroy(w http.ResponseWriter, r *http.Request, category *Category, flag int) {

	/* flag == 0  累加 */
	/* flag == 1  删除 */
	log.Printf("Come to update category")

	c := appengine.NewContext(r)
	/* 先设置 key */
	k := datastore.NewKey(c, "Category", category.Title, 0, nil)

	log.Printf("save category %v ***** %v", category, k)

	/*存在 不再插 */
	log.Printf("title [%s]", category.Title)
	_, err := GetCategory(w, r, category.Title)

	if err == nil {
		log.Printf("存在了，不再插入 %v", err)
		if flag == 0 {
			category.TopicCount++
		} else if flag == 1 {
			category.TopicCount--
		}
	} else {
		if flag == 0 {
			SaveCategroy(w, r, category)
		} else if flag == 1 {
			log.Fatal("Error update But have no data !")
			return
		}
	}

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

	log.Println(category, k)

	return category, err
}

func GetAllTopic(w http.ResponseWriter, r *http.Request) ([]template.FuncMap, int64) {
	var err error
	var i int64
	var topic []Topic

	q := datastore.NewQuery("Topic").Order("Id")

	c := appengine.NewContext(r)

	data := []template.FuncMap{}

	_, err = q.GetAll(c, &topic)
	if err != nil {
		log.Fatal(err)
		return nil, 0
	}

	min, err := q.Count(c)
	if err != nil {
		log.Fatal(err)
		return nil, 0
	}

	for i = 0; i < int64(min); i++ {
		log.Println(topic[i])
		data = append(data, template.FuncMap{"Topic": topic[i]})
	}

	log.Printf("If here number is Need [%v]", topic[i-1].Id)

	return data, topic[i-1].Id
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

func DeleteTopic(w http.ResponseWriter, r *http.Request, i int64) error {
	// 删除同时减少计数器！不要忘了！

	var topic Topic

	c := appengine.NewContext(r)

	k := datastore.NewKey(c, "Topic", "", i, nil)
	log.Println(k)

	err := datastore.Get(c, k, &topic)

	err = DeleteCategory(w, r, topic.Category)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = datastore.Delete(c, k)
	if err != nil {
		log.Fatal(err)
		return err
	}

	http.Redirect(w, r, "/", 303)

	return err
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, title string) error {
	c := appengine.NewContext(r)

	category, err := GetCategory(w, r, title)
	if err != nil {
		log.Fatalf("类型不存在[%v]", title)
		return err
	}

	if category.TopicCount > 1 {

		log.Println("更新 category [%v]", category)
		UpdateCategroy(w, r, &category, 1)

	} else if category.TopicCount == 1 {

		k := datastore.NewKey(c, "Category", title, 0, nil)
		log.Println(k)
		err = datastore.Delete(c, k)
		if err != nil {
			log.Fatal(err)
			return err
		}
	} else {
		log.Fatalf("出错了！[%v]", category)
		return err
	}

	log.Printf("结束 delete category")

	return err
}
