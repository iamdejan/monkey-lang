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
