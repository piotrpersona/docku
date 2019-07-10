package docker_test

import (
	"testing"

	"github.com/piotrpersona/docker-upload/docker"
)

var testCases = []struct {
	registry, name, tag, expected string
}{
	{
		registry: "docker.io",
		name:     "nginx",
		tag:      "latest",
		expected: "docker.io/nginx:latest",
	},
	{
		registry: "docker.io/",
		name:     "/nginx",
		tag:      "latest",
		expected: "docker.io/nginx:latest",
	},
	{
		registry: "docker.io///",
		name:     "//nginx",
		tag:      "latest",
		expected: "docker.io/nginx:latest",
	},
	{
		registry: "localhost:8080",
		name:     "library/nginx",
		tag:      "latest",
		expected: "localhost:8080/library/nginx:latest",
	},
	{
		registry: "localhost:8080",
		name:     "library/nginx::",
		tag:      "::latest",
		expected: "localhost:8080/library/nginx:latest",
	},
}

func TestImageURL(t *testing.T) {
	for _, tc := range testCases {
		actual := docker.ImageURL(tc.registry, tc.name, tc.tag)
		if actual != tc.expected {
			t.Errorf("Expected %s got %s", tc.expected, actual)
		}
	}
}
