package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {

        fileName1 := "./exp_1-2-8-1.txt"
        fileName2 := "./log/expected.txt"

        fp1, err := os.Open(fileName1)
        if err != nil {
                fmt.Println(err)
                return
        }
        fp2, err := os.Create(fileName2)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer fp1.Close()
        defer fp2.Close()

        scanner := bufio.NewScanner(fp1)

        for scanner.Scan() {
                fp2.WriteString(scanner.Text())
                fp2.WriteString("\n")
        }
        if err = scanner.Err(); err != nil {
                fmt.Println("ファイル書き込みでエラー発生", err)
        }
}
