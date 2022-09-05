package main

import (
	"fmt"
	"time"
)

var minAge = time.Duration(0)

func main() {
	t := time.Now()
	s := t.Add(-minAge)
	fmt.Printf("minAge=%d \n t=%d\n s=%d", minAge, t, s)

}
