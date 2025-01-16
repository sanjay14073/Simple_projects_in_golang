package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type MXStructure struct {
    Exchange string `json:"exchange"`
    Priority int    `json:"priority"`
}

type Response struct {
    Domain     string       `json:"domain"`
    CreateDate string       `json:"create_date"`
    UpdateDate string       `json:"update_date"`
    Country    string       `json:"country"`
    IsDead     string         `json:"isDead"`
    RecordA    []string     `json:"A"`
    NsRecords  []string     `json:"NS"`
    Cname      string       `json:"CNAME"`
    Mx         []MXStructure `json:"MX"`
    Txt        []string       `json:"TXT"`
}

type APIResponse struct {
    Domains []Response `json:"domains"`
}

//A basic 

func main() {
    var companyName string
    fmt.Print("Enter company name: ")
    fmt.Scan(&companyName)

    conn := fmt.Sprintf("https://api.domainsdb.info/v1/domains/search?domain=%s&zone=com", companyName)
    res, err := http.Get(conn)
    if err != nil {
        log.Fatal("Error making HTTP request:", err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal("Error reading response body:", err)
    }

    var apiResponse APIResponse
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }

    fmt.Printf("Response: %+v\n", apiResponse.Domains)
}
