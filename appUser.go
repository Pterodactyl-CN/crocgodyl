package crocgodyl

import (
	"encoding/json"
	"fmt"
	"time"
)

// Application User API

// Users is the struct for all the panel users.
// GET this from the '/api/application/users` endpoint
type Users struct {
	Object string `json:"object,omitempty"`
	User   []User `json:"data,omitempty"`
	Meta   Meta   `json:"meta,omitempty"`
}

// User is the struct for all the panel users.
// GET this from the '/api/application/users/<user_ID>` endpoint
// The panel does not and will not return a password.
// You can update a password using the API.
type User struct {
	Object     string         `json:"object,omitempty"`
	Attributes UserAttributes `json:"attributes,omitempty"`
}

// UserAttributes is the struct for all the panel users.
type UserAttributes struct {
	ID         int       `json:"id,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`
	UUID       string    `json:"uuid,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Language   string    `json:"language,omitempty"`
	RootAdmin  bool      `json:"root_admin,omitempty"`
	TwoFa      bool      `json:"2fa,omitempty"`
	Password   string    `json:"password,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

// GetUsers returns Information on all users.
func GetUsers() (Users, error) {
	var users Users

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI("users", "get", nil)
	if err != nil {
		return users, err
	}

	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetUser returns Information on a single user.
func GetUser(userID int) (User, error) {
	var user User
	endpoint := fmt.Sprintf("users/%d", userID)

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI(endpoint, "get", nil)
	if err != nil {
		return user, err
	}

	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByExternal returns Information on a single user by their externalID.
// The externalID is a string as that is what the panel requires.
func GetUserByExternal(externalID string) (User, error) {
	var user User
	endpoint := fmt.Sprintf("users/%s", externalID)

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI(endpoint, "get", nil)
	if err != nil {
		return user, err
	}

	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByPage returns Information on users by their page number.
// The externalID is a string as that is what the panel requires.
func GetUserByPage(pageID int) (User, error) {
	var user User
	endpoint := fmt.Sprintf("users/%d", pageID)

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI(endpoint, "get", nil)
	if err != nil {
		return user, err
	}

	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser creates a user.
func CreateUser(newUser UserAttributes) (User, error) {
	var userDetails User

	nubytes, err := json.Marshal(newUser)
	if err != nil {
		return userDetails, err
	}

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI("users", "post", nubytes)
	if err != nil {
		return userDetails, err
	}

	// Get user info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &userDetails)
	if err != nil {
		return userDetails, err
	}

	return userDetails, nil
}

// EditUser creates a user.
// Send a UserAttributes to the panel to update the user.
// You cannot edit the id or created/updated fields for the user.
func EditUser(editUser UserAttributes, userID int) (User, error) {
	var userDetails User
	endpoint := fmt.Sprintf("users/%d", userID)

	eubytes, err := json.Marshal(editUser)
	if err != nil {
		return userDetails, err
	}

	// get json bytes from the panel.
	ubytes, err := queryPanelAPI(endpoint, "patch", eubytes)
	if err != nil {
		return userDetails, err
	}

	// Get user info from the panel
	// Unmarshal the bytes to a usable struct.
	err = json.Unmarshal(ubytes, &userDetails)
	if err != nil {
		return userDetails, err
	}

	return userDetails, nil
}

// DeleteUser deletes a user.
// It only requires a user id as a string
func DeleteUser(userID int) error {
	endpoint := fmt.Sprintf("users/%d", userID)

	// get json bytes from the panel.
	_, err := queryPanelAPI(endpoint, "delete", nil)
	if err != nil {
		return err
	}

	return nil
}
