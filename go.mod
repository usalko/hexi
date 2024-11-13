module github.com/usalko/hexsi

go 1.23

require github.com/usalko/hexsi/internal v0.1.1

replace (
	github.com/usalko/hexsi/internal => ./internal
	github.com/usalko/hexsi/tests => ./tests
)
