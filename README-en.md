# VecchioClient

[![English](https://img.shields.io/badge/lang-en-blue.svg)](README-en.md) [![Italiano](https://img.shields.io/badge/lang-it-blue.svg)](README.md) 

**VecchioClient** is a Go library that provides a client for posting on [vecchiochan.com](https://vecchiochan.com/).  
 
It wraps around the reverse-engineered `/post.php` endpoint of [NPFchan](https://github.com/fallenPineapple/NPFchan), abstracting away the details of form submission and request handling. 

## Features

- Post new threads to specific boards
- Reply to existing threads

Custom configuration is also possible for injecting custom `http.Client`, `User-Agent` and logger and for enabling verbose logging.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
   - [Basic Client Usage](#basic-client-usage)
     - [Posting a New Thread](#posting-a-new-thread)
     - [Posting a Reply](#posting-a-reply)
   - [Custom Client Configuration](#custom-client-configuration)

## Installation

To install VecchioClient, use `go get`:

```bash
go get github.com/anond0rf/vecchioclient
```

## Usage

### Basic Client Usage

VecchioClient offers a simple and straightforward API to interact with vecchiochan. Here's how to get started:

1. Import the client into your Go code:

    ```go
    import "github.com/anond0rf/vecchioclient/client"
    ```

2. Create a client:
   
    ```go
    vc := client.NewVecchioClient()
    ```

3. Use the client to interact with vecchiochan, such as posting a new thread or replying to an existing one.  

    - ##### Posting a New Thread
    
    ```go
    thread := client.Thread{
		Board:    "b",
		Name:     "",
		Subject:  "",
		Email:    "",
		Spoiler:  false,
		Body:     "This is a new thread on board /b/",   // Thread message
		Embed:    "",
		Password: "",
		Sage:     false,               // Prevents bumping and replaces email with "rabbia"
		Files:    []string{`C:\path\to\file.jpg`},
	}

    id, err := vc.NewThread(thread)
	if err != nil {
		log.Fatalf("Unable to post thread %+v. Error: %v", thread, err)
	}
	fmt.Printf("Thread posted successfully (id: %d) - %+v\n", id, thread)
    ```

    Please note that you do not need to set all fields to instantiate the `Thread` struct and you can do so with a smaller set:

    ```go
    thread := client.Thread{
		Board:    "b",
		Body:     "This is a new thread on board /b/",   // Thread message
		Files:    []string{`C:\path\to\file.jpg`},
	}
    ```

    In this case, default values will be assigned to the other fields.  
    **Board** is the only **mandatory** field checked by the client but keep in mind that, as the rules vary across boards and because of board settings, more fields are probably required for posting (e.g. you can't post a new thread with no embed nor files on /b/).

    - ##### Posting a Reply

    ```go
    reply := client.Reply{
		Thread:   1,
		Board:    "b",
		Name:     "",
		Email:    "",
		Spoiler:  false,
		Body:     "This is a new reply to thread #1 of board /b/",    // Reply message
		Embed:    "",
		Password: "",
		Sage:     false,            // Prevents bumping and replaces email with "rabbia"
		Files:    []string{`C:\path\to\file1.mp4`, `C:\path\to\file2.webm`},
	}

    id, err = vc.PostReply(reply)
	if err != nil {
		log.Fatalf("Unable to post reply %+v. Error: %v", reply, err)
	}
	fmt.Printf("Reply posted successfully (id: %d) - %+v\n", id, reply)
    ```

    Please note that you do not need to set all fields to instantiate the `Reply` struct and you can do so with a smaller set:

    ```go
    reply := client.Reply{
        Thread:   1,
		Board:    "b",
		Body:     "This is a new reply to thread #1 of board /b/",   // Reply message
	}
    ```

    In this case, default values will be assigned to the other fields.  
    **Thread** is the only **mandatory** field checked by the client but keep in mind that, as the rules vary across boards and because of board settings, more fields are probably required for replying.

### Custom Client Configuration

Custom client configuration is done by creating a `Config` struct with the needed values like in the example below:

```go
config := client.Config{
    Client:    &http.Client{Timeout: 10 * time.Second},                 // Custom HTTP client
    Verbose:   true,                                                    // Enable/Disable detailed logging
    UserAgent: "MyCustomUserAgent/1.0",                                 // Custom User-Agent
    Logger:    log.New(os.Stdout, "vecchioclient: ", log.LstdFlags),    // Custom logger
}

vc := client.NewVecchioClientWithConfig(config)
```
