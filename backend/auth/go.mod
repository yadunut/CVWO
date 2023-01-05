module github.com/yadunut/CVWO/backend/auth

go 1.19

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jackc/pgconn v1.13.0
	github.com/joho/godotenv v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/yadunut/CVWO/backend/database v0.0.0
	github.com/yadunut/CVWO/backend/proto v0.0.0
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	google.golang.org/grpc v1.51.0
	gorm.io/driver/postgres v1.4.5
	gorm.io/gorm v1.24.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/yadunut/CVWO/backend/database v0.0.0 => ../database

replace github.com/yadunut/CVWO/backend/proto v0.0.0 => ../proto
