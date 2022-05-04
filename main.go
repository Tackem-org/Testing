package main

import "fmt"

var (
	Version    string
	Commit     string
	CommitDate string
)

func main() {
	fmt.Println("Testing", Version, Commit, CommitDate)
	fmt.Println("Testing2", Version, Commit, CommitDate)
}
