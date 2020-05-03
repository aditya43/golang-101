package common

import (
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
