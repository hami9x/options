package options

import (
	"testing"
	. "launchpad.net/gocheck"
)

//An option
type WhateverName struct{ Value string }

//An option
type WhateverNum struct{ Value int }

//WhateverSpec specifies the options
type WhateverSpec struct {
	Name WhateverName `js:"name"`
	Num  WhateverNum  `js:"num"`
}

func WhateverOptions(opts ...Option) *OptionsProvider {
	return NewOptions(&WhateverSpec{}).Options(opts...)
}

func GetWhateverOptions(o *OptionsProvider) *WhateverSpec {
	return o.Get().(*WhateverSpec)
}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestOptions(c *C) {
	name, num := "n0t9r34t6cz...", 999999
	opts := WhateverOptions(
		WhateverName{name},
		WhateverNum{num},
	)
	m := opts.ExportToMapWithTag("js")
	c.Check(GetWhateverOptions(opts).Name.Value, Equals, name)
	c.Check(m["name"].(string), Equals, name)
	c.Check(m["num"].(int), Equals, num)
}
