go env -w GOOS=linux
go env -w CGO_ENABLED=0
timeout /t 1
go build -ldflags "-w -s" -o ./emqx-grpc .\main.go
timeout /t 3
plink root@10.0.0.34 -pw P@ssw0rd (sudo service EmqxGrpc stop)
timeout /t 3
pscp -v -P 22 -pw P@ssw0rd ./emqx-grpc root@10.0.0.34:/usr/local/bin
timeout /t 3
plink root@10.0.0.34 -pw P@ssw0rd (sudo chmod +x /usr/local/bin/emqx-grpc)
timeout /t 3
plink root@10.0.0.34 -pw P@ssw0rd (sudo service EmqxGrpc start)
timeout /t 3
go env -w GOOS=windows
go env -w CGO_ENABLED=1
