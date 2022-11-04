package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/replit/database-go"
	"net/http"
	"os"
	"strings"
	"time"
)

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello\n")
}

func main() {

	myfunc := func() {
		for {
			config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
			token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))

			httpClient := config.Client(oauth1.NoContext, token)
			client := twitter.NewClient(httpClient)

			arr, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
				UserID: 798241286385868800,
			})
			if err != nil {
				fmt.Printf("err:%v\n", err)
			}
      
			parse := arr[0]
			//fmt.Println(parse.Text)

			//database.Set("key","gegsfgsg")
			post:=parse.Text
		

			value, _ := database.Get("key")
			if value != post {
				database.Set("key", post)

      if parse.InReplyToUserID==0||parse.InReplyToScreenName=="skrossigg"{
        	
        b := true
_, _, err := client.Statuses.Update(reverseString(post), &twitter.StatusUpdateParams{
					InReplyToStatusID:         parse.ID,
					AutoPopulateReplyMetadata: &b,
				})
				if err != nil {
					fmt.Printf("errr:%v", err)
				}
        
      }else{
        sr := strings.Split(parse.Text, " ")
	     var news string
	for _, n := range sr {
		if string(n[0]) == "@" {
			n = ""
			news = news +" "+ n
		}
		news = news +" "+ n
	}
    post=news
    b := true
    _, _, err := client.Statuses.Update(reverseString(post), &twitter.StatusUpdateParams{
					InReplyToStatusID:         parse.ID,
					AutoPopulateReplyMetadata: &b,
				})
				if err != nil {
					fmt.Printf("errr:%v", err)
				}
}
			}
			time.Sleep(5 * time.Second)
		}
	}
	go myfunc()
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
