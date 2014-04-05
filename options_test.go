package options

import (
	"testing"
	. "launchpad.net/gocheck"
)

//An option
type WhateverName struct{ Value string }

//An option
type WhateverNum struct{ Value int }

//Another option
type WhateverNum2 struct{ Value float32 }

//WhateverSpec specifies the options
type WhateverSpec struct {
	Name WhateverName `js:"name"`
	Num  WhateverNum  `js:"num"`
	Num2 WhateverNum2 `js:"num2"`
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
	c.Check(opts.IsSet("Num2"), Equals, false)
	_, ok := m["num2"] //num2 doesn't appear in the result because it has not been set
	c.Check(ok, Equals, false)
}
