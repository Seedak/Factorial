# Factorial

The following project was developed using Goland IDE version 2019.1.3 and requires Go version 1.6 or above, Protocol Buffers V3, grpc to execute. And Intellij Idea Ultimate and requires gradle version 3.0+ to run client grpc with go grpc server.

The program performs the following tasks using grpc:
1. Calculate average of a number.
2. Calculate Factorial of a number.
3. Calculate prime factors of a number.
4. Calculate average of a given set of numbers.
5. Determine whether given numbers are odd or even.

*****To run the program with Go server and Go client*****
1. Open Factorial directory as project in GoLand IDE.
2. Run main.go file

Note:-> Before running the program please chech the imports in the following go files first
        1.) client.go in Factorial -> client directory
        2.) server.go and calculator_handler in Factorial -> server directory
        
*****Output*****


![Screenshot from 2019-06-05 12-16-16](https://user-images.githubusercontent.com/36620386/58937270-91574700-878f-11e9-8ffa-bf74cb4d3561.png)




*****To run the program with Go server and Java client*****
1. Open Factorial directory as project in Goland IDE.
2. Run only grpc server from main.go file 
3. Open grpcjava directory as project in IntelliJ IDE.
4. After opening the project, open the project directory in terminal and run the following command 
    "gradle clean build"
5. After running gradle build, in grpcjava navigate to src/main/java and run client.java

*****Output*****

![Screenshot from 2019-06-07 12-46-53](https://user-images.githubusercontent.com/36620386/59087491-63563c00-8922-11e9-84a6-10f2820899f3.png)





