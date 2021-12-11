package servaices

import (
	"fmt"
	m_tweets "github.com/archi6830/workgo/projectwork/domen/tweet"
	"github.com/gin-gonic/gin"
	"net/http"
)

var TweetArr []m_tweets.Tweet

func FindTweetById(idForSearch int64) *m_tweets.Tweet {

	for i, v := range TweetArr {
		if idForSearch == v.Id {
			return &v
		}
		fmt.Printf("%d,%+v,\n", i, v)
	}
	return nil
}
func FindTweetByMessage(c *gin.Context) {
	tweetMessage := c.Query("searchString")
	var foundedTweet []m_tweets.Tweet
	foundedTweet = findTweetByMessage(tweetMessage)
	if foundedTweet == nil {
		fmt.Println("не знаю такого твита")
		c.String(http.StatusBadRequest, fmt.Sprintf("я с %s не нашел", tweetMessage))
		return
	}
	c.JSON(http.StatusOK, foundedTweet)
}
func findTweetByMessage(NewMessage string) []m_tweets.Tweet {
	var MessageArr []m_tweets.Tweet
	for i, v := range TweetArr {
		if NewMessage == v.Message {
			MessageArr = append(MessageArr, v)
		}

		fmt.Printf("%d,%+v,\n", i, v)
	}
	return MessageArr
}
