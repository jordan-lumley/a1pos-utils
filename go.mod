module github.com/jordan-lumley/a1pos

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/jordan-lumley/service v0.0.0
	github.com/shirou/gopsutil v3.20.10+incompatible
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)

replace github.com/jordan-lumley/service => /pkg/service

go 1.15
