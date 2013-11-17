package models

import (
	"appengine"
	"appengine/datastore"

	"log"
	"net/http"
	"time"
)

type Topic struct {
	TopicNum int64
	Title    string
	Content  string
	Created  time.Time
}

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

/*
func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	e1 := Employee{
		Name:     "Joe Citizen",
		Role:     "Manager",
		HireDate: time.Now(),
		Account:  user.Current(c).String(),
	}

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "employee", nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 Employee
	if err = datastore.Get(c, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the Employee named %q", e2.Name)
}
*/
