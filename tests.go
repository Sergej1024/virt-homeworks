package main

import "testing"

func TestConvert(t *testing.T) {
    expected := 3.28
    received := convert(1)
    if received != expected {
	t.Errorf("Error, got: %f, want: %f", received, expected)
    }
}  
func TestMin(t *testing.T) {
    array := []int {15, 20, 30, 40, 50, 10, 60, 777}
    expected_index := 5
    expected_value := 10
    received_index, received_value := min(array)
    if received_index != expected_index || received_value != expected_value {
	t.Errorf("Error, got: %d, %d, want: %d, %d", received_index, received_value, expected_index, expected_value)
    }
}