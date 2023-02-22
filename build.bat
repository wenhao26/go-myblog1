SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ./bin/coinskyManage ./main.go  
go build -o ./bin/coinskyApi ./cmd_api/main.go 
go build -o ./bin/coinskyWebApi ./cmd_webapi/main.go  
go build -o ./bin/coinskyAdminApi ./cmd_admin/main.go 


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o ./bin/coinskyManageArm ./main.go
go build -o ./bin/coinskyApiArm ./cmd_api/main.go
go build -o ./bin/coinskyWebApiArm ./cmd_webapi/main.go
go build -o ./bin/coinskyAdminApiArm ./cmd_admin/main.go
