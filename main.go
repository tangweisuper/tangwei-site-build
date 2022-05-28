package main

import "tangwei-site-build/builder/build"

func main() {
	err := build.Build()
	if err != nil {
		println(err)
		return
	}
}
