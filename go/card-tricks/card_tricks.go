package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
  if index >= len(slice) || index < 0 {
    return 0, false
  } else {
    return slice[index], true
  }
  // panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
  if index >= len(slice) || index < 0 {
		slice := append(slice, len(slice))
    return slice
	} else {
		slice[index] = value
    return slice
	}
  // panic("Please implement the SetItem function")
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
  if length < 0 {
    return nil
  }
  slice := make([]int, length)
  for i,_ := range slice {
    slice = SetItem(slice, i, value)
  }
  return slice
  // panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
  if index >= len(slice) || index < 0 {
    return slice
  }
  return append(slice[:index], slice[index+1:]...)
  // panic("Please implement the RemoveItem function")
}
