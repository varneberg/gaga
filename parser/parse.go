package parser

import (
	"bufio"
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
	//fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
	return out
}

//func ParseTerraformPlan() {
//	s := "Plan: 9 to add, 0 to change, 0 to destroy."
//	match, err := regexp.MatchString("Plan:", s)
//	fmt.Println(match, err)
//
//tfplan := ReadPipeInput()
//fmt.Println(tfplan)

//toSlice := strings.Split(tfplan, "\n")
//fmt.Println(toSlice)
// fmt.Println(toSlice)
// 	// found := false

// 	// var out []string
// 	for i, line := range toSlice{
// 		// Separator used in Terraform plan
// 		if strings.Contains(line, "─────────────────────────────────────────────────────────────────────────────"){
// 			fmt.Println(toSlice[i+1:len(toSlice)-4])
// 			// fmt.Println(len(toSlice))
// 			// fmt.Println(toSlice[777])
// 			break
// 		}
// 	}
//}
