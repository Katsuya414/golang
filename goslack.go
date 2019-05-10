package main

import (
    "bytes"
    "fmt"
    "net/http"
)

func main() {
    name := "Go"
    text := "Hello from go"
    channel := "random"

    jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

    req, err := http.NewRequest(
        "POST",
        "https://hooks.slack.com/services/ほげほげ",
        bytes.NewBuffer([]byte(jsonStr)),
    )

    if err != nil {
        fmt.Print(err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Print(err)
    }

    fmt.Print(resp)
    defer resp.Body.Close()
}
