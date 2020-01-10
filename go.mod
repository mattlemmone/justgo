module github.com/mattlemmone/popo

go 1.12

require (
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/stretchr/testify v1.4.0
	github.com/therecipe/qt v0.0.0-20200103041036-2b818d970888
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace github.com/mattlemmone/popo/ => ./

// replace github.com/mattlemmone/popo/desktop => ./desktop/
