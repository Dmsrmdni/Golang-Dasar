package main

import "os"

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)
	return nil
}

// func readFile(name string) (string, error) {
// 	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

// 	if err != nil {
// 		return "", err
// 	}

// 	defer file.Close()

// 	render := bufio.NewReader(file)

// 	var message string

// 	for {
// 		line, _, err := render.ReadLine()
// 		if err == io.EOF {
// 			break
// 		}

// 		message += string(line) + "\n"
// 	}

// 	return message, nil
// }

// func addToFile(name string, message string) error {
// 	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	file.WriteString(message)
// 	return nil
// }

// func deleteFile() {
// 	var err = os.Remove("file_package/sample.log")
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println("file berhasil di delete")
// }

func main() {
	createNewFile("file_package/sample.log", "Sample logs")

	// result, _ := readFile("file_package/sample.log")
	// fmt.Println(result)

	// addToFile("file_package/sample.log", "\nSample logs")

	// deleteFile()

}
