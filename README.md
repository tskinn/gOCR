# gOCR
An attempt to make a primitive self-organizing map that learns to recognize characters that are 9x9. 
The project is made up of a go web server that serves up static html and then connects via websockets and allows user to specify the details of the neuralnetwork.

How to:

Before installation:
 - Install golang on your system.
 - Create a directory for your GOPATH environment variable
   example:
        mkdir $HOME/go
        echo "export GOPATH=$HOME/go" >> $HOME/.bashrc

Installation:
 - Install gorilla/sockets (GoCR uses websockets from this package)
   go get github.com/gorilla/sockets
 - Install GoCR
   go get github.com/tskinn/GoCR

Running GoCR:
 - Go to the GoCR source directory
   cd $HOME/go/src/github.com/tskinn/GoCR
 - Run the server and the neuralnetwork
   go run server.go neuralnet.go