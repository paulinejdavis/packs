package main

import "testing"

func TestCalculatePacks(t *testing.T) {
    testCases := []struct {
        orderSize int
        expectedPacks map[int]int
    }{
        {1, map[int]int{250: 1}},
        {250, map[int]int{250: 1}},
        {251, map[int]int{500: 1}},
        {501, map[int]int{500: 1, 250: 1}},
        {12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
    }

    for _, tc := range testCases {
        t.Run("OrderSize"+string(tc.orderSize), func(t *testing.T) {
            packs := calculatePacks(tc.orderSize)
            for size, count := range tc.expectedPacks {
                if packs[size] != count {
                    t.Errorf("For order size %d, expected %d packs of size %d, got %d", tc.orderSize, count, size, packs[size])
                }
            }
        })
    }
}