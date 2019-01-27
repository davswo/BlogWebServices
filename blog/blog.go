package blog

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Blog struct{}

func GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://192.168.64.14/blogs")
	if err != nil {
		log.Panicf("Not able to reach backend node %v\n", err.Error())
		w.Write([]byte("Not able to reach backend 1"))
	}
	respBody, bodyerr := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Not able to convert backend response %v\n", bodyerr.Error())
		w.Write([]byte("Not able to reach backend 2"))
		return
	}
	if respBody == nil || len(respBody) == 0 {
		w.Write([]byte("Nothing to see here..."))
	}

	w.Write([]byte("Something actually happened"))
}
