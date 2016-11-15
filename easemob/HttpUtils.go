package easemob

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HTTPGet http get
func HTTPGet(url string, header map[string]string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return body
}

// HTTPPost http get
func HTTPPost(url string, value []byte) []byte {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(value))
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	// fmt.Printf(string(body))
	return body
}
