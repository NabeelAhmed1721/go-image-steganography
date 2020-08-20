package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// fmt.Println(len(getRawFileData("example.jpg")))
	// writeRawFileData("nabeel.ahmed", []byte("NABEEL WAS HERE"))
	adder := []byte("Nabeel Ahmed")
	writeRawFileData("example.jpg", (append(getRawFileData("example.jpg"), adder...)))
}

func writeRawFileData(filename string, data []byte) int {
	output, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	bytes, writeErr := output.Write(data)
	if writeErr != nil {
		log.Fatal(writeErr)
	}

	return bytes
}

func getRawFileData(filename string) []byte {
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// Buffer Size
	stat, staterr := input.Stat()
	if staterr != nil {
		log.Fatal(err)
	}

	size := stat.Size()
	byteBuffer := make([]byte, size)

	// Byte Buffer
	buffer := bufio.NewReader(input)
	buffer.Read(byteBuffer)

	return byteBuffer
}
