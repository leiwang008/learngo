module example.com/hello

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/leiwang008/utils v0.0.5
	golang.org/x/tour v0.0.0-20210526031051-3891a3eb15c0
	rsc.io/quote v1.5.2
)

replace github.com/leiwang008/utils => ../utils
