rd pack /q /s
if not exist pack md pack
if not exist pack\static md pack\static
if not exist pack\static\log md pack\static\log

if not exist pack\http md pack\http
go get -u github.com/intmian/mian_go_lib
go mod tidy
go build -o pack
copy setting_tep.json pack\setting.json
Xcopy http\static\ pack\http\static\ /Y /E
msg * "suc"