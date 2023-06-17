package main

import "tangwei-site-build/builder/build"

func main() {
	//err := build.Build24()
	err := build.BuildFlaskTemplate()
	if err != nil {
		println(err)
		return
	}
}
