package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Defer_demo() {
	if err := write("readme.txt", "hello world"); err != nil {
		log.Fatal("Failed to write file due to ", err)
	}

	if err := fileCopy("readme.txt", "dst.txt"); err != nil {
		log.Fatal("Failed to copy due to : ", err)
	}
}

func write(fileName, message string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, message)
	if err != nil {
		return err
	}
	return nil
}

func fileCopy(source, destination string) error {
	src, err := os.Open(source)

	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(destination)

	if err != nil {
		return err
	}

	defer dst.Close()

	n, err := io.Copy(dst, src)

	if err != nil {
		return err
	}

	fmt.Printf("Copied %d from %s to %s\n", n, source, destination)

	if err := src.Close(); err != nil {
		return err
	}
	return dst.Close()
}

func Multiple_defers() {
	defer fmt.Println("Hello")
	defer fmt.Println("Bye")
	defer fmt.Println("World")
}
