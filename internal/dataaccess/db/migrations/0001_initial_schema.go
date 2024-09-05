package migrations

import (
	"database/sql"
	"errors"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"google.golang.org/protobuf/proto"
)

func Register0001Migration(app core.App) error {
	downloadTaskCollection, err := app.Dao().FindCollectionByNameOrId("download_tasks")
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if downloadTaskCollection != nil {
		return nil
	}

	form := forms.NewCollectionUpsert(app, &models.Collection{})
	form.Name = "download_tasks"
	form.Type = models.CollectionTypeBase
	form.Schema.AddField(&schema.SchemaField{
		Name:     "of_account_id",
		Type:     schema.FieldTypeRelation,
		Required: true,
		Options: &schema.RelationOptions{
			CollectionId:  "_pb_user_auth_",
			CascadeDelete: true,
		},
	})
	form.Schema.AddField(&schema.SchemaField{
		Name:     "download_type",
		Type:     schema.FieldTypeNumber,
		Required: true,
		Options: &schema.NumberOptions{
			Min:       proto.Float64(1),
			Max:       proto.Float64(1),
			NoDecimal: true,
		},
	})
	form.Schema.AddField(&schema.SchemaField{
		Name:     "url",
		Type:     schema.FieldTypeUrl,
		Required: true,
	})
	form.Schema.AddField(&schema.SchemaField{
		Name:     "download_status",
		Type:     schema.FieldTypeNumber,
		Required: true,
		Options: &schema.NumberOptions{
			Min:       proto.Float64(0),
			Max:       proto.Float64(4),
			NoDecimal: true,
		},
	})
	form.ListRule = proto.String("@request.auth.id = of_account_id")
	form.ViewRule = proto.String("@request.auth.id = of_account_id")
	form.UpdateRule = proto.String("@request.auth.id = of_account_id")
	form.DeleteRule = proto.String("@request.auth.id = of_account_id")

	return form.Submit()
}
