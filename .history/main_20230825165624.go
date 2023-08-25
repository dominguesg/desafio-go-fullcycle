package main

func main() {
	// START OMIT
	// Create a channel of strings.
	ch := make(chan string)

	// Send a string to the channel.
	ch <- "Hello World!"

	// Receive a string from the channel.
	str := <-ch
	// END OMIT
	println(str)
}