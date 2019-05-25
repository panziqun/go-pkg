module github.com/laughmaker/go-pkg

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190228050851-31a38585487a
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190221220918-438050ddec5e
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190228124157-a34e9553db1e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190221204921-83362c3779f5
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190227213309-4f5b463f9597
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.0
)

require (
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190524120938-d471c41f6e1f
	github.com/aliyun/aliyun-log-go-sdk v0.0.0-20190514033836-f7ccbaffaef0
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.42.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.8
	github.com/jmoiron/sqlx v1.2.0
	github.com/olivere/elastic/v7 v7.0.1
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.2
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
