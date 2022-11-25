package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			connec, err := net.Dial("tcp", fmt.Sprintf("%v:%v", "scanme.nmap.org", port))
			if err != nil {
				return
			}
			connec.Close()
			fmt.Printf("El Puerto %v Esta Abierto\n", port)
		}(i)
	}

	wg.Wait()
}
