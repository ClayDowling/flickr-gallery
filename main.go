package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const KEY = "c11703fdb8e578fc8b3ba3ef1806cdc9"
const USER_ID = "63028497@N06"

func makeUrl(method string, properties map[string]string) string {

	args := fmt.Sprintf("api_key=%s&user_id=%s", KEY, USER_ID)
	for k, v := range properties {
		args = fmt.Sprintf("%s&%s=%s", args, k, v)
	}

	requesturl := fmt.Sprintf("https://www.flickr.com/services/rest/?method=%s&%s&format=json",
		method, args)

	return requesturl
}

func getContent(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body[14 : len(body)-1]
}

func getPhotosets() []Photoset {
	requesturl := makeUrl("flickr.photosets.getList", nil)
	body := getContent(requesturl)

	var response PhotosetList
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response.Photosets.Albums
}

func getPhotos(id string) PhotosetMeta {

	requesturl := makeUrl("flickr.photosets.getPhotos", map[string]string{
		"photoset_id": id})

	body := getContent(requesturl)
	var response Pset
	err := json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response.Content
}

func main() {

	listOnly := true
	var albumName string
	var templatename string

	flag.StringVar(&albumName, "album", "", "Album name to list")
	flag.StringVar(&templatename, "template", "hugo.tpl", "Template for link generation")
	flag.Parse()

	if albumName != "" {
		listOnly = false
	}

	albums := getPhotosets()

	fmt.Printf("Found %d photosets.\n", len(albums))

	for _, ps := range albums {
		if listOnly == true {
			fmt.Printf("\"%s\" id: %s\n", ps.Title, ps.Description)
		} else if albumName == ps.Title.String() {
			fmt.Printf("Getting album %s\n", ps.Title)
			pset := getPhotos(ps.Id)

			tmpl := template.Must(template.New("album").ParseFiles(templatename))
			err := tmpl.Execute(os.Stdout, pset)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
