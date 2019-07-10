package docker

func imageURL(registry, name, tag string) string {
	return registry + "/" + name + ":" + tag
}
