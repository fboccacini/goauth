module github.com/fboccacini/goauth

go 1.15

replace (
  github.com/fboccacini/goauth/pkg/cmd => ./pkg/cmd
)

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)
