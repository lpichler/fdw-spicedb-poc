package fdwspicedbpoc

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableWorkspacesList(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "workspaces",
		Description: "workspaces to query by permission and user",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "user_name", Require: plugin.Required},
				{Name: "permission", Require: plugin.Required},
			},
			Hydrate: listWorkspaces,
		},
		Columns: workspaceCols(),
	}
}
