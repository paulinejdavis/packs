package main

import "testing"

func TestCalculatePacks(t *testing.T) {
    testCases := []struct {
        orderSize int
        expectedPacks map[int]int
    }{
        {1, map[int]int{250: 1}},
        {250, map[int]int{250: 1}},
        {251, map[int]int{250: 2}},
        {501, map[int]int{500: 1, 250: 1}},
        {12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
    }


}