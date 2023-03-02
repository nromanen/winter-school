package kata

type FParam func(int, int) int

func OperArray(fct FParam, arr []int, init int) []int {
result := make([]int, len(arr))
result[0] = fct(init, arr[0])
for i := 1; i < len(arr); i++ {
result[i] = fct(result[i-1], arr[i])
}
return result
}

func Som(x, y int) int {
return x + y
}

func Mini(x, y int) int {
if x < y {
return x
}
return y
}

func Maxi(x, y int) int {
if x > y {
return x
}
return y
}

func Lcmu(x, y int) int {
return abs(x*y) / Gcdi(x, y)
}

func Gcdi(x, y int) int {
if x == 0 {
return y
}
if y == 0 {
return x
}
if x < 0 {
x = -x
}
if y < 0 {
y = -y
}
for y != 0 {
x, y = y, x%y
}
return x
}

func abs(n int) int {
if n < 0 {
return -n
}
return n
}