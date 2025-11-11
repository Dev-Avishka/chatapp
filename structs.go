package main

//this file will contain struct definitions

type Message struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	UserName string `json:"UserName"`
}

type ReqMessage struct {
	Content  string `json:"content"`
	UserName string `json:"UserName"`
}
