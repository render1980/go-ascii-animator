package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var frameDelay = flag.Float64("d", 0.5, "Frame delay")
	var file = flag.String("f", "", "File with ASCII graphics to animate")
	var loop = flag.Bool("l", false, "If the animation should loop")

	flag.Parse()

	if len(*file) == 0 {
		fmt.Println("You should define file! (use -f)")
		os.Exit(1)
	}

	// fmt.Println("frameDelay:", *frameDelay)
	//fmt.Println("file:", *file)
	//fmt.Println("loop:", *loop)

	process(*frameDelay, *file, *loop)
}

func process(frameDelay float64, file string, loop bool) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Read file line error: %v", err)
			return
		}
		fmt.Printf(line)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
