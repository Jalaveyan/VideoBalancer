package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
	"os"
	"regexp"
	"sync/atomic"

	pb "VideoBalancer/balancer/proto/package/balancer"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedBalancerServer
	counter uint64
	cdnHost string
}

func (s *server) Redirect(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	videoURL := req.Video
	s1 := extractServer(videoURL)           // функция для извлечения s1 из URL
	videoPath := extractVideoPath(videoURL) // функция для извлечения пути видео

	if atomic.AddUint64(&s.counter, 1)%10 == 0 {
		// Каждый 10й запрос отправляем на оригинальный сервер
		return &pb.Response{Url: videoURL}, nil
	}

	// Если cdnHost пуст, используем оригинальный URL
	if s.cdnHost == "" {
		return &pb.Response{Url: videoURL}, nil
	}

	// Все остальные запросы отправляем на CDN
	cdnURL := fmt.Sprintf("http://%s/%s/%s", s.cdnHost, s1, videoPath)
	return &pb.Response{Url: cdnURL}, nil
}

func extractServer(videoURL string) string {
	// Используем регулярное выражение для извлечения поддомена (s1, s2, ...)
	re := regexp.MustCompile(`http://(s\d+)\.origin-cluster`)
	match := re.FindStringSubmatch(videoURL)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func extractVideoPath(videoURL string) string {
	// Используем регулярное выражение для извлечения пути к видео
	re := regexp.MustCompile(`http://s\d+\.origin-cluster/(.*)`)
	match := re.FindStringSubmatch(videoURL)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func main() {
	cdnHost := os.Getenv("CDN_HOST")
	cdnHost = "storage.googleapis.com"
	if cdnHost == "" {
		log.Printf("CDN_HOST environment variable is not set. Using default CDN host: storage.googleapis.com")
	} else {
		log.Printf("Using CDN host: %s", cdnHost)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterBalancerServer(grpcServer, &server{cdnHost: cdnHost})

	slog.Debug("Server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("failed to serve: %v", err)
	}
}
