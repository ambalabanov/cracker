package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/ambalabanov/nthash"
)

var wg sync.WaitGroup
var hash = flag.String("hash", "de26cce0356891a4a020e7c4957afc72", "NTHash")
var fWordlistName = flag.String("wordlist", "wordlist.txt", "wordlist file")

func main() {
	runtime.GOMAXPROCS(4)
	flag.Parse()
	fWordlist, err := os.Open(*fWordlistName)
	if err != nil {
		panic(err)
	}
	defer fWordlist.Close()

	scanner := bufio.NewScanner(fWordlist)
	for scanner.Scan() {
		pass := scanner.Text()
		wg.Add(1)
		go brute(pass)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	wg.Wait()
}
func brute(pass string) {
	if nthash.CheckHash(pass, *hash) {
		fmt.Println("Congratz!!!")
		fmt.Printf("Pass:\t%s\n", pass)
		fmt.Printf("Hash:\t%s\n", *hash)
		os.Exit(0)
	}
	wg.Done()
}
