package main

import (
	"strconv"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
    testCases := []struct {
        orderSize     int
        expectedPacks map[int]int
    }{
        {1, map[int]int{250: 1}},
        {250, map[int]int{250: 1}},
        {251, map[int]int{500: 1}},
        {501, map[int]int{500: 1, 250: 1}},
        {12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
    }

    for _, tc := range testCases {
        t.Run("OrderSize"+strconv.Itoa(tc.orderSize), func(t *testing.T) {
            packs := calculatePacks(tc.orderSize)
            if len(packs) != len(tc.expectedPacks) {
                t.Errorf("For order size %d, expected %v packs, got %v", tc.orderSize, tc.expectedPacks, packs)
            }
            for size, count := range tc.expectedPacks {
                if gotCount, ok := packs[size]; !ok || gotCount != count {
                    t.Errorf("For order size %d, expected %d packs of size %d, got %d", tc.orderSize, count, size, gotCount)
                }
            }
        })
    }
}
