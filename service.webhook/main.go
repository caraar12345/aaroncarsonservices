package main

import (
  "net/http"
  "log"
  "fmt"
)

func handleAqlSMSWebhook(w http.ResponseWriter, r *http.Request) {
  fmt.Println("POST params were:", r.URL.Query())
}

func main() {
  port := ":8080"
  log.Println("Server started on", port)
  http.HandleFunc("/aql-sms", handleAqlSMSWebhook)
  log.Fatal(http.ListenAndServe(port, nil))
}
