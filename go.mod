module github.com/usalko/hexsi

go 1.23

require (
	github.com/usalko/hexsi/internal v0.0.0-20241113001323-d73806f784e9
	github.com/usalko/hexsi/tests v0.0.0-20241113001323-d73806f784e9
)

replace (
	github.com/usalko/hexsi/internal => ./internal
	github.com/usalko/hexsi/tests => ./tests
)
