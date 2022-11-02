package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/replit/database-go"
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
		// database.Set("key", "sfsfsf")
		parse := arr[0]

		name := "@" + parse.InReplyToScreenName + " "
		b := true
		post := strings.ReplaceAll(parse.Text, name, "")

		value, _ := database.Get("key")
		if value != post {
			database.Set("key", post)

			_, _, err := client.Statuses.Update(reverseString(post), &twitter.StatusUpdateParams{
				InReplyToStatusID:         parse.ID,
				AutoPopulateReplyMetadata: &b,
			})
			if err != nil {
				fmt.Printf("errr:%v", err)
			}

		}
		
	
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080",nil)

}
