package service

import (
	"testing"

	"user/internal/model"
)

func TestParseCachedUserReturnsUserForValidJSON(t *testing.T) {
	user, ok := parseCachedUser(`{"id":18,"username":"day18","nickname":"缓存用户"}`)

	if !ok {
		t.Fatal("expected valid cached user")
	}
	if user.ID != 18 || user.Username != "day18" || user.Nickname != "缓存用户" {
		t.Fatalf("unexpected cached user: %#v", user)
	}
}

func TestParseCachedUserRejectsInvalidJSON(t *testing.T) {
	user, ok := parseCachedUser(`{bad json}`)

	if ok {
		t.Fatal("expected invalid cache value to be rejected")
	}
	if user != nil {
		t.Fatalf("expected nil user for invalid cache value, got %#v", user)
	}
}

func TestMarshalCachedUserReturnsJSON(t *testing.T) {
	cacheBytes, err := marshalCachedUser(&model.User{Username: "day17"})

	if err != nil {
		t.Fatalf("expected marshal success, got %v", err)
	}
	if len(cacheBytes) == 0 {
		t.Fatal("expected non-empty cache bytes")
	}
}
