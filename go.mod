module github.com/eddiewebb/blync-studio-light

//goblync uses hid which needs an update due to changes in cgo 1.10 - https://github.com/boombuler/hid/pull/15
replace github.com/boombuler/hid => github.com/eddiewebb/hid v0.0.0-20190226232454-cc6173fefbcb

require (
	github.com/boombuler/hid v0.0.0-20180620055412-8263579894f5 // indirect
	github.com/eddiewebb/goblync v0.0.0-20151214232719-d5f54f59e81b
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.1
	golang.org/x/net v0.0.0-20190227160552-c95aed5357e7
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
	google.golang.org/api v0.1.0
	gotest.tools v2.2.0+incompatible
)

go 1.13
