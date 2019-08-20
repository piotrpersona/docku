# docku

[![Build Status](https://travis-ci.org/piotrpersona/docku.svg?branch=master)](https://travis-ci.org/piotrpersona/docku)
[![Go Report Card](https://goreportcard.com/badge/github.com/piotrpersona/docku)](https://goreportcard.com/report/github.com/piotrpersona/docku)

Upload images to remote registry at the speed of light ⚡️

![docku](images/docku.png)

## Configuration

Provide images config

```json
{
  "registry": "localhost:5000",
  "images": {
    "nginx": {
      "registry": "docker.io",
      "tag": "latest"
    },
    "alpine": {
      "registry": "docker.io",
      "tag": "3.9"
    }
  }
}
```

## Run

Docker

```bash
docker run \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    --volume "${config}:/config.json" \
    --network host \
    piotrpersona/docku config.json "${@}"
```

## Installation

* docker
* as a package (TBD)
