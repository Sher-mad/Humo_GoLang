package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

func main() {
	// filename := "my_file.txt"

	// count, err := CountCharters(filename)
	// if err != nil {
	// 	fmt.Println("Ошибка при чтении файла:", err)
	// 	return
	// }

	// fmt.Printf("В файле %s %d символов\n", filename, count)
}

// func CountCharters(filename string) (int, error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer f.Close()

// 	buf := make([]byte, 0, 1024)
// 	_, err = io.Copy(buf, f)
// 	if err != nil {
// 		return 0, err
// 	}

// 	str := string(buf)
// 	return len(str), nil
// }
