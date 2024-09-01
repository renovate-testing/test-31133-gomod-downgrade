# 31133

First, read the [Renovate minimal reproduction instructions](https://github.com/renovatebot/renovate/blob/main/docs/development/minimal-reproductions.md).

Then replace the current `h1` with the Renovate Issue/Discussion number.

## Current behavior

When a Go module dependency is declared at a pseudo version (e.g. the latest commit on a branch), and there are tagged releases on older commits, Renovate downgrades the dependency to the latest tag on an older commit.

## Expected behavior

Renovate updates the dependency only when there are newer commits or tags.

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/discussions/31133
