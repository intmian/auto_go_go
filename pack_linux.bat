rd pack /q /s
if not exist pack md pack
if not exist pack\static md pack\static

if not exist pack\http md pack\http
go get -u github.com/intmian/mian_go_lib
go mod tidy

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64

go build -o pack
copy setting_tep.json pack\setting.json
Xcopy http\static\ pack\http\static\ /Y /E
msg * "suc"