package main

import (
	"os"
	"log"
	"fmt"
	"bytes"
	"io"
)

func getLinesChannel(file io.ReadCloser ) <-chan string{
	out := make(chan string,1)
	go func(){
		defer file.Close()
		defer close(out)
		
			str := ""	
		for {
			data := make([]byte,8)
			n, err := file.Read(data)

			if err!= nil {
				break
			}
			data = data[:n]
			
			if i := bytes.IndexByte(data, '\n'); i!=-1{
				str += string(data[:i])
				data = data[i+1:]
				out <- str
				str = ""
			}
			str += string(data)
		}
		if len(str) != 0{
			out <- str 
		}
	}()
	return out
}

func main() {
	file, err := os.Open("messages.txt")

	if err != nil {
		log.Fatal("error",err)
	}
	lines := getLinesChannel(file)
	for line := range lines{
		fmt.Printf("read: %s \n", line)
	}

}
