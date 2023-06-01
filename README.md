# Monkey Language

My implementation from "Writing An Interpreter in Go".

## Prerequisites

Before you can try this, there are some things you need to setup. The setup will be different, depending on the platform.

### Ubuntu

#### Go Language

```sh
sudo apt install golang-go
```

#### Pre-commit Hook

First, you need to install Python 3 and Pip. Then, install `pre-commit` program.
```sh
sudo apt install python3 python3-pip
pip3 install pre-commit
```

Close the terminal. Then, test whether you can execute `pre-commit` or not.
```sh
pre-commit --version
```

If it shows error, try [following this guide](https://stackoverflow.com/a/71043830).

### Windows 10 and 11

I suggest you install WSL 2, install Ubuntu in WSL (I use 22.04), then follow the guide for Ubuntu.

Below guide(s) and link(s) are for you who want to use native Windows platform.

#### Go Language

Follow the guide from [Go Language official website](https://go.dev/doc/install).

#### Pre-commit Hook

You can install Python by downloading the setup from [here](https://www.python.org/downloads/windows/). Pip is also included in the setup.

Reference: https://stackoverflow.com/a/12476379

## How to Run

The easiest way to try this is to run the REPL:
```sh
go run main.go
```

However, the *slight* inconvenience is that you cannot use arrow keys, since that will only return escape sequence(s) of the key (e.g. `^[[D`). To solve that issue, you can use [rlfe](https://manpages.debian.org/unstable/rlfe/rlfe.1.en.html) like this:
```sh
rlfe go run main.go
```

However, I'm not sure if `rlfe` exists on Mac OS. I only tested this on Ubuntu 22.04.

## Pre-Commit Hook

Everytime you create / edit `.pre-commit-config.yaml` file, don't forget to run this command:
```sh
pre-commit install --hook-type pre-push
```

NOTE: to ease the life of the developer, we need to add a little bit of leniency for the validation, which is why the validation is done before push, not before commit. The commit will be squashed anyway, so it's okay to validate before push.
