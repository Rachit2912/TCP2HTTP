package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
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
			if err != nil {
				if err != io.EOF {
					log.Println("read error:", err)
				}
				break
			}
			
			buffer = buffer[:n]	
			// searching for a new line character ('\n'):
			// if found an index, break it into chunks there : 
			if pos_n := bytes.IndexByte(buffer,'\n'); pos_n != -1 {
				str += string(buffer[:pos_n])
				out <- str
				// fmt.Printf("%s\n",str)
				buffer = buffer[pos_n+1:]
				str=""	
			}
			// adding to the string :
			str += string(buffer)
		}

		// for last part of string if left any : 
		if len(str)!=0 {
			out <- str
			// fmt.Printf("%s\n",str)
		}
	}()

	// returning the channel : 
	return out
}

func main() {

	// reading onto the TCP connection : 
	listener, err := net.Listen("tcp",":42069");
	if err != nil{
		log.Fatal("error in listening the tcp connection : ",err);
	}
	fmt.Println("Listening on port 42069...")


	
	// getting lines from channel : 
	for{
		// accepting the tcp connections : 
		conn, err := listener.Accept();
		if err != nil{
			log.Println("error in accepting the tcp connection : ",err);
		}
		// reading lines from each tcp connection : 
		fmt.Println("a connection is established");
		for line := range getLinesChannel(conn){
			fmt.Printf("%s\n",line);
		}
	}

	






	





}


