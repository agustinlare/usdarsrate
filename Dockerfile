FROM golang:1.17.8-alpine3.15

RUN mkdir -p dolarchat
COPY . dolarchat
WORKDIR dolarchat

RUN go mod download
RUN go build

USER 1001

CMD [ "./main" ]