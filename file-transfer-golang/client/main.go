package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	// Nama file yang akan dikirim
	filePath := "file_to_send.txt"

	// Buka file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Dapatkan informasi file (ukuran)
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()

	// Hubungkan ke server di port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)

	// Kirim nama file diikuti dengan newline
	_, err = writer.WriteString(fileName + "\n")
	if err != nil {
		fmt.Println("Error sending file name:", err)
		return
	}

	// Kirim ukuran file diikuti dengan newline
	_, err = writer.WriteString(strconv.FormatInt(fileSize, 10) + "\n")
	if err != nil {
		fmt.Println("Error sending file size:", err)
		return
	}
	writer.Flush()

	// Kirim konten file
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}

	fmt.Println("File sent successfully!")
}
