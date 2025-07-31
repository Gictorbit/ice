# Ice Global Backend challenge

### Available recipes:
> **Note**: This project uses [`just`](https://github.com/casey/just) as a task runner instead of `make`.  
> Make sure to install `just` before running any project commands.  
> You can install it via [Homebrew](https://formulae.brew.sh/formula/just) (`brew install just`) or follow instructions in the official repo.

```shell
just --list 
Available recipes:
    benchmark # Run benchmarks
    build     # Build the Go binary using the local Go environment (not Docker)
    clean     # Clean up compiled binaries
    docker    # Build Docker image
    generate  # generate go commands and mocks
    run       # Run the whole application stack with Docker Compose
    test      # Run unit tests
```