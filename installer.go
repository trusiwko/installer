package main

import (
    //"os"
	"os/exec"
	"io"
	"fmt"
	"log"
	"bufio"
)

func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}

func read(reader io.ReadCloser) {
	nBytes, nChunks := int64(0), int64(0)
    r := bufio.NewReader(reader)
    buf := make([]byte, 0, 4*1024)
    for {
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
        if n == 0 {
            if err == nil {
                continue
            }
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
        nChunks++
        nBytes += int64(len(buf))
        // process buf
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
		fmt.Print(">>>", string(buf[:]))
    }
	fmt.Println("end")
}

func main() {
    cmd := exec.Command("sqlplus", "/nolog")

    // Create stdout, stderr streams of type io.Reader
    stdout, err := cmd.StdoutPipe()
    checkError(err)
    //stderr, err := cmd.StderrPipe()
    //checkError(err)

    // Start command
    err = cmd.Start()
    checkError(err)

    // Don't let main() exit before our command has finished running
    defer cmd.Wait()  // Doesn't block

    // Non-blockingly echo command output to terminal
	
	read(stdout)
    // io.Copy(os.Stdout, stdout)
    // io.Copy(os.Stderr, stderr)

    // I love Go's trivial concurrency :-D
    fmt.Printf("Do other stuff here! No need to wait.\n\n")
}