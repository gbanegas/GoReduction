package reduction

import (
  matrix "github.com/skelterjohn/go.matrix"
  "fmt"
)

func Reduct() int {
  ret := 0

  s := `[1 2 3;4 5 6]`
	A, err := matrix.ParseMatlab(s)
	Ar := matrix.MakeDenseMatrix([]float64{1,2,3,4,5,6}, 2, 3)
  fmt.Print(A)
  fmt.Print(err)
  fmt.Print(Ar)
  return ret
}
