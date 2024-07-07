package model

import (
	"time"

	"github.com/google/uuid"
)

type UserID = uuid.UUID
type User struct {
	ID           UserID    `json:"id"`
	Username     string    `json:"username"`
	DisplayName  string    `json:"displayName"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	AvatarURL    string    `json:"avatarUrl"`
	IsBot        bool      `json:"isBot"`
	CreatedAt    time.Time `json:"createdAt"`
	LastActiveAt time.Time `json:"lastActiveAt"`
}

type GuildID = uuid.UUID
type Guild struct {
	ID          GuildID   `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     UserID    `json:"ownerId"`
	IconURL     string    `json:"iconUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type RoleID = uuid.UUID
type Role struct {
	ID          RoleID    `json:"id"`
	GuildID     GuildID   `json:"guildId"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Permissions uint64    `json:"permissions"`
	Mentionable bool      `json:"mentionable"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type MemberID = uuid.UUID
type Member struct {
	ID        MemberID  `json:"id"`
	GuildID   GuildID   `json:"guildId"`
	UserID    UserID    `json:"userId"`
	Nickname  string    `json:"nickname"`
	JoinedAt  time.Time `json:"joinedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryID = uuid.UUID
type Category struct {
	ID        CategoryID `json:"id"`
	GuildID   GuildID    `json:"guildId"`
	Name      string     `json:"name"`
	Position  int        `json:"position"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type ChannelID = uuid.UUID
type ChannelType string

const (
	ChannelTypeText     ChannelType = "text"
	ChannelTypeVoice    ChannelType = "voice"
	ChannelTypeAnnounce ChannelType = "announce"
)

type Channel struct {
	ID         ChannelID   `json:"id"`
	GuildID    GuildID     `json:"guildId"`
	CategoryID CategoryID  `json:"categoryId"`
	Name       string      `json:"name"`
	Type       ChannelType `json:"type"`
	Position   int         `json:"position"`
	Topic      string      `json:"topic"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

type MessageID = uuid.UUID
type Message struct {
	ID        MessageID `json:"id"`
	ChannelID ChannelID `json:"channelId"`
	AuthorID  UserID    `json:"authorId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Reaction struct {
	MessageID MessageID `json:"messageId"`
	UserID    UserID    `json:"userId"`
	Emoji     string    `json:"emoji"`
	CreatedAt time.Time `json:"createdAt"`
}

type Invite struct {
	Code      string    `json:"code"`
	GuildID   GuildID   `json:"guildId"`
	CreatorID UserID    `json:"creatorId"`
	ExpiresAt time.Time `json:"expiresAt"`
	MaxUses   int       `json:"maxUses"`
	Uses      int       `json:"uses"`
	CreatedAt time.Time `json:"createdAt"`
}
