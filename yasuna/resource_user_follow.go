package yasuna

import (
	"net/url"
	"strconv"
)

// GetUserFollow returns a list of users the specified user ID is following.
//
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func (t *Twitter) GetUserFollowing(id int64, opts ...GetUserFollowOption) (*Response[[]*User], error) {
	return t.getUserFollows("following", id, opts...)
}

// GetUserFollowers returns a list of users who are followers of the specified user ID.
//
// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func (t *Twitter) GetUserFollowers(id int64, opts ...GetUserFollowOption) (*Response[[]*User], error) {
	return t.getUserFollows("followers", id, opts...)
}

func (t *Twitter) getUserFollows(f string, id int64, opts ...GetUserFollowOption) (*Response[[]*User], error) {
	v := url.Values{}
	for _, opt := range opts {
		opt.optGetUserFollow(v)
	}
	path := "users/" + strconv.FormatInt(id, 10) + "/" + f
	data := new(Response[[]*User])
	return data, t.get(path, v, data)
}

type GetUserFollowOption interface {
	optGetUserFollow(v url.Values)
}
