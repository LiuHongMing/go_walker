/**
  常量
 */

package main

const Pi = 3.14159265358979323846
const zero = 0.0

const (
	size int64 = 1024
	eof = -1
)

const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"

const mask = 1 << 3

func main() {
	println(a, b, mask)
}
