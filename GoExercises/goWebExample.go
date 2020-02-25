package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"encoding/xml"
	"html/template"
)

type NewsItem struct {
	Title string `xml:"title"`
	Desc string `xml:"description"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type NewsIndex struct {
	Item []NewsItem `xml:"channel>item"`
}

func templateHandler(writer http.ResponseWriter, request *http.Request) {
	Msg := NewsItem{"This was the message", "Description", "Link", "PubDate"}
	t, _ := template.ParseFiles("basicHtml.html")
	t.Execute(writer, Msg)
}

func index_handler(writer http.ResponseWriter, request *http.Request) {
	res, error := http.Get("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	if error != nil {
		fmt.Println(error)
	}
	defer res.Body.Close() // postpone this call till the current function returns
	res_bytes, _ := ioutil.ReadAll(res.Body)
	
	res_string := string(res_bytes)
	fmt.Fprintf(writer, "Hello! This is the index page")
	fmt.Fprintf(writer, res_string)
	/*
	
	//Parsing XML
	var newsList NewsIndex
	xml.Unmarshal(res_bytes, &newsList)
	fmt.Println(newsList.Item)
	*/
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/parse", templateHandler)
	http.ListenAndServe(":9000", nil)
}
