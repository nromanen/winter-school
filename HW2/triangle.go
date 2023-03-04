package kata

func RowSumOddNumbers(n int) int {
	ret := 1
	for i := 0; i < 3; i++ {
		ret *= n
	}
	return ret

}

//triangle of consecutive odd numbers
//               1						== 1^3
//           3     5					== 2^3
//        7     9    11					== 3^3
//    13    15    17    19				== 4^3
// 21    23    25    27    29			== 5^3
