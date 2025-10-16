package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// fn. for reading the lines from channel :
func getLinesChannel(file io.ReadCloser) <-chan string{
	// making a channel of string : 
	out := make(chan string,1)

	// goroutine : 
	go func(){
		defer close(out)
		defer file.Close()
		
		str := "" 
		for{
			// making & writing from file into buffer : 
			buffer := make([]byte,8)
			n, err := file.Read(buffer)
			if err != nil{
				break
			}
			
			buffer = buffer[:n]	
			// searching for a new line character ('\n'):
			// if found an index, break it into chunks there : 
			if pos_n := bytes.IndexByte(buffer,'\n'); pos_n != -1 {
				str += string(buffer[:pos_n])
				fmt.Printf("%s\n",str)
				buffer = buffer[pos_n+1:]
				str=""	
			}
			// adding to the string :
			str += string(buffer)
		}

		// for last part of string if left any : 
		if len(str)!=0 {
			fmt.Printf("%s\n",str)
		}
	}()

	// returning the channel : 
	return out
}

func main() {

	// reading from the file : 
	file, err := os.Open("messages.txt")
	if err != nil{
		log.Fatal("error in reading the file",err);
	}

	// getting lines from channel : 
	lines := getLinesChannel(file)
	for line := range lines{
		// printing each line :
		fmt.Printf("%s\n",line)
	}

	 

	
		







	





}


