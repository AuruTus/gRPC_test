package test

import (
	"fmt"
)

// public static int hashCode(byte[] value) {
// 	int h = 0;
// 	for (byte v : value) {
// 		h = 31 * h + (v & 0xff);
// 	}
// 	return h;
// }

func LatinHashCode(str string) int {
	h := 0
	for _, v := range str {
		h = 31 * h + (int(v) & 0xff)
	}
	return h
}

// static final int hash(Object key) {
// 	int h;
// 	return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);
// }

func Hash(str string) int {
	h := LatinHashCode(str)
	return h ^ int(uint(h) >> 16)
}

func Test() {
	fmt.Printf("1234")
}
