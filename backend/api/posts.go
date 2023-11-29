package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Niflnir/Dreame/database"
)

type createPost struct {
  title string
  body string
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }

  var c createPost
  err = json.Unmarshal(body, &c)
  if err != nil {
    panic(err)
  }

  database.CreatePost(database.DBCon, c.title, c.body)
}

