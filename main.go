
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

var coins = []string{"sol-idr", "bnb-idr", "eth-idr", "ptu-idr", "doge-idr"}

func main() {
    fmt.Println("🚀 AI Trader Monitor – Pintu Exchange")
    for _, coin := range coins {
        url := "https://api.pintu.co.id/v2/trade/price?id=" + coin
        resp, err := http.Get(url)
        if err != nil {
            fmt.Printf("Gagal mengambil data %s: %s\n", coin, err)
            continue
        }
        defer resp.Body.Close()

        var result map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
            fmt.Printf("Gagal decoding data %s: %s\n", coin, err)
            continue
        }

        if pairData, ok := result["payload"].(map[string]interface{}); ok {
            fmt.Printf("%s → Harga Beli: %s | Harga Jual: %s\n",
                coin,
                pairData["buy"],
                pairData["sell"],
            )
        }
    }
}
