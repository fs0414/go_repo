package main

import "fmt"

func main() {
  arr := []int{43, 12, 55, 2, 81, 37, 65}
  fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
  if len(arr) <= 1 {
    return arr
  }

  pivot := len(arr) / 2
  left, right := []int{}, []int{}

  for i, v := range arr {
    if v > arr[pivot] {
      right = append(right, arr[i])
    } else if v < arr[pivot] {
      left = append(left, arr[i])
    }
  }

  left_sorted := quickSort(left)
  right_sorted := quickSort(right)

  left_res := append(left_sorted, arr[pivot])
  full_sorted := append(left_res, right_sorted...)
  return full_sorted
}
