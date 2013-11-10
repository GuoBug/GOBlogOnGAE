package hello

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	/* 取当前回话的 */
	u := user.Current(c)
	if u == nil {
		/* 这是啥? */
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, Google App Engine, you are %v", u)

}
