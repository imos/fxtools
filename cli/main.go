package main

import (
	"../hirose"
	"encoding/json"
	"flag"
	"log"
	"os"
	// utils "../utils"
	"fmt"
)

func Command(result interface{}, err error) {
	if err != nil {
		log.Fatal(err)
	}
	buf, err := json.Marshal(result)
	if err != nil {
		log.Fatal("failed to marshal: %v", err)
	}
	fmt.Println(string(buf))
	os.Exit(0)
}

func main() {
	// c := utils.HttpClient{}
	// c.Load("/tmp/cookie")
	// resp, err := c.Do("GET", "https://lionfx-mob.hirose-fx.co.jp/L001.html", map[string]string{})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", resp)
	// c.Save("/tmp/cookie")
	flag.Parse()
	h := hirose.HiroseClient{}
	command := flag.Args()[0]
	if command == "signin" {
		userId := os.Getenv("FX_USER_ID")
		if userId == "" {
			log.Fatalf("FX_USER_ID is required.")
		}
		password := os.Getenv("FX_PASSWORD")
		if password == "" {
			log.Fatalf("FX_PASSWORD is required.")
		}

		result, err := h.SignIn(userId, password)
		if err != nil {
			panic(err)
		}
		if !result {
			os.Exit(1)
		}

		result, err = h.IsSignedIn()
		if err != nil {
			panic(err)
		}
		if result {
			os.Exit(0)
		}
		log.Fatalf("Failed to sign in.")
		os.Exit(1)
	}

	result, err := h.IsSignedIn()
	if err != nil {
		panic(err)
	}
	if !result {
		log.Printf("Session is expired.")
		os.Exit(1)
	}
	if command == "check" {
		os.Exit(0)
	}

	switch command {
	case "status":
		{
			Command(h.GetStatus())
		}
	case "positions":
		{
			Command(h.GetPositions())
		}
	case "orders":
		{
			Command(h.GetOrders())
		}
	}
}
