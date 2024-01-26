# How do we run the code in our project ?
```
use run main.go 
```

# GO CLI 
    - go run : compiles and execute one or two files
    - go build : only compiles a bunch of source code files
    - go fmt : formate all the code in each file in the current working directory 
    - go install : compiles and install a package 
    - go get : Download the raw source code of someone else's package 
    - go test : Runs any tests associtated with the current project 

 # What does package main mean
    - Type of packages
        1) Executable - Generate a file that we cam run 
            - Must have a func call 'main' (*Important*)
        2) Reusable - Code used as 'helpers' Good place to put resuable logic

    - How to classify that 
        - Specifically the keyword main in package is Executable package 
        - All other package names are reusable 

# What does import fmt means?
    - Standard library package for formatting in GO 
    - we use import statement to import other packages in our package 
    - example we have main package; we need to import otherpackage name fmt then we use import fmt in main pacjkage so basically it will create a link from our package main to package fmt.

# What is func in Go?
    - func is used to make a function in GO 
    - Syntax: func FunctionName() { FUnctionBody }
