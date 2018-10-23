package main

import (
	"time"

	graphqlsample "github.com/svobodal/graphqlsample"
)

type UserResolver struct {
	user *graphqlsample.User
}

func (resolver *UserResolver) Id() int32 {
	return resolver.user.Id
}

func (resolver *UserResolver) Name() string {
	return resolver.user.Name
}

func (resolver *UserResolver) Email() *string {
	return &resolver.user.Email
}

func (resolver *UserResolver) Videos(args struct{ Count int32 }) []*VideoResolver {
	list := make([]*VideoResolver, 0)
	videos, err := graphqlsample.Model.GetVideosForUser(resolver.user.Id)
	if err != nil {
		panic(err)
	}

	for i, v := range videos {
		if args.Count >= 0 && i >= int(args.Count) {
			break
		}

		list = append(list, &VideoResolver{v})
	}
	return list
}

// CommentResolver returns data about a comment to a video.
type CommentResolver struct {
	comment *graphqlsample.Comment
}

func (resolver *CommentResolver) Id() int32 {
	return resolver.comment.Id
}

func (resolver *CommentResolver) Text() string {
	return resolver.comment.Text
}

func (resolver *CommentResolver) CreationTime() time.Time {
	return resolver.comment.CreationTime
}

func (resolver *CommentResolver) Creator() *UserResolver {
	return &UserResolver{resolver.comment.User}
}

// VideoResolver returns data about a single video.
type VideoResolver struct {
	video *graphqlsample.Video
}

func (resolver *VideoResolver) Id() int32 {
	return resolver.video.Id
}

func (resolver *VideoResolver) Name() string {
	return resolver.video.Name
}

func (resolver *VideoResolver) Views() int32 {
	return resolver.video.Views
}

func (resolver *VideoResolver) Description() string {
	return resolver.video.Name
}

func (resolver *VideoResolver) CreationTime() time.Time {
	return resolver.video.CreationTime
}

func (resolver *VideoResolver) Creator() *UserResolver {
	return &UserResolver{resolver.video.Creator}
}

func (resolver *VideoResolver) Comments(args struct{ Count int32 }) []*CommentResolver {
	list := make([]*CommentResolver, 0)
	comments, err := graphqlsample.Model.GetCommentsForVideo(resolver.video.Id)
	if err != nil {
		panic(err)
	}

	for i, comment := range comments {
		if args.Count >= 0 && i >= int(args.Count) {
			break
		}

		list = append(list, &CommentResolver{comment})
	}
	return list
}

func (resolver *VideoResolver) Related(args struct{ Count int32 }) []*VideoResolver {
	list := make([]*VideoResolver, 0)
	related, err := graphqlsample.Model.GetRelatedVideos(resolver.video.Id)
	if err != nil {
		panic(err)
	}

	for i, v := range related {
		if args.Count >= 0 && i >= int(args.Count) {
			break
		}

		list = append(list, &VideoResolver{v})
	}
	return list
}

// Query is GraphQL root query resolver.
type QueryResolver struct{}

func (resolver *QueryResolver) Video(args struct{ Id int32 }) (*VideoResolver, error) {
	video, err := graphqlsample.Model.GetVideo(args.Id)
	if err != nil {
		return nil, err
	}
	return &VideoResolver{video}, nil
}

func (resolver *QueryResolver) User(args struct{ Id int32 }) (*UserResolver, error) {
	user, err := graphqlsample.Model.GetUser(args.Id)
	if err != nil {
		return nil, err
	}
	return &UserResolver{user}, nil
}
