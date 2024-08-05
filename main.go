package main

func main() {

	server := NewApiserver(":3000")
	server.Run()
}
