package datastruct

import (
	"io/ioutil"

	"github.com/influxdata/toml"
	"github.com/influxdata/toml/ast"
)

func parseConfig(f string) (*ast.Table, error) {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return toml.Parse(content)
}
