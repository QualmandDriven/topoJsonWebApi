package main

type Country struct {
	Name string `json:"name"`
}

type City struct {
	Name string `json:"name"`
}

type Countries []Country
type Cities []City
