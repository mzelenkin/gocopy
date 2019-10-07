package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Копирует файл srcPath в dstPath
func copyFile(srcPath string, dstPath string, offset int64, limit int64, progressCallback func(int64, int64)) (err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error opening source: %v", err)
	}
	defer src.Close()

	dst, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening destination: %v", err)
	}
	defer dst.Close()

	// Пользуемся тем, что src у нас типа *File
	// и определяем размер файла через Seek
	srcSize, err := src.Seek(0, io.SeekEnd)
	if err != nil {
		return fmt.Errorf("filesize determination error: %v", err)
	}

	// Проверяем offset на границы
	if offset < 0 || offset >= srcSize {
		return errors.New("offers must be 0...filesize -1")
	}

	// Проверяем limit на границы
	maxLimit := srcSize - offset
	if limit < 0 || limit >= maxLimit {
		return fmt.Errorf("limit must be 0...filesize - offset (%d)", maxLimit)
	}

	// Выставляем позицию в файле с которой начнем копирование
	_, err = src.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("cannot set position in source: %v", err)
	}

	// Переменная с количеством копируемых байт
	var fileSize int64

	if limit > 0 {
		fileSize = limit
	} else {
		fileSize = srcSize - offset
	}
	// Создаем writer, который будет собирать статистику
	statWriter := NewStatisticsWriter(dst, func(pos int64) {
		progressCallback(pos, fileSize)
	})

	if limit == 0 {
		_, err = io.Copy(statWriter, src)
	} else {
		reader := io.LimitedReader{R: src, N: limit}
		_, err = io.Copy(statWriter, &reader)
	}

	return
}
