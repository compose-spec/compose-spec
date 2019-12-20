FROM node:12 as generator
RUN npm install -g @adobe/jsonschema2md
COPY schema /schema
RUN jsonschema2md -d /schema/ -e json

FROM scratch as docs
COPY --from=generator /out /


FROM golang:1.13 as codegen
RUN go get -u github.com/mjibson/esc
COPY schema /schema
RUN esc -o /schema.go -pkg spec /schema/


FROM scratch as golang
COPY --from=codegen /schema.go /
