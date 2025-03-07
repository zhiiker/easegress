module github.com/megaease/easegress

go 1.16

require (
	github.com/ArthurHlt/go-eureka-client v1.1.0
	github.com/Shopify/sarama v1.29.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/fatih/color v1.12.0
	github.com/fatih/structs v1.1.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-chi/chi/v5 v5.0.3
	github.com/go-zookeeper/zk v1.0.2
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/golang-lru v0.5.4
	github.com/json-iterator/go v1.1.11
	github.com/klauspost/compress v1.13.1
	github.com/lucas-clemente/quic-go v0.21.1
	github.com/megaease/easemesh-api v0.0.0-20210604095307-27c2d1f7cf09
	github.com/megaease/grace v1.0.0
	github.com/megaease/jsonschema v0.0.0-20191230042224-e92108cfafc5
	github.com/mitchellh/mapstructure v1.4.1
	github.com/nacos-group/nacos-sdk-go v1.0.8
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.0
	github.com/tcnksm/go-httpstat v0.2.1-0.20191008022543-e866bb274419
	github.com/tidwall/gjson v1.8.0
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/valyala/fasttemplate v1.2.1
	github.com/xeipuuv/gojsonschema v1.2.0
	github.com/yl2chen/cidranger v1.0.2
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.etcd.io/etcd/server/v3 v3.5.0
	go.uber.org/zap v1.17.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/cli-runtime v0.20.7 // indirect
	knative.dev/client v0.23.1
	knative.dev/serving v0.23.1-0.20210614141420-380a090c2039
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
