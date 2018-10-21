package graphqlsample

import "time"

type UserResolver struct {
	user *User
}

func (resolver *UserResolver) Id() int32 {
	return resolver.user.ID
}

func (resolver *UserResolver) Name() string {
	return resolver.user.Name
}

func (resolver *UserResolver) Email() *string {
	return &resolver.user.Email
}

type CommentResolver struct {
	comment *Comment
}

func (resolver *CommentResolver) Id() int32 {
	return resolver.comment.ID
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
	video *Video
}

func (resolver *VideoResolver) Id() int32 {
	return resolver.video.ID
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
	comments, err := Model.GetCommentForVideo(resolver.video.ID)
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
	related, err := Model.GetRelatedVideos(resolver.video.ID)
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
type Query struct{}

func (resolver *Query) Video(args struct{ Id int32 }) *VideoResolver {
	video, err := Model.GetVideo(args.Id)
	if err != nil {
		return nil
	}

	return &VideoResolver{video}
}
