package reduction

import (
  "github.com/skelterjohn/go.matrix"
)

func Reduct(exp List) ret int {

  ret := 0
  s := []
  for i:=exp[0]; i >=0; i-- {
    s[i] = i

  }
  A, err := ParseMatlab(s)


  return ret;
}
