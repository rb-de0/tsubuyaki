package main

import (
	"fmt"
	"os"
)

type Router struct {
	Args        []string
	HandlerList []Handler
}

type Handler struct {
	Operation  string
	HandleFunc func([]string)
}

func NewRouter(args []string) *Router {
	router := Router{Args: args}
	return &router
}

func (r *Router) AddHandler(operation string, handleFunc func([]string)) {
	handler := Handler{Operation: operation, HandleFunc: handleFunc}
	r.HandlerList = append(r.HandlerList, handler)
}

func (r *Router) Execute() {
	if len(r.Args) <= 1 {
		fmt.Println("Error, no arguments")
		os.Exit(-1)
	}

	for index := range r.HandlerList {
		if r.Args[1] == r.HandlerList[index].Operation {
			r.HandlerList[index].HandleFunc(r.Args)
			return
		}
	}

	fmt.Println("Error, no matched operation")
	os.Exit(-1)
}
