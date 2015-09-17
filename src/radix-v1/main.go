package main

import (

	"github.com/fzzy/radix/redis"
	"log"
	"fmt"
	"github.com/fzzy/radix/extra/pubsub"
)


/*
  This code requires the Redis Server to be running on port 6379 (default).
*/

func main() {

	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	r := client.Cmd("PING")
	if r.Err != nil {
		log.Fatal(err)
	}
	log.Printf("Results from PING: %s\n", r.String())


	r = client.Cmd("SET", "foo", "bar")
	if r.Err != nil {
		log.Fatal(err)
	}

	log.Printf("r.String(): %s\n", r.String())

	r = client.Cmd("GET", "foo")
	log.Printf("r.String(): %s\n", r.String())

	r = client.Cmd("SET", "foo_int", 123)

	r = client.Cmd("GET", "foo_int")
	i, _ := r.Int()
	log.Printf("i: %d\n", i)


	// Batching Commands
	client.Append("SET", "name", "Roberto")
	client.Append("GET", "name")

	r = client.GetReply() // Response from the SET
	log.Printf("Response from the SET: r.String(): %s\n", r.String())

	r = client.GetReply() // Response from the GET
	log.Printf("Response from the GET: r.String(): %s\n", r.String())

	// Use MGET
	client.Cmd("SET", "first_name", "Roberto")
	client.Cmd("SET", "last_name", "Rojas")

	r = client.Cmd("MGET", "first_name", "last_name")
	list, _ := r.List()
	for idx, value := range list {
		log.Printf("MGET - value(%d): %s\n", idx, value)
	}

	client.Close()

	// Pub-Sub Sample

	go func() {
		cli, _ := redis.Dial("tcp", "localhost:6379")
		i := 0
		for {
			i++
			cli.Cmd("PUBLISH", "news.tech", fmt.Sprintf("This is tech story #%d", i))
			cli.Cmd("PUBLISH", "news.sports", fmt.Sprintf("This is sports story #%d", i))
			if i > 10 {
				break
			}
		}
	}()

	go func() {
		cli, _ := redis.Dial("tcp", "localhost:6379")
		sub := pubsub.NewSubClient(cli)
		sr  := sub. PSubscribe("news.*")
		if sr.Err != nil {
			log.Fatal(sr.Err)
		}

		for  {
			r := sub.Receive()
			if r.Err != nil {
				log.Fatal(r.Err)
			}
			log.Printf("Pubsub Message: %s\n", r.Message)
		}


	}()

	fmt.Println("Press any key to stop program...")
	fmt.Scanln() // Keep the program running until any key is pressed...

}
