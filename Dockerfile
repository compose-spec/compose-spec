FROM node:12 as generator
RUN npm install -g @adobe/jsonschema2md
COPY schema /schema
RUN jsonschema2md -d /schema/ -e json

FROM scratch as docs
COPY --from=generator /out /
