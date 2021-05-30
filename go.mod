module example.com/hello

go 1.16

require (
	example.com/morestrings v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.6
	github.com/leiwang008/utils v0.0.4
	rsc.io/quote v1.5.2
)

replace example.com/morestrings => ../morestrings
