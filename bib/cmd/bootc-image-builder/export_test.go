package main

var CanChownInPath = canChownInPath
var IsBootcImage = isBootcImage

func MockOsGetuid(new func() int) (restore func()) {
	saved := osGetuid
	osGetuid = new
	return func() {
		osGetuid = saved
	}
}
