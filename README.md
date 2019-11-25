# test-in-go
Project with an example of unit test in go  

The conventional way of doing unit tests in Go lang, is to leave them at the same level of the source code and executing them separately, in this example a specialized folder is made for the tests, and through a shell-script file all the tests are centralized, making it easier to execute and read results.

The objective of this project is to do the tests in a more orderly way, allowing to make a productive and maintainable code.

To understand a little about how unit tests work in go you can read the following document:

https://docs.google.com/document/d/1veGEbFmQ-e1dZvASFF1wOZTTAcYV0nN9AC3yQa3yQ_s/edit?usp=sharing

To run all the tests simultaneously, you can run the shell-script file located in the test folder called "coverage.sh", its execution requires as a parameter the path of the project you want to test, the execution can be done as follows in a UNIX console:

sh ./coverage.sh ~/go/src/<project-name>
