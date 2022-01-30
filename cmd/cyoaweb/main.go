package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port to start web app")
	filename := flag.String("file", "gopher.json", "JSON file with cyoa story")
	flag.Parse()
	fmt.Printf("using the story in %s. \n.", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("starting the server at port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
