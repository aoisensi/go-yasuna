package yasuna

import (
	"strings"
)

// Scope
//
//https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code#scopes
type Scope string

const (
	// All the Tweets you can view, including Tweets from protected accounts.
	ScopeTweetRead Scope = "tweet.read"
	// Tweet and Retweet for you.
	ScopeTweetWrite Scope = "tweet.write"
	// Hide and unhide replies to your Tweets.
	ScopeTweetModerateWrite Scope = "tweet.moderate.write"
	// Any account you can view, including protected accounts.
	ScopeUsersRead Scope = "users.read"
	// People who follow you and people who you follow.
	ScopeFollowsRead Scope = "follows.read"
	// Follow and unfollow people for you.
	ScopeFollowsWrite Scope = "follows.write"
	// Stay connected to your account until you revoke access.
	ScopeOfflineAccess Scope = "offline.access"
	// All the Spaces you can view.
	ScopeSpaceRead Scope = "space.read"
	// Accounts you’ve muted.
	ScopeMuteRead Scope = "mute.read"
	// Mute and unmute accounts for you.
	ScopeMuteWrite Scope = "mute.write"
	// Tweets you’ve liked and likes you can view.
	ScopeLikeRead Scope = "like.read"
	// Like and un-like Tweets for you.
	ScopeLikeWrite Scope = "like.write"
	// Lists, list members, and list followers of lists you’ve created or are a member of, including private lists.
	ScopeListRead Scope = "list.read"
	// Create and manage Lists for you.
	ScopeListWrite Scope = "list.write"
	// Accounts you’ve blocked.
	ScopeBlockRead Scope = "block.read"
	// Block and unblock accounts for you.
	ScopeBlockWrite Scope = "block.write"
	// Get Bookmarked Tweets from an authenticated user.
	ScopeBookmarkRead Scope = "bookmark.read"
	// Bookmark and remove Bookmarks from Tweets
	ScopeBookmarkWrite Scope = "bookmark.write"
)

var (
	ScopeTweet    Scope = fmtScope(ScopeTweetRead, ScopeTweetWrite, ScopeTweetModerateWrite)
	ScopeMute     Scope = fmtScope(ScopeMuteRead, ScopeMuteWrite)
	ScopeLike     Scope = fmtScope(ScopeLikeRead, ScopeLikeWrite)
	ScopeBlock    Scope = fmtScope(ScopeBlockRead, ScopeBlockWrite)
	ScopeBookmark Scope = fmtScope(ScopeBookmarkRead, ScopeBookmarkWrite)
	ScopeFollowes Scope = fmtScope(ScopeFollowsRead, ScopeFollowsWrite)

	ScopeAll Scope = fmtScope(ScopeTweet,
		ScopeUsersRead,
		ScopeFollowes,
		ScopeOfflineAccess,
		ScopeSpaceRead,
		ScopeMute,
		ScopeLike,
		ScopeBlock,
		ScopeBookmark,
	)
)

func (s Scope) Split() []Scope {
	ss := strings.Split(string(s), " ")
	scope := make([]Scope, len(ss))
	for i, s := range ss {
		scope[i] = Scope(s)
	}
	return scope
}

func (s Scope) Remove(x Scope) Scope {
	for _, x := range x.Split() {
		s = Scope(strings.Replace(string(s), string(x), "", -1))
	}
	return Scope(strings.TrimSpace(string(s)))
}
