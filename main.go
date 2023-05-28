package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "os/user"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        printPrompt()
        cmdString, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        if err = execInput(cmdString); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}

func printPrompt() {
    currentDir, _ := os.Getwd()
    dir := strings.Split(currentDir, "/")
    user, _ := user.Current()
    host, _ := os.Hostname()
    fmt.Printf("[%s@%s] in %s> ", user.Username, host, dir[len(dir)-1])
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
