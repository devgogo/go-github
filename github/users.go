package github

import (
	"context"
	"fmt"
)

type UsersService service

type User struct {
	Login     *string `json:"login,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
	Email     *string `json:"email,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	URL       *string `json:"url,omitempty"`
}

func (s *UsersService) Get(ctx context.Context, user string) (*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	uResp := new(User)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}
	return uResp, resp, nil
}
