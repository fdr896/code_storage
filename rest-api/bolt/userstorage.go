package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
)

// UserStorage stores information about all users.
type userStorage struct {
	bucketName []byte
	UsersInfo  *bolt.DB
}

// NewUserStorage creates new UserStorage with required main bucketName in required localtion.
func NewUserStorage(bucketName []byte, path string, db *bolt.DB) (core.UserStorage, error) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	}); err != nil {
		return nil, err
	}

	return &userStorage{bucketName, db}, nil
}

// AddUser adds new user to database if there's no user with same login.
func (us *userStorage) AddUser(user *core.User) error {
	return us.UsersInfo.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(us.bucketName)

		tmp := b.Get([]byte(user.Login))
		if tmp != nil {
			return core.ErrUserWithSameLogin
		}

		encoded, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put([]byte(user.Login), encoded)
	})
}

// HasPassword checks if user set his password.
func (us *userStorage) HasPassword(login string) (bool, error) {
	var userState bool
	if err := us.UsersInfo.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(us.bucketName)

		bytes := b.Get([]byte(login))
		if bytes == nil {
			return core.ErrNoSuchUser
		}

		var user core.User
		json.Unmarshal(bytes, &user)

		userState = (user.Password != "")
		return nil
	}); err != nil {
		return false, err
	}

	return userState, nil
}

// ComparePassword checks if received password is similar with real user's password.
func (us *userStorage) ComparePassword(login string, receivedPassword string) (bool, error) {
	var correctPassword bool
	if err := us.UsersInfo.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(us.bucketName)

		bytes := b.Get([]byte(login))
		if bytes == nil {
			return core.ErrNoSuchUser
		}

		var user core.User
		json.Unmarshal(bytes, &user)

		correctPassword = (receivedPassword == user.Password)
		return nil
	}); err != nil {
		return false, err
	}

	return correctPassword, nil
}

// SetPassword sets new password to user with required login.
func (us *userStorage) SetPassword(login string, newPassword string) error {
	if newPassword == "" {
		return core.ErrEmptyPassword
	}

	return us.UsersInfo.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(us.bucketName)

		bytes := b.Get([]byte(login))
		if bytes == nil {
			return core.ErrNoSuchUser
		}

		var user core.User
		json.Unmarshal(bytes, &user)

		user.Password = newPassword

		encoded, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put([]byte(login), encoded)
	})
}
