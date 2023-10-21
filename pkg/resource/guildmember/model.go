package guildmember

import "github.com/dezhishen/satori-sdk-go/pkg/resource/user"

type GuildMember struct {
	User   *user.User `json:"user"`
	Nick   string     `json:"nick"`
	Avatar string     `json:"avatar"`
	JoinAt int64      `json:"join_at"`
}

type GuildMemberList struct {
	Data []GuildMember `json:"data"`
	Next string        `json:"next,omitempty"`
}
