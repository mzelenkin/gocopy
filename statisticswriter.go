package main

import "io"

// StatisticsWriter обертка для извлечение статистики записанных байт
type StatisticsWriter struct {
	writer   io.Writer
	written  int64
	callback func(int64)
}

// NewStatisticsWriter - конструктор экземпляра StatisticsWriter
func NewStatisticsWriter(writer io.Writer, callback func(int64)) *StatisticsWriter {
	return &StatisticsWriter{writer: writer, callback: callback}
}

// Write передает запрос writer'у, после выполнения запроса сохраняет статистику
// и вызывает callback функцию
func (pw *StatisticsWriter) Write(p []byte) (n int, err error) {
	n, err = pw.writer.Write(p)
	pw.written += int64(n)
	pw.callback(pw.written)

	return
}
