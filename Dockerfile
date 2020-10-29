FROM node:13-alpine AS front-build

RUN apk add python make gcc g++
WORKDIR /frontend
COPY ./frontend .
RUN yarn install 
RUN yarn run generate

FROM golang:1.15-alpine AS build
MAINTAINER Alejandro Garrido "garridomota@gmail.com"

WORKDIR /librefit 
COPY . .
COPY --from=front-build /frontend/dist ./static
RUN apk add gcc libc-dev
RUN go build --tags "json1" -o ./librefit ./main.go

FROM alpine:latest AS prod
WORKDIR /librefit
COPY --from=build /librefit/librefit /librefit/librefit
COPY --from=front-build /frontend/dist ./static
RUN mkdir /librefit/data
EXPOSE 4000
ENTRYPOINT ["/librefit/librefit", "start"]
