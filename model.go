package graphqlsample

import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("Not found")

type User struct {
	Id    int32
	Name  string
	Email string
}

type Video struct {
	Id           int32
	Name         string
	Views        int32
	Description  string
	CreationTime time.Time
	Creator      *User
}

type Comment struct {
	Id           int32
	Video        *Video
	User         *User
	CreationTime time.Time
	Text         string
}

type model struct {
	Users    []*User
	Videos   []*Video
	Comments []*Comment
}

var Model *model

func (m *model) GetUser(id int32) (*User, error) {
	for _, u := range m.Users {
		if u.Id == id {
			return u, nil
		}
	}

	return nil, ErrNotFound
}

func (m *model) GetVideo(id int32) (*Video, error) {
	for _, v := range m.Videos {
		if v.Id == id {
			return v, nil
		}
	}

	return nil, ErrNotFound
}

func (m *model) GetCommentsForVideo(videoId int32) ([]*Comment, error) {
	list := make([]*Comment, 0)
	for _, c := range m.Comments {
		if c.Video.Id == videoId {
			list = append(list, c)
		}
	}

	return list, nil
}

func (m *model) GetRelatedVideos(videoId int32) ([]*Video, error) {
	list := make([]*Video, 0)
	for _, r := range m.Videos {
		if r.Id != videoId {
			list = append(list, r)
		}
	}

	return list, nil
}

func (m *model) GetVideosForUser(userId int32) ([]*Video, error) {
	list := make([]*Video, 0)
	for _, v := range m.Videos {
		if v.Creator.Id == userId {
			list = append(list, v)
		}
	}

	return list, nil
}

func init() {
	u1 := &User{
		Id:    1,
		Name:  "Charles",
		Email: "charles@example.com",
	}
	u2 := &User{
		Id:    2,
		Name:  "Richard",
		Email: "Richard@example.com",
	}
	u3 := &User{
		Id:    3,
		Name:  "Rob",
		Email: "rob@example.com",
	}

	v1 := &Video{
		Id:           10,
		Name:         "Funny cat",
		CreationTime: time.Now().Add(-5 * time.Hour),
		Views:        700,
		Description:  "Funny video of my cat",
		Creator:      u1,
	}

	v2 := &Video{
		Id:           20,
		Name:         "Running dog",
		CreationTime: time.Now().Add(-10 * time.Hour),
		Views:        1500,
		Description:  "Me playing with my dog",
		Creator:      u1,
	}

	v3 := &Video{
		Id:           30,
		Name:         "Hamster eating",
		CreationTime: time.Now().Add(-15 * time.Hour),
		Views:        1200,
		Description:  "My hamster eating a cabbage",
		Creator:      u2,
	}

	comments := []*Comment{
		&Comment{
			Id:           11,
			Video:        v1,
			User:         u2,
			CreationTime: time.Now().Add(-4 * time.Hour),
			Text:         "Cute!",
		},
		&Comment{
			Id:           12,
			Video:        v1,
			User:         u3,
			CreationTime: time.Now().Add(-3 * time.Hour),
			Text:         "Nice",
		},
	}

	Model = &model{
		Users:    []*User{u1, u2, u3},
		Videos:   []*Video{v1, v2, v3},
		Comments: comments,
	}
}
