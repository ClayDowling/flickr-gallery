package main

import (
	_ "encoding/json"
)

type Photo struct {
	Id     string
	Secret string
	Server string
	Farm   int
	Title  string
}

type PhotosetMeta struct {
	Id        string
	Primary   string
	Owner     string
	Ownername string
	Photos    []Photo `json:"photo"`
	Title     string
	Count     int `json:"total"`
}

type Pset struct {
	Content PhotosetMeta `json:"photoset"`
}

type EmbeddedContent struct {
	Content string `json:"_content"`
}

func (ec EmbeddedContent) String() string {
	return ec.Content
}

type Photoset struct {
	Id          string
	PhotosCount int             `json:"count_photos"`
	Title       EmbeddedContent `json:"title"`
	Owner       string
	Description EmbeddedContent `json:"description"`
	Photos      []Photo         `json:"photo"`
}

type Psets struct {
	Page   int
	Pages  int
	Total  int
	Albums []Photoset `json:"photoset"`
}

type PhotosetList struct {
	Photosets Psets
}
