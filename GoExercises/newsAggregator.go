package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"html/template"
)

type NewsItem struct {
	Title string `xml:"title"`
	Desc string `xml:"description"`
	Link string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

type NewsIndex struct {
	Title string
	Item []NewsItem `xml:"channel>item"`
}

func newsHandler(writer http.ResponseWriter, request *http.Request) {
	res, error := http.Get("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	if error != nil {
		fmt.Println(error)
	}
	defer res.Body.Close() // resoponse body must be closed : postpone this call till the current function returns
	res_bytes, _ := ioutil.ReadAll(res.Body)
	
	//Parsing XML
	var NewsList NewsIndex
	xml.Unmarshal(res_bytes, &NewsList)
	
	NewsList.Title = "News Aggregator"
	/*fmt.Println(NewsList)
	
	for i, item := range NewsList.Item {
		fmt.Printf("%d -> %s\n", i, item.Title)
	}*/
	
	//Msg := NewsItem{"This was the message", "Description", "Link", "PubDate"}
	t, _ := template.ParseFiles("newsAggregator.html")
	fmt.Println(t.Execute(writer, NewsList)) // to log errors
}

func index_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer,"<h1>Welcome!</h1>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/news", newsHandler)
	http.ListenAndServe(":9000", nil)
}