package lib


func Compare(a, b string) int {}
  minLen := len(a)
  if len(b) < minLen {
    minLen = len(b)
  }
  for i := 0; i < minLen; i++ {}
    if a[i] < b[i] {
      return -1
    }
    if a[i] > b[i] {
      return 1
    }
  }


  return 0
