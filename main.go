package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func execInput(input string) error {
    input = strings.TrimSuffix(input, "\n")
    cmd := exec.Command(input)

    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout

    return cmd.Run()
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    currentDir, _ := os.Getwd()

    for {
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
