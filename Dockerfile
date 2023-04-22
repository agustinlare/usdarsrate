FROM golang:1.17.8-alpine3.15

ENV ENDPOINT_URL=https://www.dolarito.ar/api/frontend/quotations

RUN mkdir /usdrate
COPY . /usdrate
WORKDIR /usdrate

RUN go mod download
RUN go build

USER 1001

CMD [ "./main" ]