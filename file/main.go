package main

import "os"

func main() {
	/*file, err := os.Create("mangal.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, errr := io.WriteString(file, "Mangal is a popular Indian comic book character.\n")
	if errr != nil {
		fmt.Println("Error writing to file:", errr)
		return
	}
	*/

	/*file, err := os.Open("mangal.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
	*/

	content, err := os.ReadFile("mangal.txt")
	if err != nil {
		println("Error reading file:", err)
		return
	}
	println(string(content))
}
