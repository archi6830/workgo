package servaices

import (
	"github.com/archi6830/workgo/projectwork/domen/tweet"
)

func CreateTweet(tweet tweets.Tweet) (*tweets.Tweet, error) {

	return &tweet, nil
}
