FROM golang:1.23.10

# Make Apt non-interactive
RUN echo 'APT::Get::Assume-Yes "true";' > /etc/apt/apt.conf.d/90ci \
  && echo 'DPkg::Options "--force-confnew";' >> /etc/apt/apt.conf.d/90ci

ENV DEBIAN_FRONTEND=noninteractive

# Debian Jessie is EOL'd and original repos don't work.
# Switch to the archive mirror until we can get people to
# switch to Stretch.
RUN if grep -q Debian /etc/os-release && grep -q jessie /etc/os-release; then \
	rm /etc/apt/sources.list \
    && echo "deb http://archive.debian.org/debian/ jessie main" >> /etc/apt/sources.list \
    && echo "deb http://security.debian.org/debian-security jessie/updates main" >> /etc/apt/sources.list \
	; fi

# Make sure PATH includes ~/.local/bin
# https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=839155
RUN echo 'PATH="$HOME/.local/bin:$PATH"' >> /etc/profile.d/user-local-path.sh

# man directory is missing in some base images
# https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=863199
RUN apt-get update \
  && mkdir -p /usr/share/man/man1 \
  && apt-get install -y \
    git apt \
    locales sudo openssh-client ca-certificates tar gzip \
    net-tools netcat-traditional unzip zip bzip2 gnupg curl wget jq

# Set timezone to UTC by default
RUN ln -sf /usr/share/zoneinfo/Etc/UTC /etc/localtime

# Use unicode
RUN locale-gen C.UTF-8 || true
ENV LANG=C.UTF-8

# Install Docker
RUN curl -fsSL https://get.docker.com|bash

RUN groupadd --gid 3434 ci \
  && useradd --uid 3434 --gid ci --shell /bin/bash --create-home ci \
  && echo 'ci ALL=NOPASSWD: ALL' >> /etc/sudoers.d/50-ci \
  && echo 'Defaults    env_keep += "DEBIAN_FRONTEND"' >> /etc/sudoers.d/env_keep

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | INSTALL_DIRECTORY=/usr/local/bin sh
RUN curl -L https://github.com/gotestyourself/gotestsum/releases/download/v0.3.4/gotestsum_0.3.4_linux_amd64.tar.gz | \
  tar -xz -C /usr/local/bin gotestsum
RUN go install golang.org/x/lint/golint@latest
RUN go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.5
RUN go install sigs.k8s.io/kustomize/kustomize/v4@latest
RUN go install github.com/mikefarah/yq/v4@v4.44.6

# Install kubectl cli
RUN curl -o /usr/local/bin/kubectl -L https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x /usr/local/bin/kubectl

# Install helm client
RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

# Install Operator SDK
ARG OPERATOR_SDK_VERSION
RUN curl -o /usr/local/bin/operator-sdk -L https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk_linux_amd64 \
    && chmod +x /usr/local/bin/operator-sdk

WORKDIR /go/casskop

CMD ["/bin/sh"]
