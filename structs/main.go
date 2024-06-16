package main

import "fmt"

type Movie struct {
	Director    string
	Genre       string
	ReleaseDate string
}

func Structs_demo() {
	movie := Movie{"Christopher Nolan", "Action", "2002"}

	fmt.Printf("%+v", movie)

	anime := struct {
		name   string
		author string
	}{
		name:   "One Piece",
		author: "Oda",
	}

	fmt.Printf("%+v", anime)
}
