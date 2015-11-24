package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/redis.v3"
)

func main() {
	hostAndPort := ":8080"

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	fmt.Printf("Server started on %s\n", hostAndPort)

	// HTTP handler to get request, parse subdomain, and send back the image URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		searchTerm := strings.ToLower(strings.Split(r.Host, ".")[0])
		if r.URL.Path != "/" {
			searchTerm = strings.ToLower(strings.Split(r.URL.Path[1:], ".")[0])
			searchTerm = strings.TrimSuffix(searchTerm, filepath.Ext(searchTerm))
		}

		imageURL, err := redisClient.Get("todoy:" + searchTerm).Result()
		if err == redis.Nil {
			imageURL = searcher(r.RemoteAddr, searchTerm)
			redisClient.Set("todoy:"+searchTerm, imageURL, time.Second*120)
		}
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Referer", imageURL)
		fmt.Fprintf(w, "<html><img src=\"%s\"></html>\n", imageURL)
	})

	http.ListenAndServe(hostAndPort, nil)
}

// method to perform a Google Image search
func searcher(userip string, term string) string {
	gis := GoogleImageSearch{}
	err := gis.Search(userip, term)
	if err != nil {
		fmt.Println(err)
		return "http://www.nedmartin.org/v3/amused/_img/something-went-terribly-wrong.jpg"
	}
	num := 0
	// do random stuff
	// rand.Seed(time.Now().UTC().UnixNano())
	// num := rand.Intn(len(gis.ResponseData.Results))
	return gis.ResponseData.Results[num].URL
}
