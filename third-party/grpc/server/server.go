package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"grpc/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

type R struct {
	req *http.Request
	rw  http.ResponseWriter

	closed chan int
}

var reqCh = make(chan R)

type Server struct {
	api.UnimplementedRsLocaldServer
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r := R{
		req,
		rw,
		make(chan int),
	}
	reqCh <- r
	<-r.closed
}

func (s *Server) Login(ctx context.Context, body *api.LoginBody) (*api.LoginReply, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Server) Listen(server api.RsLocald_ListenServer) error {
	for req := range reqCh {
		buf := bytes.NewBuffer([]byte(""))
		req.req.WriteProxy(buf)
		fmt.Println("req:", string(buf.String()))
		_ = server.Send(&api.ProxyRequest{
			Data: buf.Bytes(),
		})

		var resp api.ProxyResponse
		if err := server.RecvMsg(&resp); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(string(resp.Data))
		_, err := io.Copy(req.rw, bytes.NewReader(resp.Data))
		close(req.closed)
		fmt.Println(err)
	}

	return nil
}

func (s *Server) SendResponse(ctx context.Context, response *api.ProxyResponse) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
