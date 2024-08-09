package numbers

// Every signed integer number type
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Every unsigned integer number type
type Unsigned interface {
  ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Every integer (signed or not) number type
type Integer interface {
  Signed | Unsigned
}

// Every floating point number type
type Float interface {
		~float32 | ~float64
}

// Every complex number type
type Complex interface {
		~complex64 | ~complex128
}

// Every real number
type Real interface {
  Integer | Float
}

// Every number type
type Number interface {
  Real | Complex
}

// Sums any given amount of numbers of the same type T
func Sum[T Number](ns ...T) T {
	var sum T = 0
	for _, v := range ns {
		sum += v
	}

	return sum
}

// Basic division of two numbers of same type
// Returns 0 if a division by zero is caught
func Divide[T Number](a T, b T) T {
  if b == 0 {
    return 0
  }

  return a / b
}
