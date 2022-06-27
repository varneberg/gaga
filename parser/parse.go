package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Read unix pipe
func ReadPipeInput() string {
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	var out string
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
		//fmt.Println(string(buf))
		out += string(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	return out
}

func WritePipeOutput(plan string) {
	fmt.Println(plan)

}

func ReadFileInput(filename string) string{
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error openening file")
	}
	defer file.Close()

	var out string
	const maxS = 4
	buf := make([]byte, maxS)
	for  {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		out += string(buf[:n])
	}
	return out
}
