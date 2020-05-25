package main

func main() {
	port := ":3000"
	server := NewServer(port)
	server.AddHandle("/", HandleRoot)
	server.AddHandle("/api", HandleAPI)
	server.Listen()
}
