FROM node:23-slim

RUN apt update && apt install -y curl git

RUN curl -fsSL https://golang.org/dl/go1.24.1.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"

ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable && corepack prepare pnpm@10.7.0 --activate

COPY package.json pnpm-lock.yaml buf.gen.yaml ./

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

COPY . .

RUN pnpm buf:generate

WORKDIR /application