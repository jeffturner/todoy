package main

import (
	"fmt"
	"time"

	"gopkg.in/redis.v3"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	searcher := func(userip string, term string) string {
		cachedURL, err := redisClient.Get("todoy:" + term).Result()
		if err != redis.Nil {
			return cachedURL
		}

		gis := GoogleImageSearch{}
		err = gis.Search(userip, term)
		if err != nil {
			fmt.Println(err)
			return "http://www.nedmartin.org/v3/amused/_img/something-went-terribly-wrong.jpg"
		}
		num := 0
		// do random stuff
		// rand.Seed(time.Now().UTC().UnixNano())
		// num := rand.Intn(len(gis.ResponseData.Results))

		redisClient.Set("todoy:"+term, gis.ResponseData.Results[num].URL, time.Second*120)
		return gis.ResponseData.Results[num].URL
	}

	fmt.Println("Server started on :8080")
	DoServer("/", ":8080", searcher)
}
