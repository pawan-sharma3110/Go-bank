package main

func main() {
	server := NewAPIServer(":3000/account")
	server.Run()

}
