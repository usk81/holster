workflow "Build" {
  on = "release"
  resolves = [
    "release darwin/386",
    "release darwin/amd64",
    "release linux/386",
    "release linux/amd64",
    "release windows/386",
    "release windows/amd64",
  ]
}

action "release darwin/386" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "darwin"
    GOARCH = "386"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release darwin/amd64" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release linux/386" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "linux"
    GOARCH = "386"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release linux/amd64" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release windows/386" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "windows"
    GOARCH = "386"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release windows/amd64" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}