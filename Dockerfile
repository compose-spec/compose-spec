# syntax=docker/dockerfile:1

FROM node:16.1 as generator
RUN npm install -g @adobe/jsonschema2md
COPY schema /schema
RUN jsonschema2md -d /schema/ -e json

FROM scratch as docs
COPY --from=generator /out /

FROM --platform=${BUILDPLATFORM} node:19.6.1-alpine as spec-build
RUN apk add --no-cache rsync git
RUN npm install -g markdown-include
WORKDIR /src
RUN --mount=target=/context \
    --mount=target=.,type=tmpfs <<EOT
  set -e
  rsync -a /context/. .
  markdown-include markdown-include.json
  mkdir /out
  cp spec.md /out
EOT

FROM scratch AS spec-update
COPY --from=spec-build /out /out

FROM spec-build AS spec-validate
RUN --mount=target=/context \
    --mount=target=.,type=tmpfs <<EOT
   set -e
   rsync -a /context/. .
   git add -A
   rm spec.md
   cp  /out/spec.md ./spec.md
   if [ -n "$(git status --porcelain -- ./spec.md)" ]; then
     echo >&2 'ERROR: Spec result differs. Please update with "make spec"'
     git status --porcelain -- ./spec.md
     exit 1
   fi
EOT
