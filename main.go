// Copyright 2016 Google, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kelseyhightower/memq/broker"
)

var (
	memqAddr  string
	queueName string
)

func main() {
	flag.StringVar(&memqAddr, "memq", "http://127.0.0.1:80", "memq HTTP address.")
	flag.StringVar(&queueName, "queue", "", "Work queue name.")
	flag.Parse()
	log.Printf("Starting worker...")

	if queueName == "" {
		log.Fatal("The -queue flag must be set and non empty")
	}
	log.Printf("Work queue set to %s", queueName)

	url := fmt.Sprintf("%s/queues/%s/messages", memqAddr, queueName)
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == http.StatusNoContent {
			// The queue is empty.
			log.Println("Work queue is empty. Shutting down...")
			os.Exit(0)
		}
		if resp.StatusCode != http.StatusOK {
			log.Fatal("Unexpected HTTP status code:", resp.Status)
		}

		var m broker.Message
		if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Processing message id: %s", m.Id)
		time.Sleep(time.Second)
	}
}
