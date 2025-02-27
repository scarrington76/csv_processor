FROM node:16.20-alpine3.18 AS JS_BUILD
COPY web /web
WORKDIR /web
RUN npm install && npm run build

FROM golang:1.22.2-alpine3.18 AS GO_BUILD
COPY api /api
WORKDIR /api
RUN go build -o /go/bin/api

FROM alpine:3.18.6
COPY --from=JS_BUILD /web/build* ./webapp/
COPY --from=GO_BUILD /go/bin/api ./
CMD ./api
