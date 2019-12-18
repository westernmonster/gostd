package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// func MultiReader(readers ...Reader) Reader
//
// MultiReader 返回一个 Reader 串联参数中所有的 Reader
// 当所有Reader读取完毕后 返回 EOF，如果某一个Reader返回
// 非EOF的错误，MultiReader将返回该错误
func DoMultiReader() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	rAll := io.MultiReader(r1, r2, r3)

	buf2 := make([]byte, 10)
	for {
		n, err := rAll.Read(buf2)
		fmt.Println(n, err, string(buf2[:n]))
		if err == io.EOF {
			break
		}
	}
}

// io.Reader
// 1. 读取 n 个byte 到 p 中
// 2. 会将 p 作为暂存空间
// 3. 如果读到末尾，会返回 io.EOF 错误

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
func DoReader() {
	r := strings.NewReader("abcde")
	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
}

// func LimitReader(r Reader, n int64) Reader
//
// LimitReader 返回一个 Reader 读取 n 个 bytes 后返回 EOF
func DoLimitReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	buf := make([]byte, 4)
	for {
		n, err := lr.Read(buf)
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
}

// func TeeReader(r Reader, w Writer) Reader
//
// TeeReader 返回一个 Reader，所有读取的内容将同时写入到 Writer 中，读取必须在写入中
func DoTeeReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(&buf)
	printall(tee)
	printall(&buf)
}

func main() {
	DoReader()

	fmt.Println()
	DoMultiReader()

	fmt.Println()
	DoLimitReader()

	fmt.Println()
	DoTeeReader()
}
