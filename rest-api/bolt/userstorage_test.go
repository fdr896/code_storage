package bolt

import (
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
)

var (
	validUser1 = &core.User{
		Login: "admin",
	}
	validUser2 = &core.User{
		Login: "test",
	}
	badUser = &core.User{
		Login: "    ",
	}
	testPassword1 = "qwerty"
	testPassword2 = "123"
)

func TestNewUserStorage(t *testing.T) {
	os.RemoveAll(testPath)
	db, _ := bolt.Open(testPath, 0600, nil)
	defer db.Close()

	_, err := NewUserStorage(testBucketName, testPath, db)
	if err != nil {
		failTest(t, err)
	}
}

func TestAddUser(t *testing.T) {
	os.RemoveAll(testPath)
	db, _ := bolt.Open(testPath, 0600, nil)
	defer db.Close()

	us, _ := NewUserStorage(testBucketName, testPath, db)

	if err := us.AddUser(validUser1); err != nil {
		failTest(t, err)
	}

	if err := us.AddUser(validUser1); err == nil {
		t.Error("userStorage didn't handle registretaion of already registered user")
	}
}

func TestSetPassword(t *testing.T) {
	os.RemoveAll(testPath)
	db, _ := bolt.Open(testPath, 0600, nil)
	defer db.Close()

	us, _ := NewUserStorage(testBucketName, testPath, db)

	us.AddUser(validUser1)
	if err := us.SetPassword(validUser1.Login, testPassword1); err != nil {
		failTest(t, err)
	}

	if err := us.SetPassword(validUser1.Login, ""); err == nil {
		t.Errorf("userStorage didn't handle empty password")
	}
}

func TestHasPassword(t *testing.T) {
	os.RemoveAll(testPath)
	db, _ := bolt.Open(testPath, 0600, nil)
	defer db.Close()

	us, _ := NewUserStorage(testBucketName, testPath, db)

	if _, err := us.HasPassword(validUser1.Login); err == nil {
		t.Error("userStorage didn't handle unregistered user properly")
	}

	us.AddUser(validUser1)
	if hasPassword, err := us.HasPassword(validUser1.Login); hasPassword || err != nil {
		if err != nil {
			failTest(t, err)
		} else {
			t.Error("userStorage claimed that user is already set password")
		}
	}

	us.SetPassword(validUser1.Login, testPassword1)
	if hasPassword, err := us.HasPassword(validUser1.Login); !hasPassword || err != nil {
		if err != nil {
			failTest(t, err)
		} else {
			t.Error("userStorage claimed that hadn't set password")
		}
	}
}

func TestComparePassword(t *testing.T) {
	os.RemoveAll(testPath)
	db, _ := bolt.Open(testPath, 0600, nil)
	defer db.Close()

	us, _ := NewUserStorage(testBucketName, testPath, db)

	if _, err := us.ComparePassword(validUser1.Login, testPassword1); err == nil {
		t.Error("userStorage didn't handle non existing user when compare password")
	}

	us.AddUser(validUser1)
	us.SetPassword(validUser1.Login, testPassword1)
	if rightPassword, err := us.ComparePassword(validUser1.Login, testPassword1); !rightPassword || err != nil {
		if err != nil {
			failTest(t, err)
		} else {
			t.Error("userStorage claimed that correct password is incorrect")
		}
	}

	if rightPassword, err := us.ComparePassword(validUser1.Login, testPassword2); rightPassword || err != nil {
		if err != nil {
			failTest(t, err)
		} else {
			t.Error("userStorage claimed that incorrect password is correct")
		}
	}
}
