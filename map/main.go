package main

func main() {
	colors := map[string]string{
		"red": "ff0000",
		"blue": "ff0329",
		"green": "fr493409",
	}

	delete(colors, "red")

	for key, value := range colors {
		println(key, value)
	}
}
