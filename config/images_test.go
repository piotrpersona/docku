package config

import (
	"testing"
)

func TestChangeRegistry(t *testing.T) {
	testCases := []struct {
		url, expected ImageURL
		registry      string
	}{
		{
			url:      ImageURL("nginx:latest"),
			registry: "localhost:5000",
			expected: ImageURL("localhost:5000/nginx:latest"),
		},
		{
			url:      ImageURL("docker.io/nginx:latest"),
			registry: "localhost:5000",
			expected: ImageURL("localhost:5000/nginx:latest"),
		},
		{
			url:      ImageURL("docker.io/aaa/nginx:latest"),
			registry: "localhost:5000",
			expected: ImageURL("localhost:5000/aaa/nginx:latest"),
		},
	}
	for _, tc := range testCases {
		result := ImageURL(tc.url).ReplaceRegistry(tc.registry)
		if result != tc.expected {
			t.Errorf("Expected: '%s' have: '%s'", tc.expected, result)
		}
	}
}

func TestBuildURL(t *testing.T) {
	testCases := []struct {
		registry, imageName, expected string
	}{
		{
			registry:  "localhost:5000",
			imageName: "nginx:latest",
			expected:  "localhost:5000/nginx:latest",
		},
		{
			registry:  "//localhost:5000//",
			imageName: "nginx:latest",
			expected:  "localhost:5000/nginx:latest",
		},
		{
			registry:  "localhost:5000",
			imageName: "//nginx:latest//",
			expected:  "localhost:5000/nginx:latest",
		},
		{
			registry:  "//localhost:5000//",
			imageName: "//nginx:latest//",
			expected:  "localhost:5000/nginx:latest",
		},
	}
	for _, tc := range testCases {
		result := buildURL(tc.registry, tc.imageName)
		if result != tc.expected {
			t.Errorf("Expected: '%s' have: '%s'", tc.expected, result)
		}
	}
}
