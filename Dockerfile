# syntax=docker/dockerfile:1

FROM node:16.1 as generator
RUN npm install -g @adobe/jsonschema2md
COPY schema /schema
RUN jsonschema2md -d /schema/ -e json

FROM scratch as docs
COPY --from=generator /out /

FROM --platform=${BUILDPLATFORM} alpine as spec-build
WORKDIR /src
COPY *.md /src
RUN <<EOT
  set -e
  cat head.md > spec.md
  cat 01-status.md >> spec.md
  cat 02-model.md >> spec.m >> spec.md
  cat 03-compose-file.md >> spec.md
  cat 04-version-and-name.md >> spec.md
  cat 05-services.md >> spec.md
  cat 06-networks.md >> spec.md
  cat 07-volumes.md >> spec.md
  cat 08-configs.md >> spec.md
  cat 09-secrets.md >> spec.md
  cat 10-fragments.md >> spec.md
  cat 11-extension.md >> spec.md
  cat 12-interpolation.md >> spec.md
  cat 13-merge.md >> spec.md
  cat 14-include.md >> spec.md
  cat 15-profiles.md >> spec.md

  mkdir /out
  cp spec.md /out
EOT

FROM scratch AS spec-update
COPY --from=spec-build /out /out

FROM --platform=${BUILDPLATFORM} alpine as spec-validate
RUN apk add --no-cache rsync git
WORKDIR /src
COPY --from=spec-build /out /out
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
