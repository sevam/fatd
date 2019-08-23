module github.com/Factom-Asset-Tokens/fatd

go 1.12

require (
	crawshaw.io/sqlite v0.1.3-0.20190520153332-66f853b01dfb
	github.com/AdamSLevy/go-merkle v0.0.0-20190611101253-ca33344a884d
	github.com/AdamSLevy/jsonrpc2/v11 v11.3.2
	github.com/AdamSLevy/sqlitechangeset v0.0.0-20190731000048-d57789e63df5
	github.com/Factom-Asset-Tokens/base58 v0.0.0-20181227014902-61655c4dd885
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/posener/complete v1.2.1
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3 // indirect
	golang.org/x/text v0.3.2 // indirect
)

replace github.com/gocraft/dbr => github.com/AdamSLevy/dbr v0.0.0-20190429075658-5db28ac75cea

replace github.com/spf13/pflag v1.0.3 => github.com/AdamSLevy/pflag v1.0.4

replace crawshaw.io/sqlite => github.com/AdamSLevy/sqlite v0.1.3-0.20190729192944-6cbd592f144c
