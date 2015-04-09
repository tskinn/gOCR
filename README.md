# gOCR
This is an attempt to make a primitive self-organizing map that learns to recognize characters that are 9x9.  
The project is made up of a go web server that serves up static html and then connects via websockets and allows user to specify the details of the neuralnetwork.  

How to:  

Before installation:  
 - Install golang on your system.  
 - Create a directory for your GOPATH environment variable  
   example:  
<code>  
mkdir $HOME/go  
</code>  
<code>
echo "export GOPATH=$HOME/go" >> $HOME/.bashrc
</code>

Installation:  
 - Install gorilla/sockets (GoCR uses websockets from this package)  
<code>
go get github.com/gorilla/sockets
</code>
 - Install GoCR  
<code>
go get github.com/tskinn/GoCR
</code>

Running GoCR:  
 - Go to the GoCR source directory  
<code>
cd $HOME/go/src/github.com/tskinn/GoCR  
</code>
 - Run the server and the neuralnetwork  
<code>
go run server.go neuralnet.go  
</code>
 - Open a browser and go to <code>localhost:3000</code>