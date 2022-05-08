package repos

import "io"

// DepStorer serves as a registry/cache for repositories informations
// We use it to rebuild the Dependency Graph
type DepStorer interface {
	io.ReadWriter
}

type DepEntry struct {
	Name, Version, Lang string
}

type RepositoryAdapter interface {
	UpdateDepRegistry(reg DepStorer) error
}
