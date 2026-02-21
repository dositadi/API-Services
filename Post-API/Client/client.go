package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "http://localhost:8080"

func HomeRequest() {
	home := "/"
	url := URL + home

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	defer response.Body.Close()

	responseBuffer, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Printf("Error reading response body: %v", err2)
		return
	}

	fmt.Println(string(responseBuffer))
}

func PostRequest() {
	postURL := URL + "/Posts"

	postJson := `{"id": "494","title": "My first post!","body":"You can't give up on your self"}`

	post := strings.NewReader(postJson)

	resp, err := http.Post(postURL, "application/json", post)
	if err != nil {
		fmt.Printf("Error reading response: %v", err.Error())
		return
	}

	defer resp.Body.Close()

	response, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Printf("Error reading response: %v", err1.Error())
	}

	fmt.Println(string(response))
}

func DeleteRequest() {
	postURL := URL + "/Posts/404"

	request, err := http.NewRequest(http.MethodDelete, postURL, http.NoBody)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}

	client := &http.Client{}

	resp, err2 := client.Do(request)
	if err2 != nil {
		fmt.Println("Error sending request: ", err2)
		return
	}

	defer resp.Body.Close()

	response, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("Error reading response: ", err)
	}
	fmt.Println(string(response))
}

func ListPosts() {
	home := "/Posts"
	url := URL + home

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	defer response.Body.Close()

	responseBuffer, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		fmt.Printf("Error reading response body: %v", err2)
		return
	}

	fmt.Println(string(responseBuffer))
}

func UpdatePosts() {
	postURL := URL + "/Posts/494"

	postJson := `{"id": "494","title": "My first post!","body":"To God be the Glory!","tags":["dositadi","learn2earn"]}`

	post := strings.NewReader(postJson)

	req, err := http.NewRequest(http.MethodPut, postURL, post)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}

	client := &http.Client{}

	resp, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("Error sending request: ", err2)
		return
	}

	defer resp.Body.Close()

	response, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("Error reading response: ", err)
	}
	fmt.Println(string(response))
}

func main() {
	HomeRequest()
	//PostRequest()
	DeleteRequest()
	UpdatePosts()
	ListPosts()
}
