package main

import (
	"io"
	"os"
	//"fmt"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (rot *rot13Reader) Read(bytes []byte) (n int,err error) {
	n, err = rot.r.Read(bytes)
	for i:=0;i<len(bytes);i++ {
		ch:=bytes[i]
		if ch>=65 && ch<=77 {
			bytes[i]+=13	
		} else if ch>=97 && ch<=109 {
			bytes[i]+=13
		} else if ch>=78 && ch<=90 {
			bytes[i]-=13
		} else if ch>=110 && ch<=122 {
			bytes[i]-=13
		}
		//fmt.Printf("b[:n] = %q\n", b[:n])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	//fmt.Printf("%v",r.r)
}
