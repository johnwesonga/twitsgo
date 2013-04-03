package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type twitterResult struct {
	Results []struct {
		Text     string `json:"text"`
		Ids      string `json:"id_str"`
		Name     string `json:"from_user_name"`
		Username string `json:"from_user"`
		UserId   string `json:"from_user_id_str"`
	}
}

var twitterUrl = "http://search.twitter.com/search.json?q=%23KOT"

func retrieveTweets(c chan<- *twitterResult) {
	for {
		resp, err := http.Get(twitterUrl)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		r := new(twitterResult)
		err = json.Unmarshal(body, &r)
    if err != nil{
      log.Fatal(err)
    }	
		c <- r
	}

}

func displayTweets(c chan *twitterResult) {
  tweets := <-c
  for _, v := range tweets.Results{
    fmt.Printf("%v:%v\n", v.Username, v.Text)
  }
	
}

func main() {
	c := make(chan *twitterResult)
	go retrieveTweets(c)
	for {
		displayTweets(c)
	}

}
