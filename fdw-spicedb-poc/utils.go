package fdwspicedbpoc

import (
	"context"
	e "errors"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/lpichler/fdw-spicedb-poc/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"io"
)

type Workspace struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	UserName   string `json:"user_name"`
	Permission string `json:"permission"`
}

var (
	spiceDBURL   = "localhost:50051"
	spiceDBToken = "foobar"
)

var SpiceDbClient *authzed.Client

func workspaceCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "id", Type: proto.ColumnType_INT, Description: "id"},
		{Name: "name", Type: proto.ColumnType_STRING, Description: "name"},
		{Name: "user_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("user_name"), Description: "user_name"},
		{Name: "permission", Type: proto.ColumnType_STRING, Transform: transform.FromQual("permission"), Description: "permission"},
	}
}

func initSpiceDB() {
	if SpiceDbClient == nil {
		SpiceDbClient = client.InitServer(spiceDBURL, spiceDBToken)
	}
}

func listWorkspaces(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	initSpiceDB()

	userName := d.EqualsQualString("user_name")
	if userName == "" {
		// Empty command returns zero rows
		return nil, nil
	}

	permission := d.EqualsQualString("permission")
	if permission == "" {
		// Empty command returns zero rows
		return nil, nil
	}

	lrClient, err := SpiceDbClient.LookupResources(ctx, &v1.LookupResourcesRequest{
		ResourceObjectType: "workspace",
		Permission:         permission,
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   userName,
			},
		},
	})

	if err != nil {
		plugin.Logger(ctx).Error("spicedb error: %v", err)
		return nil, err
	}

	var workspaces []string
	for {
		next, err := lrClient.Recv()
		if e.Is(err, io.EOF) {
			break
		}
		if err != nil {
			plugin.Logger(ctx).Error("spicedb error: %v", err)
			return nil, err
		}

		workspaces = append(workspaces, next.GetResourceObjectId())
	}

	for i, workspace := range workspaces {
		w := Workspace{i + 1, workspace, permission, userName}
		d.StreamListItem(ctx, &w)
	}
	return nil, nil
}
