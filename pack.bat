go get -u github.com/intmian/mian_go_lib
go mod tidy
go build -o pack
copy setting_tep.json pack\setting.json
msg * "suc"