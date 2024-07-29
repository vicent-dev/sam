package main

import "sam/app"

func main() {
	s := app.NewServer()

	if err := s.Migrate(); err != nil {
		panic(err)
	}
}
