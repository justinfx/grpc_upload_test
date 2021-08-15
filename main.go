//go:generate ./generate_proto.sh
package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "grpc_file_test/uploader"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port       = ":50051"
	maxReceive = 500 * MB
)

const MB = 1 << 20

type server struct {
	pb.UnimplementedUploaderServer
}

func (s *server) Upload(ctx context.Context, req *pb.FileTransferRequest) (*empty.Empty, error) {
	var (
		header *pb.FileHeader
		buf    bytes.Buffer
	)

	start := time.Now()

	header = req.GetHeader()
	log.Printf("[Upload] Got file request header: %s\n", header.GetName())
	if header.FileSize != nil {
		log.Printf("  Reported file size should be: %d\n", header.GetFileSize())
	}

	if data := req.GetData(); data != nil {
		buf.Write(data)
	}

	took := time.Since(start)
	log.Printf("  Total bytes received: %d\n", buf.Len())
	log.Printf("[Upload] Took: %s\n", took)
	return &empty.Empty{}, nil
}

func (s *server) UploadStream(stream pb.Uploader_UploadStreamServer) error {
	var (
		header     *pb.FileHeader
		buf        bytes.Buffer
		chunkCount int
	)
	start := time.Now()
	for {
		req, err := stream.Recv()
		if err == io.EOF || stream.Context().Err() != nil {
			break
		}
		if err != nil {
			log.Println(err)
			stream.SendAndClose(&empty.Empty{})
			return err
		}

		if req.GetHeader() != nil {
			header = req.GetHeader()
			log.Printf("[UploadStream] Got file request header: %s\n", header.GetName())
			if header.FileSize != nil {
				log.Printf("  Reported file size should be: %d\n", header.GetFileSize())
			}
			continue
		}

		if req.GetChunk() != nil {
			chunkCount++
			buf.Write(req.GetChunk())
			//fmt.Printf("Got file data chunk #%d of size %d\n", chunkCount, n)
		}
	}
	stream.SendAndClose(&empty.Empty{})
	took := time.Since(start)
	log.Printf("  Total bytes received (in %d chunk(s)): %d\n", chunkCount, buf.Len())
	log.Printf("[UploadStream] Took: %s\n", took)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxReceive))
	pb.RegisterUploaderServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
