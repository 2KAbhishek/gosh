package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        currentDir, _ := os.Getwd()
        fmt.Printf("%s> ", currentDir)
        cmdString, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if err = execInput(cmdString); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}

func execInput(input string) error {
    input = strings.TrimSuffix(input, "\n")
    args := strings.Split(input, " ")

    switch args[0] {
    case "cd":
        if len(args) < 2 {
            return errors.New("path required")
        }
        return os.Chdir(args[1])
    case "exit":
        os.Exit(0)
    }

    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    return cmd.Run()
}
