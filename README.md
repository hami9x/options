options
=======

This package provides a simple, flexible and convenient way of defining config options in Go.  
  
Without the need for explaination, when we need to specify config options for something, especially when there are a lot of optional options, the naive way of using struct or function with a long list of arguments doesn't work.  

This package uses the power of Go reflections to make that task very simple. This is how you define and use a set of options with the help of this package:  

	//An option
	type WhateverName struct{ Value string }
	
	//An option
	type WhateverNum struct{ Value int }
	
	//WhateverSpec specifies the list of options
	type WhateverSpec struct {
		Name WhateverName `js:"name"` //You may specify a tag here
		Num  WhateverNum  `js:"num"`
	}
	
	//Example convenient method for creating the master config
	func WhateverOptions(opts ...Option) *OptionsProvider {
		return NewOptions(&WhateverSpec{}).Options(opts...)
	}
	
	//Example convenient method for retrieving the options
	func GetWhateverOptions(o *OptionsProvider) *WhateverSpec {
		return o.Get().(*WhateverSpec)
	}
	
	func main() {
		//This is how you use the defined options
		opts := WhateverOptions(
			WhateverName{"n0t9r34t6cz..."},
			WhateverNum{99999},
		)
	
		println(GetWhateverOptions(opts).Name.Value) //Get the Name option value
	
		//You can export the options to a map
		m := opts.ExportToMapWithTag("js")
		_ = m["name"].(string) //Name exported to key "name", as specified by the tag
	}


###Note:  
Rob Pike once described a pattern for solving this problem ([link](http://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)). The client interface is good, it has rollbacks which is nice, but it's really really tedious to have to copy-paste dozens of lines of duplicated code for making the options, repeating the same logic for every options pack we create. It might be good for some, but for me it's unbearable. That's why this simple library is born.
