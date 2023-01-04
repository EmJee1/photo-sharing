# Photo Sharing

<!-- Describe where this readme is about.  -->

This is a project made for a CHE assignment (university of applied sciences).

The repo contains the website & backend for the IP Development assignment I made. The website is made with Go,
echo and goview.

## ✨ See it in action ✨

| Environment  | Web              |
|--------------|------------------|
| `Production` | Not deployed yet |

<!--
## 🧰 External tooling

- [Figma for the designs](https://www.figma.com/file/2BVrpjljMfNUsaEuTBbsLU/CHE-FDD?node-id=0%3A1&t=JneFtIXnLIle8Axq-1)
- [Firebase for the backend](https://console.firebase.google.com/u/0/project/che-fdd-assignment)
-->

## 🚀 Getting started

### Prerequisites

<!-- 
   Which software or library's are needed to be able to install this project?
 -->

- [Goland](https://www.jetbrains.com/go/) or [Visual Studio Code](https://code.visualstudio.com/)
- The [Go binary](https://go.dev/dl/)
- Docker ([docker desktop recommended](https://www.docker.com/products/docker-desktop/))

### Installation

<!-- How to install this project (after having the prerequisites)? -->

```shell
# Setup and run postgresql & pgadmin containers
docker compose up -d

# Install the project dependencies
go get
```

### Development

```shell
go run main.go
```

### Deployment / Release process

Not deployed yet

<!--
## 🤚 Good to know

- We use Dependabot for automatic package updates
- This creates pull requests every monday, for all npm packages that are outdated
-->
