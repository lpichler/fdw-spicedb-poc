package main

import (
	fdwspicedbpoc "github.com/lpichler/fdw-spicedb-poc/fdw-spicedb-poc"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: fdwspicedbpoc.Plugin})
}
