package main

import (
	graphqlsample "../.."
	"github.com/graphql-go/graphql"
)

var UserResolver *graphql.Object
var CommentResolver *graphql.Object
var VideoResolver *graphql.Object
var QueryResolver *graphql.Object

func init() {
	UserResolver = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	CommentResolver = graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"creator": &graphql.Field{
				Type: UserResolver,
			},
		},
	})

	VideoResolver = graphql.NewObject(graphql.ObjectConfig{
		Name: "Video",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"views": &graphql.Field{
				Type: graphql.Int,
			},
			"creator": &graphql.Field{
				Type: UserResolver,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					video := p.Source.(*graphqlsample.Video)
					return video.Creator, nil
				},
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(CommentResolver),
				Args: graphql.FieldConfigArgument{
					"count": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: -1,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					video := p.Source.(*graphqlsample.Video)
					comments, err := graphqlsample.Model.GetCommentForVideo(video.Id)
					if err != nil {
						return nil, err
					}

					count := p.Args["count"].(int)
					if count < 0 || count > len(comments) {
						count = len(comments)
					}

					return comments[0:count], nil
				},
			},
		},
	})

	VideoResolver.AddFieldConfig("related", &graphql.Field{
		Type: graphql.NewList(VideoResolver),
		Args: graphql.FieldConfigArgument{
			"count": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: -1,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			video := p.Source.(*graphqlsample.Video)
			videos, err := graphqlsample.Model.GetRelatedVideos(video.Id)
			if err != nil {
				return nil, err
			}

			count := p.Args["count"].(int)
			if count < 0 || count > len(videos) {
				count = len(videos)
			}

			return videos[0:count], nil
		},
	})

	QueryResolver = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"video": &graphql.Field{
				Type: VideoResolver,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := int32(p.Args["id"].(int))
					return graphqlsample.Model.GetVideo(id)
				},
			},
		},
	})
}
