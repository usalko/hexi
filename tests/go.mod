module github.com/usalko/hexsi/tests

go 1.23

require (
	github.com/stretchr/testify v1.9.0
	github.com/usalko/hexsi v0.1.3
	github.com/usalko/hexsi/internal v0.1.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/usalko/hexsi => ../
	github.com/usalko/hexsi/internal => ../internal
)
