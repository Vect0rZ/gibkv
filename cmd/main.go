package main

import (
	"fmt"

	"github.com/Vect0rZ/gibkv/db"
)

func main() {
	m := db.CreateIndexMap("BlogPosts")
	m.CreateTable("Users")
	m.CreateTable("Blogs")
	m.CreateTable("Posts")

	for i, k := range m.Tables {
		fmt.Println(i, k)
	}
}
