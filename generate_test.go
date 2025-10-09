package main

import (
    "testing"
)

func TestGenerateLength(t *testing.T) {
    pass := Generate(10, "uln")

    if len(pass) != 10 {
        t.Errorf("expected length 10, got %d", len(pass))
    }
}