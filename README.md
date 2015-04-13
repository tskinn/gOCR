# gOCR
This is an attempt to make a primitive self-organizing map that learns to recognize characters that are 9x9.  
The project is made up of a go web server that serves up static html and then connects via websockets and allows user to specify the details of the neuralnetwork.  

## How to use:  

### Prepare to install:  
 1. Install golang on your system.  
 2. Create a directory for your GOPATH environment variable  
example:  
<code>  mkdir $HOME/go  </code>  
<code>echo "export GOPATH=$HOME/go" >> $HOME/.bashrc</code>

### Install:  
 1. Install gorilla/sockets (GoCR uses websockets from this package)  
<code>
go get github.com/gorilla/sockets
</code>
 2. Install GoCR  
<code>
go get github.com/tskinn/GoCR
</code>

### Running GoCR:  
 1. Go to the GoCR source directory  
<code>
cd $HOME/go/src/github.com/tskinn/GoCR  
</code>
 2. Run the server and the neuralnetwork  
<code>
go run server.go neuralnet.go  
</code>
 3. Open a browser and go to <code>localhost:3000</code>