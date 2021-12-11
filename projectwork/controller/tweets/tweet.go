package tweets

import (
	"encoding/json"
	"fmt"
	m_tweets "github.com/archi6830/workgo/projectwork/domen/tweet"
	"github.com/archi6830/workgo/projectwork/servaices"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
	for i, v := range servaices.TweetArr { //создаем цыкл который пробигает по массиву
		if v.Id == newTweetToAddIntoArray.Id { //проверка по ид
			servaices.TweetArr[i] = newTweetToAddIntoArray //меняю значение по индексу
			return
		}
	}
	servaices.TweetArr = append(servaices.TweetArr, newTweetToAddIntoArray) //добавляем в массив новый твит
}
func GetTweetById(c *gin.Context) {
	tweetId, err := strconv.ParseInt(c.Param("tweet_id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Can't parse id %s", c.Param("tweet_id")))
		return
	}
	foundedTweet := servaices.FindTweetById(tweetId)

	if foundedTweet == nil {
		fmt.Println("не знаю такого твита")
		c.String(http.StatusBadRequest, fmt.Sprintf("я с %d не нашел", tweetId))
		return
	}
	c.JSON(http.StatusOK, foundedTweet)
}
func GetAllTweets(c *gin.Context) {
	c.JSON(http.StatusOK, servaices.TweetArr)
}
