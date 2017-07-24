/*
  defer
 */

package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	dst := "f:/dst.txt"
	src := "f:/vfs/upload1.log"
	CopyFile(dst, src)
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}