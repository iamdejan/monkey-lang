# Monkey Language

My implementation from "Writing An Interpreter in Go".

## Prerequisites

Before you can try this, there are some things you need to setup.

### Go Language

You need to install Go language.

#### Ubuntu

```sh
sudo apt install golang-go
```

#### Mac

```sh
brew install go
```

#### Windows 10 and later

##### WSL (Recommended)

I suggest you install WSL, install Ubuntu in WSL (I use 22.04), then follow the guide for Ubuntu.

##### Native

Follow the guide from [Go Language official website](https://go.dev/doc/install).

### Pre-Commit Hook

#### Ubuntu

First, you need to install Python 3 and PIP. Then, install `pre-commit` program.
```sh
sudo apt install python3 python3-pip
pip3 install pre-commit
```

Close the terminal. Then, test whether you can execute `pre-commit` or not.
```sh
pre-commit --version
```

If it shows error, try [following this guide](https://stackoverflow.com/a/71043830).
