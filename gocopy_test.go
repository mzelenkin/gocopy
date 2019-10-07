package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const workingDir = "/tmp"

// Проверка ошибки при некорретном смещении
func TestInvalidSource(t *testing.T) {
	srcFilename := makeFullpath("/testrandom")
	dstFilename := makeFullpath("/testrandomcopy")

	err := copyFile(srcFilename, dstFilename, 0, 0, nil)
	if err == nil {
		t.Fail()
	}

	_ = os.Remove(srcFilename)
	_ = os.Remove(dstFilename)
}

// Проверка ошибки при некорретном смещении
func TestInvalidOffset(t *testing.T) {
	srcFilename := makeFullpath("/test1")
	dstFilename := makeFullpath("/test1copy")
	data := []byte("12345") // 5 Байт
	err := ioutil.WriteFile(srcFilename, data, 0644)
	if err != nil {
		t.Errorf("Cannot create src file %s", srcFilename)
	}

	err = copyFile(srcFilename, dstFilename, 20, 0, nil)
	if err == nil {
		t.Fail()
	}

	err = copyFile(srcFilename, dstFilename, 0, 20, nil)
	if err == nil {
		t.Fail()
	}

	_ = os.Remove(srcFilename)
	_ = os.Remove(dstFilename)
}

func makeFullpath(filename string) string {
	return path.Join(workingDir, filename)
}
