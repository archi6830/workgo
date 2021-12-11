package servaices

import (
	"fmt"
	m_tweets "github.com/archi6830/workgo/projectwork/domen/tweet"
)

var TweetArr []m_tweets.Tweet

func AddOrUpdate(newTweetToAddIntoArray m_tweets.Tweet) {
	for i, v := range TweetArr { //создаем цыкл который пробигает по массиву
		if v.Id == newTweetToAddIntoArray.Id { //проверка по ид
			TweetArr[i] = newTweetToAddIntoArray //меняю значение по индексу
			return
		}
	}
	TweetArr = append(TweetArr, newTweetToAddIntoArray) //добавляем в массив новый твит
}
func FindTweetById(idForSearch int64) *m_tweets.Tweet {

	for i, v := range TweetArr {
		if idForSearch == v.Id {
			return &v
		}
		fmt.Printf("%d,%+v,\n", i, v)
	}
	return nil
}
func FindTweetByMessage(NewMessage string) []m_tweets.Tweet {
	var MessageArr []m_tweets.Tweet
	for i, v := range TweetArr {
		if NewMessage == v.Message {
			MessageArr = append(MessageArr, v)
		}

		fmt.Printf("%d,%+v,\n", i, v)
	}
	return MessageArr
}
