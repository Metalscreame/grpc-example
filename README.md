To generate a pb file use (protoc bin must be installed):

    protoc --go_out=. filename.proto
    
To generate grpc service file

    protoc --go_out=plugins=grpc:package filename.proto    
    
Concrete case    
    
    protoc --go_out=plugins=grpc:chat chat.proto