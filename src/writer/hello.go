package writer

// Hello is a function to greet an user
func Hello(name string) string {
	if name == "" {
		name = "Go"
	}

	return "Hello " + name
}
