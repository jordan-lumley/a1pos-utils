module github.com/jordan-lumley/a1pos

require (
	github.com/jordan-lumley/service v0.0.0
)

replace (
	github.com/jordan-lumley/service => /pkg/service
)

go 1.15
