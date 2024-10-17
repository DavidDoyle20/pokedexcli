package response

import (
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
	"time"
)

var cache = pokecache.NewCache(time.Second * 20)

func GetResponse(url string) ([]byte, error) {
	body, OK := cache.Get(url)

	if !OK {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			fmt.Println(res.Status)
			return nil, fmt.Errorf(res.Status)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cache.Add(url, body)
		//fmt.Println("Added to cache")
	}

	return body, nil
}
