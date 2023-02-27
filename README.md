# JWT command line interface tool

<!-- [![License](https://img.shields.io/github/license/bufbuild/buf?color=blue)][badges_license]
[![Release](https://img.shields.io/github/v/release/bufbuild/buf?include_prereleases)][badges_release]
[![CI](https://github.com/bufbuild/buf/workflows/ci/badge.svg)][badges_ci]
[![Docker](https://img.shields.io/docker/pulls/bufbuild/buf)][badges_docker]
[![Homebrew](https://img.shields.io/badge/homebrew-v1.19.0-blue)][badges_homebrew]
[![AUR](https://img.shields.io/aur/version/buf)][badges_aur]
[![Slack](https://img.shields.io/badge/slack-buf-%23e01563)][badges_slack] -->

A super fast CLI tool to decode, encode, and verify JWTs built in [Go](https://go.dev).


`jwt` is a command line tool to help you work with JSON Web Tokens (JWTs). Like most JWT command line tools out there, you can decode almost any JWT header and claims body. Unlike any that I've found, however, `jwt` allows you to encode a new JWT with nearly any piece of data you can think of. Custom header values (some), custom claim bodies (as long as it's JSON, it's game), and using any secret you need.

On top of all that, it's written in Go so it's fast and portable (windows, macOS, and linux supported right now).

## Installation

### Homebrew

You can install `jwt` using [Homebrew][brew] (macOS or Linux):

```sh
brew install duckhue01/duckhue01/jwt
```

### Binary

JWT offers Windows binaries for both the x86_64 and arm64 architectures. You can download the latest binaries from  [GitHub Releases](https://github.com/duckhue01/jwt/releases).

## Usage

For usage info, use the `help` command.

```sh
# top level help
jwt help

# command specific help
jwt help decode
```

## Contributing



