package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
    tables := []struct {
        x int
        y int
        n int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 2, 7},
    }
    for _, table := range tables {
        total := Sum(table.x, table.y)
        if total != table.n {
            t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
        }
    }
}

func TestSub(t *testing.T) {
    total := Sub(5, 5)
    if total != 10 {
        t.Errorf("Sub was incorrect, got: %d, want: %d.", total, 10)
    }
    tables := []struct {
        x int
        y int
        n int
    }{
        {1, 1, 2},
        {1, 2, 3},
        {2, 2, 4},
        {5, 2, 7},
    }
    for _, table := range tables {
        total := Sub(table.x, table.y)
        if total != table.n {
            t.Errorf("Sub of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
        }
    }
}
