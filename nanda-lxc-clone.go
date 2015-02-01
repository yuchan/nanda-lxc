// Copyright © 2013, 2014, The Go-LXC Authors. All rights reserved.
// Use of this source code is governed by a LGPLv2.1
// license that can be found in the LICENSE file.
// +build linux,cgo

// Copyright © 2015, Yusuke Ohashi. All rights reserved.

package main

import (
	"flag"
	"gopkg.in/lxc/go-lxc.v2"
	"log"
)

var (
	lxcpath string
	name    string
	newname string
	backend lxc.BackendStore
)

func init() {
	flag.StringVar(&lxcpath, "lxcpath", lxc.DefaultConfigPath(), "Use specified container path")
	flag.StringVar(&name, "name", "rubik", "Name of the original container")
	flag.StringVar(&newname, "newname", "rubik_new", "New Name of the cloned container")
	flag.Var(&backend, "backend", "Backend type to use, possible values are [dir, zfs, btrfs, lvm, aufs, overlayfs, loopback, best]")
	flag.Parse()
}

func main() {
	c, err := lxc.NewContainer(name, lxcpath)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	if backend == 0 {
		log.Fatalf("ERROR: %s\n", lxc.ErrUnknownBackendStore)
	}

	log.Printf("Cloning the container using %s backend...\n", backend)
	err = c.Clone(newname, lxc.CloneOptions{
		Backend: backend,
	})
	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}
}
