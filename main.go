package main

import (
	"fmt"
	"./utils"
	"os"
)

func main() {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("--------------------------Менеджер работы приложения--------------------------")
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("-------------------------Примерный порядок выполнения-------------------------")
	fmt.Println("При первом запуске необходимо загрузить данные через сеть с помощью 1 команды")
	fmt.Println("Далее нужно разархивировать полученные данные с помощью команды 2")
	fmt.Println("Данные помещаются в папку output, из которой дальше производится считывание и запись")
	fmt.Println("в базу данных посредством команды 3, после чего успешно загруженные файлы помещаются в")
	fmt.Println("папку finished")
	fmt.Println("После этого можно закрыть программу сочетанием клавиш ctrl+c или командой 4")
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("---------------------------Список доступных комманд---------------------------")
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("1. Загрузить данные ")
	fmt.Println("2. Провести разархивацию файлов")
	fmt.Println("3. Запись в базу данных")
	fmt.Println("4. Выход")
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println()
	var i string
	for {
		fmt.Scanln(&i)
		switch i {
		case "1":
			utils.DownLoad()
			continue
		case "2":
			utils.Unzip()
			continue
		case "3":
			utils.ToDB()
			continue
		case "4":
			os.Exit(0)
		}

	}
}
