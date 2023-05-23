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

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := MainGo(name); err != nil {
		log.Fatal(err)
	}
	if err := Taskfile(name); err != nil {
		log.Fatal(err)
	}
	if err := GoModInit(name); err != nil {
		log.Fatal(err)
	}
	if err := OpenVSCode(name); err != nil {
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

func Taskfile(path string) error {
	content := fmt.Sprintf(`version: "3"

tasks:
  git:
    desc: fast commit
    cmds:
      - git add .
      - git commit -m "%s"
      - git push -u origin $(git rev-parse --abbrev-ref HEAD)
`, path)
	f, err := os.Create(path + "/Taskfile.yaml")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
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

	cmd := exec.Command("go", "mod", "init", "kek.com")
	cmd.Dir = currDir + "/" + path
	_, err = cmd.Output()
	return err
}

func OpenVSCode(path string) error {
	if runtime.GOOS == "windows" {
		return ErrArch
	}

	_, err := exec.Command("code", path).Output()
	return err
}
