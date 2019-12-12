package pkg

import (
	"fmt"

	"github.com/NanXiao/queue"
)

func Test() {
	s := queue.New()
	s.Enqueue(0)
	fmt.Println(s)
}
