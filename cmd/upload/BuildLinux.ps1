go build  -ldflags '-w -s'  -o ./upload.exe ./main.go

timeout /t 1

go env -w GOOS=linux 
go env -w CGO_ENABLED=0

go build  -ldflags "-w -s"  -o ./upload main.go

timeout /t 1

pscp -v -P 22 -pw "P@ssw0rd" ./upload "root@10.0.1.83:/var/www/html/upload"

timeout /t 1

pscp -v -P 22 -pw "P@ssw0rd" ./upload.exe "root@10.0.1.83:/var/www/html/upload.exe"

go env -w GOOS=windows 
go env -w CGO_ENABLED=1

