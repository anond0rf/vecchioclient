package main

import (
	"fmt"
	"log"

	"github.com/anond0rf/vecchioclient/client"
)

func main() {

	vc := client.NewVecchioClient()

	// POSTING NEW THREAD
	thread := client.Thread{
		Board:    "b",
		Name:     "",
		Subject:  "",
		Email:    "",
		Spoiler:  false,
		Body:     "This is a new thread on board /b/",
		Embed:    "",
		Password: "",
		Sage:     false,
		Files:    []string{`C:\path\to\image1.jpg`, `C:\path\to\image2.png`},
	}

	id, err := vc.NewThread(thread)
	if err != nil {
		log.Fatalf("Unable to post thread %+v. Error: %v", thread, err)
	}
	fmt.Printf("Thread posted successfully (id: %d) - %+v\n", id, thread)

	// POSTING REPLY
	reply := client.Reply{
		Thread:   1,
		Board:    "b",
		Name:     "",
		Email:    "",
		Spoiler:  false,
		Body:     "This is a new reply to thread #1 of board /b/",
		Embed:    "",
		Password: "",
		Sage:     false,
		Files:    []string{},
	}

	id, err = vc.PostReply(reply)
	if err != nil {
		log.Fatalf("Unable to post reply %+v. Error: %v", reply, err)
	}
	fmt.Printf("Reply posted successfully (id: %d) - %+v\n", id, reply)
}
