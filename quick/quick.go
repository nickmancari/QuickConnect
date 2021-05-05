package quick

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Connect(method string, url string, data io.Reader, token string) []byte {

	request, err := http.NewRequest(method, url, data)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}
