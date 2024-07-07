-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    display_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    is_bot BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    last_active_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Guilds table
CREATE TABLE guilds (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(id),
    icon_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    guild_id UUID NOT NULL REFERENCES guilds(id),
    name VARCHAR(255) NOT NULL,
    color VARCHAR(7),
    permissions BIGINT NOT NULL,
    mentionable BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Members table
CREATE TABLE members (
    id UUID PRIMARY KEY,
    guild_id UUID NOT NULL REFERENCES guilds(id),
    user_id UUID NOT NULL REFERENCES users(id),
    nickname VARCHAR(255),
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    UNIQUE (guild_id, user_id)
);

-- Categories table
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    guild_id UUID NOT NULL REFERENCES guilds(id),
    name VARCHAR(255) NOT NULL,
    position INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Channels table
CREATE TABLE channels (
    id UUID PRIMARY KEY,
    guild_id UUID NOT NULL REFERENCES guilds(id),
    category_id UUID REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('text', 'voice', 'announce')),
    position INTEGER NOT NULL,
    topic TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY,
    channel_id UUID NOT NULL REFERENCES channels(id),
    author_id UUID NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Reactions table
CREATE TABLE reactions (
    message_id UUID NOT NULL REFERENCES messages(id),
    user_id UUID NOT NULL REFERENCES users(id),
    emoji VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (message_id, user_id, emoji)
);

-- Invites table
CREATE TABLE invites (
    code VARCHAR(255) PRIMARY KEY,
    guild_id UUID NOT NULL REFERENCES guilds(id),
    creator_id UUID NOT NULL REFERENCES users(id),
    expires_at TIMESTAMP WITH TIME ZONE,
    max_uses INTEGER,
    uses INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);