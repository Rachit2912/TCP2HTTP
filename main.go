package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(file io.ReadCloser) <-chan string{
	out := make(chan string,1)

	go func(){
		defer close(out)
		defer file.Close()
		
		str := "" 
		for{
			buffer := make([]byte,8)
			n, err := file.Read(buffer)
			if err != nil{
				break
			}
			
			buffer = buffer[:n]	
			if pos_n := bytes.IndexByte(buffer,'\n'); pos_n != -1 {
				str += string(buffer[:pos_n])
				fmt.Printf("%s\n",str)
				buffer = buffer[pos_n+1:]
				str=""	
			}
			str += string(buffer)
		}

		if len(str)!=0 {
			fmt.Printf("%s\n",str)
		}
	}()

	return out
}

func main() {

	file, err := os.Open("messages.txt")
	if err != nil{
		log.Fatal("error in reading the file",err);
	}

	lines := getLinesChannel(file)
	for line := range lines{
		fmt.Printf("%s\n",line)
	}

	 

	
		







	





}


