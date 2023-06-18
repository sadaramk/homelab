ARG VARIANT=16-bullseye
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
	bash-completion \
	figlet \
    make \
    python3-pip \
    python3-venv \
    default-jdk \
    graphviz \
    wget \
    # golang \
    && rm -rf /var/lib/apt/lists/*

RUN update-ca-certificates

# env/vscode support: LANG must be supported, requires installing the locale package first
# https://github.com/Microsoft/vscode/issues/58015
# https://stackoverflow.com/questions/28405902/how-to-set-the-locale-inside-a-debian-ubuntu-docker-container
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

ENV LANG en_US.UTF-8

# go gotestsum: (this package should NOT be installed via go get)
# https://github.com/gotestyourself/gotestsum/releases
RUN mkdir -p /tmp/gotestsum \
    && cd /tmp/gotestsum \
    && wget https://github.com/gotestyourself/gotestsum/releases/download/v1.8.0/gotestsum_1.8.0_linux_amd64.tar.gz \
    && tar xzf gotestsum_1.8.0_linux_amd64.tar.gz \
    && cp gotestsum /usr/local/bin/gotestsum \
    && rm -rf /tmp/gotestsum

# go linting: (this package should NOT be installed via go get)
# https://github.com/golangci/golangci-lint#binary
# https://github.com/golangci/golangci-lint/releases
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b $(go env GOPATH)/bin v1.45.2




# Upgrade pip
RUN pip3 install --progress-bar off --upgrade pip

RUN pip3 --disable-pip-version-check --no-cache-dir install ansible pre-commit


RUN wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | sudo apt-key add - && \
    echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list && \
    sudo apt-get update && \
    sudo apt-get install -y mongodb-mongosh mongodb-database-tools

RUN pip3 --disable-pip-version-check --no-cache-dir install ansible pre-commit

# Gomplate
ARG GOMPLATE_VERSION=3.10.0
RUN mkdir -p /tmp/download \
	&& wget https://github.com/hairyhenderson/gomplate/releases/download/v${GOMPLATE_VERSION}/gomplate_linux-amd64-slim -O /tmp/download/gomplate --quiet --no-check-certificate \
	&& chmod +x /tmp/download/gomplate \
	&& mv /tmp/download/gomplate /usr/local/bin/ \
	&& rm -rf /tmp/download


# Install Go 1.20
RUN wget  https://go.dev/dl/go1.20.5.linux-arm64.tar.gz
RUN tar -C /usr/local -xzf go1.20.5.linux-arm64.tar.gz
RUN rm -rf go1.20.5.linux-arm64.tar.gz

RUN export GOPATH=/usr/local/go
RUN export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin

# RUN go version

# Install Go-Task
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d