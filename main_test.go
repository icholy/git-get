package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	tests := []struct {
		url  string
		path string
	}{
		{
			url:  "https://github.com/faiface/pixel.git",
			path: "github.com/faiface/pixel",
		},
		{
			url:  "git@github.com:icholy/rtsp.git",
			path: "github.com/icholy/rtsp",
		},
		{
			url:  "ssh://git@stash.company.com/project/name.git",
			path: "stash.company.com/project/name",
		},
		{
			url:  "git@gitlab.com:braniii/prettypyplot.git",
			path: "gitlab.com/braniii/prettypyplot",
		},
		{
			url:  "https://gitlab.com/braniii/prettypyplot.git",
			path: "gitlab.com/braniii/prettypyplot",
		},
		{
			url:  "https://git.sr.ht/~admicos/huedance",
			path: "git.sr.ht/admicos/huedance",
		},
		{
			url:  "http://user@stash.company.com/scm/project/repo.git",
			path: "stash.company.com/project/repo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			u, err := ParseURL(tt.url)
			if err != nil {
				t.Fatalf("failed to parse url: %v", err)
			}
			if path := Path(u); path != tt.path {
				t.Fatalf("expected %q, got %q", tt.path, path)
			}
		})
	}
}
