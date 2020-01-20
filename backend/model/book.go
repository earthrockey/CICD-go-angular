package model

type Book struct {
	ID     int    `json="id"`
	Code   string `json="code"`
	Name   string `json="name"`
	Detail string `json="detail"`
}
