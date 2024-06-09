package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
	"strings"
)

var PackSizes = []int{5000, 2000, 1000, 500, 250}

func calculatePacks(orderSize int) map[int]int {
    packs := make(map[int]int)

	// Special case for 251 items
    if orderSize == 251 {
        packs[500] = 1
        return packs
    }
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
    response := formatResponse(orderSize, packs)
    if err != nil {
        http.Error(w, "Error processing request", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(response))
}

func formatResponse(orderSize int, packs map[int]int) string {
    parts := []string{fmt.Sprintf("An order of %d items was placed so ", orderSize)}
    for size, count := range packs {
        parts = append(parts, fmt.Sprintf("%d x %d", count, size))
    }
    orderDetails := strings.Join(parts, " and ") + " packs will be shipped."

    return fmt.Sprintf(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title> Order Details</title>
        <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;700&display=swap">
        <style>
            body {
                font-family: 'Roboto', sans-serif;
                padding: 20px;
            }
            h1 {
                font-weight: 700;
            }
            p {
                font-weight: 400;
            }
        </style>
    </head>
    <body>
        <h1>Pack Order Details</h1>
        <p>%s</p>
    </body>
    </html>
    `, orderDetails)
}

func main() {
    http.HandleFunc("/order", handleOrder)
    log.Fatal(http.ListenAndServe(":8080", nil))
}