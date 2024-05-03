package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	// add/override key-value pairs to map by simply targeting a key that exists/doesn't exists and assign a new
	// value to it
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	// deletes they key and value attached to it from the underlying map itself
	delete(websites, "Facebook")
	fmt.Println(websites)
}
