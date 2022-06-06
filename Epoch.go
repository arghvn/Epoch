package main

//A common requirement in programs is getting the number of seconds, milliseconds,
//or nanoseconds since the Unix epoch. Here’s how to do it in Go.

import (
	"fmt"
	"time"
)

func main() {

	//Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time
	//since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

//we can also convert integer seconds or nanoseconds since the epoch into the corresponding time.

// output :
//2022-06-06 17:10:58.292387 +0000 UTC
//1351700038
//1351700038292
//1351700038292387000
//2022-06-06 17:10:58 +0000 UTC
//2022-06-06 17:10:58.292387 +0000 UTC
