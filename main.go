package main

import (
	"flag"
	"fmt"
	"os"
)

// Конфигурация приложения
// Пользовательские данные
var src string   // Полный путь к файлу, который нужно скопировать
var dst string   // Файл в который будет скопировано содержимое src
var offset int64 // Смещение в файле, с которого нужно начать копирование, по умолчанию 0
var limit int64  // Кол-во байт которые нужно скопировать в файл-получатель

func main() {
	var err error

	// Получаем пользовательские параметры
	// Не проводим валидацию, т.к. она будет в copyFile
	readUserParams()

	// Копируем файл
	// Т.к. в unix системах все - файл, то копировать можно не только файлы в их базовом понимании,
	// а например, диск или ввод/вывод на терминал
	err = copyFile(src, dst, offset, limit, progress)
	if err != nil {
		// В случае ошибки пишем в специально для этого предназначенный поток stderr
		_, _ = fmt.Fprintf(os.Stderr, "Copy error: %s\n", err)
		flag.Usage() // Выводим правила использования и выходим с кодом завершения 2
		os.Exit(2)
	}

	fmt.Println("\nSuccess!")
}

// readUserParams считывает параметры, переданные пользователем в качестве аргументов
func readUserParams() {
	flag.StringVar(&src, "from", "", "source file")
	flag.StringVar(&dst, "to", "", "destination file")
	flag.Int64Var(&offset, "offset", 0, "offset")
	flag.Int64Var(&limit, "limit", 0, "limit")

	flag.Parse()
}

// Выводит текущий прогресс копирования. В аргументе written необходимо передать кол-во записанных байт
// Конечно, можно было бы воспользоваться прогресс-баром вроде https://github.com/cheggaaa/pb
// Но пока по максимуму пишу сам, больше практики - больше знаний :)
func progress(written int64, filesize int64) {
	percents := int(float32(written) / float32(filesize) * 100)
	fmt.Printf("\rWritten %d of %d bytes - %d %%", written, filesize, percents)
}
