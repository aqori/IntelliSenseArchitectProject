// cmd/intellisensearchitectproject/main.go
package main

import (
"flag"
"log"
"os"

"intellisensearchitectproject/internal/intellisensearchitectproject"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := intellisensearchitectproject.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
