package main

type ID int

type command struct {
	id        ID
	recipient string
	sender    string
	body      []byte
}