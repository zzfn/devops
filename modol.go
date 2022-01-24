package main

type Demo struct {
	Projects []Project
}
type Project struct {
	Name  string
	Url   string
	Build string
}
