package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var frameDelay = flag.Int("d", 1000, "Delay between frames (milliseconds)")
	var filename = flag.String("f", "", "File with ASCII graphics to animate")
	var loop = flag.Bool("l", false, "If the animation should loop")

	flag.Parse()

	checkFileName(*filename)
	process(time.Duration(*frameDelay), *filename, *loop)
}

func process(frameDelay time.Duration, file string, loop bool) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	frames := readFrames(reader)
	if frames == nil {
		log.Fatalf("Read file line error: %v", err)
	}
	animate(frames, frameDelay, loop)
}

func readFrames(reader *bufio.Reader) *FrameBuffer {
	frame := ""
	var frames FrameBuffer

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		if line == "[yaaa]\n" || line == "[end]\n" {
			frames.frames = append(frames.frames, frame)
			frame = ""
		} else {
			frame = frame + line
		}
	}
	return &frames
}

func animate(frames *FrameBuffer, frameDelay time.Duration, loop bool) {
	if len(frames.frames) < 1 {
		log.Fatalf("File is empty or has incorrect format!")
	}
	for {
		for _, s := range frames.frames {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			fmt.Println(s)
			time.Sleep(frameDelay * time.Millisecond)
		}
		if loop == false {
			break
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileName(filename string) {
	if len(filename) == 0 {
		fmt.Println("You should define file! (use -f)")
		os.Exit(1)
	}
}

/**
 * Sequence of frames to animate.
 * Frames go in sequence.
**/
type FrameBuffer struct {
	frames []string
}
