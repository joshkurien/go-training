package main

import (
"fmt"
"time"
)

type user struct {
	name string
	DOB time.Time

}

func main() {
t:=user{"josh",time.Now()}
fmt.Printf("%v",t)

}