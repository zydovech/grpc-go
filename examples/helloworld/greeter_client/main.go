/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server .
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		//logger.Println("ListenAndServe at:localhost:11010")
		err := http.ListenAndServe(":9899", nil)
		if err != nil {
			//logger.Fatal("ListenAndServe: ", err)
		}
	}()
	for i := 0; i < 1; i++ {

		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
			grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}

		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		name := defaultName
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		ctx = context.WithValue(ctx, "aaaaa", "xxxxxxx")
		_, err = c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.MaxCallRecvMsgSize(1024), grpc.MaxCallSendMsgSize(1024))
		fmt.Println(err)
	}

	time.Sleep(1000 * time.Second)
	//log.Printf("Greeting: %s", r.GetMessage())
}
