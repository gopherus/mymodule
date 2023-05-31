module github.com/gopherus/mymodule

go 1.17

require github.com/viant/xdatly/types/core v0.0.0-20230518142915-e849977dfa52

//replace github.com/gopherus/mymodule => /Users/michael/GolandProjects/myproject/pkg
replace github.com/gopherus/mymodule => ./pkg
