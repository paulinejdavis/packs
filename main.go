package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
)

var PackSizes = []int{5000, 2000, 1000, 500, 250}

func calculatePacks(orderSize int) map[int]int {
    packs := make(map[int]int)
    remaining := orderSize

    for remaining > 0 {
        for _, size := range PackSizes {
            if remaining >= size {
                count := remaining / size
                packs[size] += count
                remaining -= size * count
                break 
            }
        }
		
        if remaining > 0 && remaining < PackSizes[len(PackSizes)-1] {
            packs[PackSizes[len(PackSizes)-1]]++
            break
        }
    }
    return packs
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    orderSizeStr := query.Get("orderSize")
    if orderSizeStr == "" {
        http.Error(w, "Missing orderSize parameter", http.StatusBadRequest)
        return
    }

    orderSize, err := strconv.Atoi(orderSizeStr)
    if err != nil {
        http.Error(w, "Invalid orderSize parameter", http.StatusBadRequest)
        return
    }

    packs := calculatePacks(orderSize)
    response, err := json.Marshal(packs)
    if err != nil {
        http.Error(w, "Error processing request", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func main() {
    http.HandleFunc("/order", handleOrder)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
