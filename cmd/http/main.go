package main

import (
	"github.com/otridelvi/go-svc-eivor/internal/interfaces/server/container"
	"github.com/otridelvi/go-svc-eivor/internal/interfaces/server/http"
)

func main() {
	// start http server
	http.StartHttpService(container.Setup())
}
