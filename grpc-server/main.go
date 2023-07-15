package main

import (
	"context"
	"log"
	"net"

	invoicer "github.com/naqet/learning-go/grps-server/invoicer"
	"google.golang.org/grpc"
)

type myInterfaceServer struct {
    invoicer.UnimplementedInvoicerServer

}

func (s myInterfaceServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
    return &invoicer.CreateResponse{
         Pdf: []byte("test"),
         Docx: []byte("test"),
    }, nil;
}

func main() {
    listener, err := net.Listen("tcp", ":8080");

    if err != nil {
        log.Fatalf("Cannot create the listener: %s", err);
        return;
    }

    serverRegister := grpc.NewServer();
    service := &myInterfaceServer{};

    invoicer.RegisterInvoicerServer(serverRegister, service);
    
    err = serverRegister.Serve(listener);

    if err != nil {
        log.Fatalf("Cannot create the listener: %s", err);
        return;
    }
    
}
