module github.com/hxx258456/pyramidel-chain-baas

go 1.17

require (
	github.com/docker/docker v23.0.0+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/fsnotify/fsnotify v1.6.0
	github.com/gin-gonic/gin v1.8.2
	github.com/hxx258456/ccgo v0.0.4
	github.com/jinzhu/copier v0.3.5
	github.com/melbahja/goph v1.3.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus-community/pro-bing v0.1.0
	github.com/shirou/gopsutil/v3 v3.23.1
	github.com/spf13/cast v1.5.0
	github.com/spf13/viper v1.14.0
	github.com/streadway/amqp v1.0.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.5.0
	gorm.io/driver/mysql v1.4.6
	gorm.io/gorm v1.24.5
)

require (
	gitee.com/zhaochuninhefei/zcgolog v0.0.16 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/sftp v1.13.4 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/ugorji/go/codec v1.2.8 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.4.0 // indirect
)

replace (
	// github.com/hxx258456/fabric-protos-go-gm => /home/zhaochun/work/sources/gcbaas/fabric-protos-go-gm
	github.com/zmap/zcrypto => github.com/zmap/zcrypto v0.0.0-20190729165852-9051775e6a2e
	github.com/zmap/zlint => github.com/zmap/zlint v0.0.0-20190806154020-fd021b4cfbeb
)
