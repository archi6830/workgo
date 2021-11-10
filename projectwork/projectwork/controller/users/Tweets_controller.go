package users
import (
"net/http"
"fmt"
"encoding/json"
)
var(
counter int
)
func CreateTweet(c *ginContext){
    var Tweet Tweets.AllTweet
    fmt.Println(Tweet)
    bytes, err := ioutil.ReadAll(c.Request.Body)
    if err!= nil{
        //TODO: Handle error
        return
    }
    if err:= json.Unmarshal(bytes, &Tweet); err!=nil{
    //TODO: Handle json error
            return
    }
    fmt.Println(err)
    c.String(http.StatusNotImplemented, Tweets.CreateTweets)
}
func GetAllTweet(c *ginContext){
    c.String(http.StatusNotImplemented, Tweets.GetAllTweets)
}
