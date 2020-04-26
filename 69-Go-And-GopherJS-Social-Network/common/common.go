package common

import (
	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
