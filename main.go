package github.com/mikevidotto/gmk 

import (
	"log"
    "github.com/mikevidotto/ff"
)

func main() {
	body := []byte("package main\n\nimport (\n  \"fmt\"\n)\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}")
    if err := ff.WriteIfNotExist("main.go", body); err != nil {
        log.Fatal(err)
    }
}
