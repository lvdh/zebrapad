# Contributing to zebrapad

Thank you for your interest in contributing to zebrapad!

## Getting Started

- Fork the repository.
- Create a new branch from `main`.
- Ensure that, on your machine, the project:
  - runs: `make run`
  - builds: `make build`
  - passes all tests: `make test`
- Read and understand the [Code of Conduct].
- Develop your change.
- Send your [Pull Request].

## Commit Messages

Ideally, all commits adhere to [Conventional Commits]. Examples:

```
feat(label): add preview endpoint
fix(printer): handle connection errors
docs: update README with setup instructions
```

## Code Style

- backend
  - Write tests for new features and bugfixes.
  - Go code should be formatted with `gofmt`.
- frontend
  - Static assets are served from `internal/server/web/`.

## Pull Requests

- Ensure your branch is up to date with `main`.
- Briefly describe your changes and reference related issues, if any.
- Linting and tests must pass before merging.

## Questions or Suggestions?

Open an issue or start a discussion. Constructive feedback and ideas are welcome!

[Code of Conduct]: CODE_OF_CONDUCT.md
[Pull Request]: https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests
[Conventional Commits]: https://www.conventionalcommits.org/
