//Copyright(c) 2025 a42-good
//本程序是自由软件：你可以再分发之和/或依照由自由软件基金会发布的 GNU 通用公共许可证修改之，无论是版本 3 许可证，还是任何以后版都可以。

//发布该程序是希望它能有用，但是并无保障;甚至连可销售和符合某个特定的目的都不保证。请参看 GNU 通用公共许可证，了解详情。

//你应该随程序获得一份 GNU 通用公共许可证的复本。如果没有，请看 <https://www.gnu.org/licenses/>。 






package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main () {
	var fileName string
	fmt.Print("fileName:")
	fmt.Scanln(&fileName)
	
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
	
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("NO",err)
		return
	}
	defer file.Close()
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("exit结束编辑")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("NO", err)
			return
		}
		line = line[:len(line)-1]
		if line == "exit" {
			break
		}
		_, err = file.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("NO", err)
			return
		}
	}

	fmt.Println("Ok")	
}


