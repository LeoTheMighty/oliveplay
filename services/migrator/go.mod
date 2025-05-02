module catescafe.com/migrator

go 1.24

require (
	cate/cate_utils v0.0.0
	github.com/golang-migrate/migrate/v4 v4.18.2
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-github/v39 v39.2.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/oauth2 v0.18.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)

replace cate/cate_utils => ../../libraries/cate_utils
