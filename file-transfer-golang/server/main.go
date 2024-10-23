package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// Membuat listener di port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		// Menerima koneksi dari client
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Menerima nama file
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving file name:", err)
		return
	}
	fileName = fileName[:len(fileName)-1] // Menghapus newline character

	// Menerima ukuran file
	fileSizeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving file size:", err)
		return
	}
	fileSizeStr = fileSizeStr[:len(fileSizeStr)-1]
	fileSize, err := strconv.ParseInt(fileSizeStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid file size:", err)
		return
	}

	// Membuat folder untuk menyimpan file
	saveDir := "received_files"
	os.MkdirAll(saveDir, os.ModePerm)

	// Membuat file di folder
	filePath := filepath.Join(saveDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Menerima konten file
	n, err := io.CopyN(file, conn, fileSize)
	if err != nil {
		fmt.Println("Error receiving file data:", err)
		return
	}
	fmt.Printf("File %s received, %d bytes written\n", fileName, n)
}
