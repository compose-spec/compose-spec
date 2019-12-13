FROM node:12 as generator
RUN npm install -g @adobe/jsonschema2md
COPY schema /schema
RUN jsonschema2md -d /schema/config_schema_v3.9.json -v "04" -o /docs

FROM scratch as docs
COPY --from=generator /docs / 
