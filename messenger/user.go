package messenger

import (
	"context"
	"encoding/json"
)

// User messenger user entity
type User struct {
	UserID         string `json:"user_id"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name,omitempty"`
	AvatarURL      string `json:"avatar_url,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	Phone          string `json:"phone,omitempty"`
}

// FetchProfile fetch my profile
func (m Messenger) FetchProfile(ctx context.Context) (*User, error) {
	data, err := m.Request(ctx, "GET", "/me", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.User, err
}

// ModifyProfile update my profile
func (m Messenger) ModifyProfile(ctx context.Context, fullname, avatarBase64 string) (*User, error) {
	paras := map[string]interface{}{}
	if len(fullname) > 0 {
		paras["full_name"] = fullname
	}
	if len(avatarBase64) > 0 {
		paras["avatar_base64"] = avatarBase64
	}
	payload, err := json.Marshal(paras)
	if err != nil {
		return nil, err
	}

	data, err := m.Request(ctx, "POST", "/me", payload)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.User, err
}

// ModifyPreference update my preference
func (m Messenger) ModifyPreference(ctx context.Context, receiveMessageSource, acceptConversationSource string) (*User, error) {
	paras := map[string]interface{}{}
	if len(receiveMessageSource) > 0 {
		paras["receive_message_source"] = receiveMessageSource
	}
	if len(acceptConversationSource) > 0 {
		paras["accept_conversation_source"] = acceptConversationSource
	}
	payload, err := json.Marshal(paras)
	if err != nil {
		return nil, err
	}

	data, err := m.Request(ctx, "POST", "/me/preferences", payload)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.User, err
}

// FetchUsers fetch users
func (m Messenger) FetchUsers(ctx context.Context, userIDS ...string) ([]User, error) {
	if len(userIDS) == 0 {
		return nil, nil
	}

	payload, err := json.Marshal(userIDS)
	if err != nil {
		return nil, err
	}

	data, err := m.Request(ctx, "POST", "/users/fetch", payload)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Users []User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return resp.Users, nil
}

// FetchUser fetch user
func (m Messenger) FetchUser(ctx context.Context, userID string) (*User, error) {
	data, err := m.Request(ctx, "GET", "/users/"+userID, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.User, err
}

// SearchUser search user; q is String: Mixin Id or Phone Numbe
func (m Messenger) SearchUser(ctx context.Context, q string) (*User, error) {
	data, err := m.Request(ctx, "GET", "/search/"+q, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		User User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.User, err
}

// FetchFriends fetch friends
func (m Messenger) FetchFriends(ctx context.Context) ([]User, error) {
	data, err := m.Request(ctx, "GET", "/friends", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Users []User `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return resp.Users, err
}
