package main

import "fmt" 
 
func main() {
//  arr := [4][4]string{{"J", "L", "L", "M"}, {"u", "i", "i", "a"}, {"s", "v", "f", "n"}, {"t", "e", "e", ""}} //jika ingin test "arr" yang ini maka hilangkan status commentnya dan comment "arr" yang selanjutnya
 arr := [12][8]string{{"T", "M", "i", "t", "p", "o", "t", "c" },  { "h", "i", "s", "h", "o", "f", "h", "e" }, 
  {"e", "t", "", "e", "w", "", "e", "l" }, 
  { "", "o", "", "", "e", "", "", "l" }, 
  { "", "c", "", "", "r", "", "", "" }, 
  { "", "h", "", "", "h", "", "", "" }, 
  { "", "o", "", "", "o", "", "", "" }, 
  { "", "n", "", "", "u", "", "", "" }, 
  { "", "d", "", "", "s", "", "", "" }, 
  { "", "r", "", "", "e", "", "", "" }, 
  { "", "i", "", "", "", "", "", "" }, 
  { "", "a", "", "", "", "", "", "" }} 
 for i := 0; i < len(arr[0]); i++ { 
  for j := 0; j < len(arr); j++ { 
   fmt.Print(arr[j][i]) 
   if j >= len(arr)-1 { 
    fmt.Print(" ") 
   } 
  } 
 }
}