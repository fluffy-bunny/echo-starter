package shared

import di "github.com/fluffy-bunny/sarulabsdi"

var RootContainer di.Container

func GetRootContainer() di.Container {
	return RootContainer
}
