package graphQL

import (
	"github.com/graphql-go/graphql"
)

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"album": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var merchantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Merchant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_phone_number": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_email": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_address": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_city": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_province": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_website": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_category_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_media_social_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_description": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_image_profile": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_gallery": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var socialmediaType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Card",
	Fields: graphql.Fields{
		"id" : &graphql.Field{
			Type:graphql.Int,
		},
		"created" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"created_by" : &graphql.Field{
			Type:graphql.String,
		},
		"modified" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"modified_by" : &graphql.Field{
			Type:graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"deleted" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type:graphql.String,
		},
		"social_media_name": &graphql.Field{
			Type:graphql.String,
		},
		"image_url": &graphql.Field{
			Type:graphql.String,
		},
	},
})

var cardType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Card",
	Fields: graphql.Fields{
		"id" : &graphql.Field{
			Type:graphql.Int,
		},
		"created" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"created_by" : &graphql.Field{
			Type:graphql.String,
		},
		"modified" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"modified_by" : &graphql.Field{
			Type:graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"deleted" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type:graphql.String,
		},
		"card_type_name": &graphql.Field{
			Type:graphql.String,
		},
	},
})

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Category",
	Fields: graphql.Fields{
		"id" : &graphql.Field{
			Type:graphql.Int,
		},
		"created" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"created_by" : &graphql.Field{
			Type:graphql.String,
		},
		"modified" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"modified_by" : &graphql.Field{
			Type:graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"deleted" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type:graphql.String,
		},
		"category_name": &graphql.Field{
			Type:graphql.String,
		},
		"image_url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var programType = graphql.NewObject(graphql.ObjectConfig{
	Name : "Program",
	Fields: graphql.Fields{
		"id" : &graphql.Field{
			Type:graphql.Int,
		},
		"created" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"created_by" : &graphql.Field{
			Type:graphql.String,
		},
		"modified" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"modified_by" : &graphql.Field{
			Type:graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"deleted" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type:graphql.String,
		},
		"program_name": &graphql.Field{
			Type:graphql.String,
		},
		"program_image": &graphql.Field{
			Type:graphql.String,
		},
		"program_start_date": &graphql.Field{
			Type:graphql.DateTime,
		},
		"program_end_date": &graphql.Field{
			Type:graphql.DateTime,
		},
		"program_description": &graphql.Field{
			Type:graphql.String,
		},
		"card": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_id": &graphql.Field{
			Type:graphql.Int,
		},
		"merchant_id": &graphql.Field{
			Type:graphql.Int,
		},
		"merchant_name": &graphql.Field{
			Type:graphql.String,
		},
	},
})

var outletType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Outlet",
	Fields:      graphql.Fields{
		"id" : &graphql.Field{
			Type: graphql.Int,
		},
		"created" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"created_by" : &graphql.Field{
			Type:graphql.String,
		},
		"modified" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"modified_by" : &graphql.Field{
			Type:graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"deleted" : &graphql.Field{
			Type:graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_name": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_address": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_phone": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_city": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_province": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_postal_code": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_longitude": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_latitude": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_id": &graphql.Field{
			Type:graphql.String,
		},
	},
})