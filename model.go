package main

type Country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Countries []Country
type Cities []City
