package parser

import (
	"bufio"
	"io"
	"log"
	"os"
)

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
	//fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
	return out
}
