package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	db := UserList{}

	rand.Seed(4452)
	for i := 0; i < 256; i++ {
		id := rand.Int63()
		db = append(db, User{ID: id})
	}

	proxy := &UserListProxy{
		DB:            db,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knownIDs := [3]int64{db[3].ID, db[4].ID, db[5].ID}

	t.Run("FindUser - Empty Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - Two more user from db", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		user2, _ := proxy.FindUser(knownIDs[1])
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		user3, _ := proxy.FindUser(knownIDs[2])
		if proxy.DidLastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow" +
				" more than to two")
		}

		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
