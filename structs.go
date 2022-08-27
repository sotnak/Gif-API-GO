package main

type Tag struct {
	Name  string `json:"name"`
	Count string `json:"count"`
}

type Gif struct {
	Name string   `json:"name"`
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
}

type CountStruct struct {
	Count int `json:"count"`
}
