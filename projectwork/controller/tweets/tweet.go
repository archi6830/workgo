package tweets

import (
	"encoding/json"
	"fmt"
	m_tweets "github.com/archi6830/workgo/projectwork/domen/tweet"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "strconv"
)

var tweetArr []m_tweets.Tweet

func CreateTweet(c *gin.Context) {
	var newTweet m_tweets.Tweet
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("не могу вычитать тело запроса(body)")
		return
	}

	if err := json.Unmarshal(bytes, &newTweet); err != nil {
		fmt.Println(err)
		fmt.Println("не могу анмаршл джейсон")
		return
	}
	addOrUpdate(newTweet)
	fmt.Printf("Got tweet from request: %+v\n", newTweet)
	c.String(http.StatusCreated, fmt.Sprintf("Tweet %+v was created!", newTweet))
}
func addOrUpdate(newTweetToAddIntoArray m_tweets.Tweet) {
	for i, v := range tweetArr { //создаем цыкл который пробигает по массиву
		if v.Id == newTweetToAddIntoArray.Id { //проверка по ид
			tweetArr[i] = newTweetToAddIntoArray //меняю значение по индексу
			return
		}
	}
	tweetArr = append(tweetArr, newTweetToAddIntoArray) //добавляем в массив новый твит
}
func GetTweetById(c *gin.Context) {
	tweetId, err := strconv.ParseInt(c.Param("tweet_id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Can't parse id %s", c.Param("tweet_id")))
		return
	}
	foundedTweet := findTweetById(tweetId)

	if foundedTweet == nil {
		fmt.Println("не знаю такого твита")
		c.String(http.StatusBadRequest, fmt.Sprintf("я с %d не нашел", tweetId))
		return
	}

	c.JSON(http.StatusOK, foundedTweet)
}
func findTweetById(idForSearch int64) *m_tweets.Tweet {
	for i, v := range tweetArr {
		if idForSearch == v.Id {
			return &v
		}
		fmt.Printf("%d,%+v,\n", i, v)
	}
	return nil
}
func FindTweetByMessage(c *gin.Context) {
	tweetMessage := c.Param("tweet_message")
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
	for i, v := range tweetArr {
		if NewMessage == v.Message {
			MessageArr = append(MessageArr, v)
		}

		fmt.Printf("%d,%+v,\n", i, v)
	}
	return MessageArr
}
func GetAllTweets(c *gin.Context) {
	c.JSON(http.StatusOK, tweetArr)
}
