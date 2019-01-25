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
	}
	respBody, bodyerr := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Not able to convert backend response %v\n", bodyerr.Error())
		return
	}

	w.Write([]byte(respBody))
}
