package main

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/lxc/go-lxc.v2"
)

var (
	lxcpath string
)

func init() {
	flag.StringVar(&lxcpath, "lxcpath", lxc.DefaultConfigPath(), "Use specified container path")
	flag.Parse()
}
func main() {
	c := lxc.DefinedContainers(lxcpath)

	var buffer bytes.Buffer
	for i := range c {
		buffer.WriteString(fmt.Sprintf("%s\t%s\n", c[i].Name(), c[i].State()))
	}

	fmt.Println(buffer.String())
}
