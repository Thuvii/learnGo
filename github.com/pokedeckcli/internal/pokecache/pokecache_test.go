package pokecache

import (
	"testing"
	"time"
)

func TestGetAdd(t *testing.T) {
	cases := []struct {
		input    string
		expected []byte
	}{
		{
			input:    "something",
			expected: []byte("hello this is the first test"),
		},
		{
			input:    "null",
			expected: []byte("this is a null key"),
		},
	}

	for _, e := range cases {
		cache := NewCache(5 * time.Second)
		cache.Add(e.input, e.expected)
		val, exist := cache.Get(e.input)
		if !exist {
			t.Errorf("expected to find key %s", e.input)
			return
		}
		if string(val) != string(e.expected) {
			t.Errorf("Expected to find same output: %v", e.expected)
		}
	}

}

func TestReapLoop(t *testing.T) {

	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("something", []byte("hello this is the first test"))

	_, exist := cache.Get("something")
	if !exist {
		t.Errorf("Expected to find key")
		return
	}
	time.Sleep(waitTime)

	_, exist = cache.Get("something")
	if exist {
		t.Errorf("Expected not to find key")
		return
	}

}
