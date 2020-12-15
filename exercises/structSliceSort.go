package main

/*
  - Create your own structure
  - Make a slice
  - use sort.Slice() to sort the elements of the slice
*/
import (
  "fmt"
  "sort"
)

type Tweet struct {
  tweet string
  retweets int
}

func main() {
  tweets := make([]Tweet, 0)
  tweet := Tweet{tweet: "Golang is awesome", retweets: 10}
  tweets = append(tweets, tweet)
  tweet = Tweet{tweet: "Clojure is awesome", retweets: 15}
  tweets = append(tweets, tweet)
  tweet = Tweet{tweet: "Haskell is awesome", retweets: 20}
  tweets = append(tweets, tweet)
  tweet = Tweet{tweet: "Linux is awesome", retweets: 100}
  tweets = append(tweets, tweet)

  sort.Slice(tweets, func(i, j int) bool {
    return tweets[i].retweets < tweets[j].retweets
  })

  for i := range tweets {
    fmt.Printf("The Tweet is: %v, and was retweeted: %v times\n", tweets[i].tweet, tweets[i].retweets)
  }

  sort.Slice(tweets, func(i, j int) bool {
    return tweets[i].retweets > tweets[j].retweets
  })

  for i := range tweets {
    fmt.Printf("The Tweet is: %v, and was retweeted: %v times\n", tweets[i].tweet, tweets[i].retweets)
  }
}
