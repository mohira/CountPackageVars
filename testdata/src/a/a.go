package a

var (
	name string
	age  int
)

var address string
var _ = 100
var _ = 200
var x, y = 1, 2
var a, b, c = tuple3()

const Version = "1.0.0"

type (
	MyInt    int
	MyString string
)

func tuple3() (int, int, int) {
	return 3, 4, 5
}

func f() {
}
