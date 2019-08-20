# docku

[![Build Status](https://travis-ci.org/piotrpersona/docku.svg?branch=master)](https://travis-ci.org/piotrpersona/docku)
[![Go Report Card](https://goreportcard.com/badge/github.com/piotrpersona/docku)](https://goreportcard.com/report/github.com/piotrpersona/docku)

Upload images to remote registry at the speed of light ⚡️

![docku-arch](https://raw.githubusercontent.com/piotrpersona/docku/master/svg/docku-arch.svg?sanitize=true)

## Installation

Download latest release from:
https://github.com/piotrpersona/docku/releases

```bash
curl -fsSL -o /usr/local/bin/docku https://github.com/piotrpersona/docku/releases/download/<RELEASE>/docku-<OS>-<ARCH> && chmod +x /usr/local/bin/docku
```

e.g.:

```bash
curl -o /usr/local/bin/docku https://github.com/piotrpersona/docku/releases/download/v1.0.2/docku-darwin-amd64 && chmod +x /usr/local/bin/docku
```

## Run

```bash
docku config.json
```

Docker (TBD)

```bash
docker run \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    --volume "${config}:/config.json" \
    --network host \
    piotrpersona/docku config.json "${@}"
```

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
