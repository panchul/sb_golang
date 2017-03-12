# sb_golang

Golang sandbox for variety of templates, and tutorial snippets.

This repository could be a useful starting point for writing Go code, or demoing how to do it.
It could be used on its own, or deployed on a virtual machine within the workspace [https://github.com/panchul/workspace](https://github.com/panchul/workspace)

The ```/doc``` folder of the ```workspace``` repo has the documentation. 

I will gradually migrate the notes I have to this repository. To keep track of what I am adding:

+ hello
    - Set ```$GOPATH``` to the repo's directory (where you cloned it)
    - run, for example, ```go run src/hello/hello.go```:
        
    ```
    $ go run src/hello/hello.go 
    Hello, world!
    !dlrow ,olleH
    ```

    - or tests for the package ```go test stringutil```:
        
    ```
    $ go test stringutil 
    ok  	stringutil	0.005s
    ```
    
    - or you can install it, and run from $GOPATH/bin or directly if you add that to $PATH:
    
    ```$ go install hello 
      $ $GOPATH/bin/hello
      Hello, world!
      !dlrow ,olleH
     ```
       
