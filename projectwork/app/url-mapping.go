package app

import (
	"github.com/archi6830/workgo/projectwork/controller/ping"
	"github.com/archi6830/workgo/projectwork/controller/tweets"
	_ "github.com/archi6830/workgo/projectwork/controller/tweets"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	//http://localhost:8082/tweets/1 -
	router.GET("/tweets/:tweet_id", tweets.GetTweetById)
	router.GET("/tweets", tweets.GetAllTweets)
	router.GET("/search/tweets/:tweet_message", tweets.FindTweetByMessage)
	router.POST("/tweets", tweets.CreateTweet)
}
