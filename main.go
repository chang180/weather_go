package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

const apiKey = "CWB-3EEDFECA-EEDA-4627-91CF-7E50AA5D02C2"
const apiUrl = "https://opendata.cwa.gov.tw/api/v1/rest/datastore/O-A0003-001"

func main() {
    r := gin.Default()

    // 設置首頁路由
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to the Weather API!")
    })

    // 設置API路由
    r.GET("/weather", func(c *gin.Context) {
        url := fmt.Sprintf("%s?Authorization=%s", apiUrl, apiKey)

        resp, err := http.Get(url)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        var jsonResponse map[string]interface{}
        if err := json.Unmarshal(body, &jsonResponse); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
            return
        }

        c.JSON(http.StatusOK, jsonResponse)
    })

    // 運行服務器
    r.Run(":8080")
}
