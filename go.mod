module github.com/lee31802/gotemplate

go 1.24.2

require (
	cloud.google.com/go/bigquery v1.4.0 // indirect
	cloud.google.com/go/firestore v1.1.1 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.1
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/Masterminds/squirrel v1.4.0
	github.com/Shopify/sarama v1.26.4
	github.com/agiledragon/gomonkey v2.0.2+incompatible
	github.com/aws/aws-sdk-go v1.31.12
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didi/gendry v1.5.0
	github.com/dimchansky/utfbom v1.1.0
	github.com/eapache/go-resiliency v1.2.0
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/firstrow/tcp_server v0.1.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.1
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/websocket v1.4.1
	github.com/huandu/go-clone v1.1.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/json-iterator/go v1.1.11
	github.com/kr/pretty v0.3.0 // indirect
	github.com/micro/go-micro v1.16.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olivere/elastic/v7 v7.0.17
	github.com/panjf2000/ants/v2 v2.5.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pkg/sftp v1.12.0
	github.com/robfig/cron v1.2.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/rs/xid v1.2.1
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da
	github.com/sethvargo/go-limiter v0.6.0
	github.com/shopspring/decimal v1.2.0
	github.com/sideshow/apns2 v0.20.0
	github.com/spf13/cast v1.3.0
	github.com/stretchr/testify v1.8.0
	github.com/urfave/cli/v2 v2.3.0
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.opencensus.io v0.22.4 // indirect
	go.uber.org/atomic v1.9.0
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/api v0.15.1 // indirect
	google.golang.org/genproto v0.0.0-20200122232147-0452cf42e150
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/go-playground/validator.v9 v9.30.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
	github.com/micro/h2c v1.0.0 => golang.org/x/net v0.0.0-20200707034311-ab342639438
	github.com/micro/mdns => github.com/p2pNG/mdns v0.0.0-20201105170741-aaa81c17b902
	github.com/stretchr/testify => github.com/stretchr/testify v1.6.1
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

replace github.com/lee31802/comment_lib v1.0.1 => /Users/xiaolongli/Project/lee/comment_lib
