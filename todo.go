package main

type Todo struct {
	Id    string `json:"id"`
	Value string `json:"value" binding:"required"`
}
