package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var (
	ErrArch = errors.New("can't execute this on a windows machine")
	ErrName = errors.New("provide single task name")
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal(ErrName)
	}
	name := args[0]
	if name == "" {
		log.Fatal(ErrName)
	}
	if err := MainGo(name); err != nil {
		log.Fatal(err)
	}
	if err := GoModInit(name); err != nil {
		log.Fatal(err)
	}
}

func MainGo(path string) error {
	mainGo := `package main

import (
	"fmt"
)
	
func main() {
	fmt.Println()
}
`
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(path + "/main.go")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(mainGo)
	if err != nil {
		return err
	}

	return nil
}

func GoModInit(path string) error {
	if runtime.GOOS == "windows" {
		return ErrArch
	}

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("curr dir is: ", currDir)

	cmd := exec.Command("go", "mod", "init", "kek.com")
	cmd.Dir = currDir + "/" + path
	_, err = cmd.Output()
	return err
}
