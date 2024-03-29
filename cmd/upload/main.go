package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	gm "github.com/log-agent/emqx-grpc/grpc/proto"
	"github.com/snowlyg/helper/dir"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Buildtime string
	Version   string
)

var uname = flag.String("uname", "", "登录用户名；废弃")
var pwd = flag.String("pwd", "", "登录密码；废弃")
var path = flag.String("path", "", "文件地址")
var host = flag.String("host", "10.0.0.35:18088", "grpc地址")
var serName = flag.String("serName", "", "服务名称")
var uploadType = flag.String("uploadType", "", "更新类型")
var directionType = flag.String("directionType", "0", "门旁类型:1竖屏,0横屏")

func main() {
	flag.Parse()
	if *path == "" || *host == "" || *uploadType == "" {
		log.Println("path:", *path)
		log.Println("uploadType:", *uploadType)
		log.Println("directionType:", *directionType)
		log.Println("path,host, uploadType is required")
		os.Exit(1)
	}

	hostFlag := *host
	nameFlag := *serName
	if nameFlag == "" {
		var err error
		if nameFlag, err = getSerName(hostFlag); err != nil {
			log.Println("grpc.Dial", err.Error())
			os.Exit(1)
		} else if nameFlag == "" {
			log.Println("SerName is required")
			os.Exit(1)
		}

		hostFlag = "10.0.0.35:18088"
		if nameFlag == "doisWeb" {
			hostFlag = "10.0.0.23:18088"
		}
	}

	if !strings.Contains(hostFlag, ":18088") {
		log.Println("host must be x.x.x.x:18088", hostFlag)
		os.Exit(1)
	}

	log.Println("GRPC:", hostFlag)

	fileName := filepath.Base(*path)
	conn, err := grpc.Dial(hostFlag, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("grpc.Dial", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c := gm.NewGrpcManagerClient(conn)
	stream, err := c.UploadFile(ctx)
	if err != nil {
		log.Println("UploadFile", err.Error())
		os.Exit(1)
	}

	f, err := os.Open(*path)
	if err != nil {
		log.Println("os.Open", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Println("f.Stat", err.Error())
		os.Exit(1)
	}

	var size int64 = 500 * 1024 * 1024
	if stat.Size() > size {
		fmt.Printf("文件大小 [%d] 大于 [%d] 字节\n", stat.Size(), size)
		os.Exit(1)
	}

	md5, err := dir.MD5(*path)
	if err != nil {
		log.Println("MD5", err.Error())
		os.Exit(1)
	}

	// Maximum 1KB size per stream.
	buf := make([]byte, size)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println("Read", err.Error())
				return
			}
		}

		req := &gm.UploadRequest{
			FileName:  fileName,
			FileType:  *uploadType,
			SerName:   nameFlag,
			FileMd5:   md5,
			FileChunk: buf[:n],
		}

		if err := stream.Send(req); err != nil {
			log.Println("Send", err.Error())
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("CloseAndRecv", err.Error())
		os.Exit(1)
	}
	fmt.Println(reply)
	if reply.Status != 200 {
		log.Println(reply.Message)
		os.Exit(1)
	}
}

var (
	gRC     *Grpcurl
	grpcUri string
)

// SingleGrc
func SingleGrc(host ...string) *Grpcurl {
	if host != nil && host[0] != "" {
		grpcUri = host[0]
	}
	once := sync.Once{}
	once.Do(getGrc)
	return gRC
}

// getGrc
func getGrc() {
	gRC = &Grpcurl{}
	gRC.Host = grpcUri
	if err := gRC.Init(); err != nil {
		log.Println("grpc.Init", err.Error())
		return
	}
}

func getSerName(host string) (string, error) {
	uri := "10.0.0.35:18088"
	if host == "http://10.0.0.23" {
		uri = "10.0.0.23:18088"
		return "doisWeb", nil
	}

	res, err := SingleGrc(uri).SerList()
	if err != nil {
		return "", err
	}

	for _, server := range res.Services {
		if host == server.SerUrl {
			return server.SerName, nil
		}
	}

	return "", errors.New("服务不存在")
}

type Grpcurl struct {
	Host string
	conn *grpc.ClientConn
}

func (grc *Grpcurl) initGrpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grc.Host, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (grc *Grpcurl) Start(serName string) error {
	req := &gm.Request{SerName: serName, Active: "start"}
	if _, err := grc.action(req); err != nil {
		return err
	}
	return nil
}

func (grc *Grpcurl) Restart(serName, directory string) error {
	req := &gm.Request{SerName: serName, Active: "restart", Directory: directory}
	if _, err := grc.action(req); err != nil {
		return err
	}
	return nil
}

func (grc *Grpcurl) Stop(serName string) error {
	req := &gm.Request{SerName: serName, Active: "stop"}
	if _, err := grc.action(req); err != nil {
		return err
	}
	return nil
}

func (grc *Grpcurl) Status(serName string) (*gm.Response, error) {
	req := &gm.Request{SerName: serName, Active: "status"}
	response, err := grc.action(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Init
func (grc *Grpcurl) Init() error {
	conn, err := grc.initGrpcConn()
	if err != nil {
		return err
	}
	grc.conn = conn
	return nil
}

// action
func (grc *Grpcurl) action(req *gm.Request) (*gm.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := gm.NewGrpcManagerClient(grc.conn)
	re, err := c.ExecServiceControl(ctx, req)
	if err != nil {
		return nil, err
	}

	return re, nil
}

// activeDevice
func (grc *Grpcurl) Close() {
	if grc.conn != nil {
		grc.conn.Close()
	}
}

// list
func (grc *Grpcurl) list(req *gm.ServiceRequest) (*gm.ServiceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := gm.NewGrpcManagerClient(grc.conn)
	re, err := c.GetService(ctx, req)
	if err != nil {
		return nil, err
	}

	return re, nil
}

func (grc *Grpcurl) SerList() (*gm.ServiceResponse, error) {
	req := &gm.ServiceRequest{}
	response, err := grc.list(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// list
func (grc *Grpcurl) emqx(req *gm.EmqxRequest) (*gm.EmqxResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := gm.NewGrpcManagerClient(grc.conn)
	re, err := c.ExecEmqxControl(ctx, req)
	if err != nil {
		return nil, err
	}

	return re, nil
}

func (grc *Grpcurl) EmqxStart(serName string) error {
	req := &gm.EmqxRequest{SerName: serName, Active: "start"}
	resp, err := grc.emqx(req)
	if err != nil {
		return err
	}
	log.Println("EmqxStart:", resp.Message)
	return nil
}

func (grc *Grpcurl) EmqxStop(serName string) error {
	req := &gm.EmqxRequest{SerName: serName, Active: "stop"}
	if _, err := grc.emqx(req); err != nil {
		return err
	}
	return nil
}

func (grc *Grpcurl) EmqxStatus(serName string) (bool, error) {
	req := &gm.EmqxRequest{SerName: serName, Active: "status"}
	if res, err := grc.emqx(req); err != nil {
		return false, err
	} else {
		if res.Status == http.StatusBadRequest {
			return false, fmt.Errorf(res.Message)
		}
	}
	return true, nil
}
