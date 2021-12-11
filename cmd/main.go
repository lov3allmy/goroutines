package main

import (
	"fmt"
	"io"
	"io/ioutil"
	logrus "log"
	"math/rand"
	"net/http"
	"time"
)

func f(n int) {
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", n)
	response, err := http.Get(url)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}(response.Body)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(string(data))
}

func main() {
	for i := 1; i <= 100; i++ {
		go f(i)
	}

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
