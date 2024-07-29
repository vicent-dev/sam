package main

import (
	"sam/app"
)

func main() {
	s := app.NewServer()

	if err := s.Run(); err != nil {
		panic(err)
	}
}
