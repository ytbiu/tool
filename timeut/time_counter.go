package timeut

import (
	"time"
	"fmt"
)

func Counter(name string,job func())  {
	now := time.Now()
	job()
	duration := time.Since(now).Seconds()
	fmt.Printf("name: %s, time duration: %0.3f seconds",name,duration)
}
