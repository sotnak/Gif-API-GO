package main

type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Gif struct {
	Title string   `json:"title"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
}

type CountStruct struct {
	Count int `json:"count"`
}
