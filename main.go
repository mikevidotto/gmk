package main

import (
	"github.com/mikevidotto/ff/ff"
	"log"
)

func main() {
	body := []byte("package main\n\nimport (\n  \"fmt\"\n)\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}")
	if err := ff.WriteIfNotExist("main.go", body); err != nil {
		log.Fatal(err)
	}
    if err := ff.OpenInEditor("nvim", "./main.go"); err != nil {
		log.Fatal(err)
	}
}
