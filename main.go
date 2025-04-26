package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if FileExists() {
		log.Fatal("main.go already exists")
	}
	CreateTemplate()
	OpenNeovim()
}

func FileExists() bool {
	_, err := os.Open("main.go")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateTemplate() {
	body := []byte("package main\n\nimport (\n  \"fmt\"\n)\n\nfunc main() {\n   fmt.Println(\"hello world\")\n}")
	file, err := os.Create("main.go")
	if err != nil {
		log.Fatal("error creating file: ", err)
	}
	_, err = file.Write(body)
	if err != nil {
		log.Fatal("error writing to file: ", err)
	}
}

func OpenNeovim() {
	cmd := exec.Command("nvim", "./main.go")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal("error executing nvim command: ", err)
	}
}
