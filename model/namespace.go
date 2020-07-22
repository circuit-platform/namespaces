package model

import (
	"github.com/rs/rest-layer/schema"
	"github.com/rs/rest-layer/resource"
	
	"github.com/apuigsech/rest-layer-sql"

	utils "github.com/circuit-platform/models-utils"
)

var (
	Namespace = schema.Schema{
		Fields: schema.Fields{
			"id": schema.IDField,
			"created": utils.CreatedField,
			"updated": utils.UpdatedField,
			"name": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}

	NamespaceSetting = schema.Schema{
		Fields: schema.Fields{
			"id": schema.IDField,
			"created": utils.CreatedField,
			"updated": utils.UpdatedField,
			"namespace": {
				Required:   true,
				Filterable: true,
				ReadOnly:   true,
				Validator: &schema.Reference{
					Path: "namespaces",
				},
			},
			"key": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
			"value": schema.Field{
				Required: true,
				Filterable: true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}
)

func CreateNamespacesIndex(databaseSource string, databaseSchema string) resource.Index {
	index := resource.NewIndex()

	if databaseSchema != "" {
		databaseSchema = databaseSchema + "."
	}

	namespaces := index.Bind("namespaces", Namespace, utils.CreateStorerOrDie(databaseSource, databaseSchema + "namespaces",
		&sqlStorage.Config{},
	), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	namespaces.Bind("settings", "namespace", NamespaceSetting, utils.CreateStorerOrDie(databaseSource, databaseSchema + "namespaces_settings",
		&sqlStorage.Config{},
	), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	return index
}