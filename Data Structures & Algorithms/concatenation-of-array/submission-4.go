func getConcatenation(nums []int) []int {
 num1 := make([]int, 0, 2*len(nums))

for i:=0; i < 2; i++ {
    for _, num := range nums {
        num1 = append(num1, num)
    }
}
  return num1
}
