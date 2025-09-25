package main

import (
    "fmt"
    "log"
    "net/http"
)
// change name of the function to handler

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there new pipeline build from github add webhooks :)")
}

func main() {
    http.HandleFunc("/", handler)
    // Regex to match only numbers or exactly "Happy Thuresday"
    import "regexp"

    var re = regexp.MustCompile(`^\d+$|^Happy Thuresday$`)
    // Example usage:
    testStrings := []string{"12345", "Happy Thuresday", "abc123", "Thursday"}
    for _, s := range testStrings {
        fmt.Printf("'%s' matches: %v\n", s, re.MatchString(s))
    }
    log.Fatal(http.ListenAndServe(":8080", nil))
}
