# sb_golang

Golang sandbox for variety of templates, and tutorial snippets.

This repository could be a useful starting point for writing Go code, or demoing how to do it.
It could be used on its own, or deployed on a virtual machine within the workspace [https://github.com/panchul/workspace](https://github.com/panchul/workspace)

The ```/doc``` folder of the ```workspace``` repo has the documentation. 

I will gradually migrate the notes I have to this repository. To keep track of what I am adding:

+ hello
    - Demoes how to create and run a Go project. More of a validation that Go was installed properly.
      (To run, once you follow the Go installation, and go to $GOPATH, type ```$ go install github.com/user/sb_golang/hello```
       and run ```$ $GOPATH/bin/hello```. You'd need to move that whole folder to the proper location, or make a
       symbolic link to ```/home/vagrant/golang_workspace/src/github.com/vagrant/sb_golang``` )
       ```$ go install github.com/vagrant/sb_golang/hello``` - to build and install
       ```$ go test github.com/vagrant/sb_golang/stringutil``` - to run tests
       ```$ $GOPATH/bin/hello``` - to execute (or add the path to the PATH)
       
