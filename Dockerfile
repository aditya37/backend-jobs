FROM golang

RUN mkdir /injobs

WORKDIR /injobs/


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go","run","main.go"]
