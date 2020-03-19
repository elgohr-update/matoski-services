module github.com/micro/services/home/api

go 1.13

require (
	github.com/golang/protobuf v1.3.4
	github.com/micro/go-micro/v2 v2.3.1-0.20200319095459-cbb958def528
	github.com/micro/services/apps/service v0.0.0-20200318105532-9c3078c484d5
	github.com/micro/services/users/service v0.0.0-20200313151537-5407234f5db7
)

replace github.com/micro/services/apps/service => ../../apps/service
