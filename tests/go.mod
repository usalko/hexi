module github.com/usalko/hexi/tests

go 1.23

require (
	github.com/stretchr/testify v1.9.0
	github.com/usalko/hexi v0.1.11
	github.com/usalko/hexi/ft v0.1.11
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/usalko/hexi => ../
	github.com/usalko/hexi/ft => ../ft
)
