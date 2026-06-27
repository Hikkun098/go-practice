package main

import "fmt"

func main() {
	profile := map[string]string{
		"name": "Shuu",
		"age":  "30",
		"job":  "Engineer",
		"city": "Tokyo",
	}
	fmt.Println(profile)
	fmt.Println(profile["name"])
	fmt.Println(profile["age"])
	fmt.Println(profile["job"])
	fmt.Println(profile["city"])
	value, ok := profile["name"]
	fmt.Println(value, ok)
	profile["alias"] = "Shuu-chan"
	fmt.Println(profile["alias"])
	profile["age"] = "31"
	fmt.Println(profile["age"])
	delete(profile, "job")
	fmt.Println(profile)
}
