package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		nameFlag = getSerName(hostFlag)
		if nameFlag == "" {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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

func getSerName(host string) string {
	if host == "http://10.0.0.23" {
		return "doisWeb"
	}
	switch host {
	case "http://dgzyh.server.chindeo.test":
		return "DGZYHospitalV1"
	case "http://zhrmh.server.chindeo.test":
		return "ZHRMV2"
	case "http://whsyh.server.chindeo.test":
		return "WHSYV2"
	case "http://tjhhh.server.chindeo.test":
		return "TJHHV1"
	case "http://szbdh.server.chindeo.test":
		return "SZBDHospitalV1"
	case "http://szsmh.server.chindeo.test":
		return "SMV1"
	case "http://shxhh.server.chindeo.test":
		return "SHXHHospitalV1"
	case "http://qyysh.server.chindeo.test":
		return "QYYSHospitalV1"
	case "http://nfykdh.server.chindeo.test":
		return "NFYKHospitalV1"
	case "http://lhrmh.server.chindeo.test":
		return "LHV2"
	case "http://lgzxh.server.chindeo.test":
		return "LGZXHospitalV1"
	case "http://lggkh.server.chindeo.test":
		return "LGGGHospitalV1"
	case "http://jzgah.server.chindeo.test":
		return "JZGAXRMV1"
	case "http://gzhqh.server.chindeo.test":
		return "GZHQV1"
	case "http://bazxh.server.chindeo.test":
		return "BAZXV2"
	case "http://barmh.server.chindeo.test":
		return "BAHospitalV2"
	case "http://bafyh.server.chindeo.test":
		return "BAFYBJV1"
	case "http://szfth.server.chindeo.test":
		return "FTHospital"
	case "http://group.chindeo.test":
		return "GZWYHospitalV1"
	case "http://fxkzh.server.chindeo.test":
		return "FXKZHospitalV1"
	default:
		return ""
	}
}
