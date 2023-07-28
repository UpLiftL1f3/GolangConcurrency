package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID       int
	Comments []string
	Likes    int
	Friends  []int
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}

	jsonString, err := json.MarshalIndent(userProfile, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User Profile:\n%s\n", jsonString)
	fmt.Printf("Fetching the UserProfile took the program %+s long to complete", time.Since(start))
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respCh = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	//* we are doing three requests inside their own go routine
	go getComments(id, respCh, wg)
	go getLikes(id, respCh, wg)
	go getFriends(id, respCh, wg)

	//- adding 3 to the WaitGroup
	wg.Add(3)
	wg.Wait() //* block until wg counter == 0 we unblock
	close(respCh)

	userProfile := &UserProfile{}
	for resp := range respCh {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}

	}

	return userProfile, nil
}

func getComments(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey, that was great",
		"Yeah Buddy",
		"Ow, I didn't Know That",
	}

	respCh <- Response{
		data: comments,
		err:  nil,
	}

	// work is done
	wg.Done()
}

func getLikes(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respCh <- Response{
		data: 0,
		err:  nil,
	}

	// work is done
	wg.Done()
}

func getFriends(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respCh <- Response{
		data: []int{11, 34, 845, 55},
		err:  nil,
	}

	// work is done
	wg.Done()
}
