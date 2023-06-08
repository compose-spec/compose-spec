## Dependency on other compose projects

A Compose application can declare dependency on another compose application, managed by another team
or shared with others. This allows to keep a compose file reasonably complicated for the limited
amount of resources a single team has to declare for it's own sub-domain within a larger deployment.

`include` top-level section is used to define dependency to another compose application (or subdomain).
Each path listed in this section must be loaded as an individual compose model with it's own project
directory to resolve relative paths. 

Once the application to be included has been loaded, all resources definitions are copied into the 
current compose application model. Compose must warn user if some resource name conflict, and not 
try to merge those. To enforce this, `include` must be evaluated after the compose file(s) selected 
by user to define the Compose application model have been parsed and merged, so that conflicts 
between included compose files and those selected are detected.

`include` applies recursively: an included compose file which declares it's own `include` section
will trigger those other files to be included as well.

The resulting resources can be used in the including compose model for cross-services references.

### short syntax

Short syntax only defines path to another compose file. File is loaded with parent
folder as project directory, and optional `.env` file being loaded to define variables default values
for interpolation, while local project environment can override those values. 

```yaml
include:
  - ../commons/compose.yaml
  - ../another_domain/compose.yaml

services:
  webapp:
    depends_on:
      - included-service # defined by another_domain
```

In this illustration example, when loading compose file, both `../commons/compose.yaml` and 
`../another_domain/compose.yaml` are loaded as individual compose projects. Relative paths 
in compose files being refered by `include` are resolved relative to their own compose 
file path, not based on local project directory. Variables are interpolated using values set in
`.env` optional file in same folder, and can be overriden by local project environment.

### long syntax

Long syntax offer fine-grain control over the sub-project parsing:

```yaml
include:
   - path: ../commons/compose.yaml
     project_directory: ..
     env_file: ../another/.env
```

#### path
`path` is required and defines the location of the compose file(s) to be parsed and included into
local compose model. `path` can be set either to a string when a single compose file is involved,
or to a list of strings when multiple compose files need to be [merged together](14-merge.md) to 
define the compose model to be included in local application.

```yaml
include:
   - path: 
       - ../commons/compose.yaml
       - ./commons-override.yaml
```

#### project_directory
`project_directory` defines base path to resolve relative paths set in compose file. It defaults to 
the directory of the included compose file.

#### env_file
`env_file` defines an environment file(s) to use to define default values when interpolating variables
in the compose file being parsed. It defaults to `.env` file in the `project_directory` for the compose 
file being parsed. 

`env_file` can be set either to a string or a list of strings when multiple env_file need to be merged
to define project environment.

```yaml
include:
   - path: ../another/compose.yaml
     env_file:
       - ../another/.env
       - ../another/dev.env
```

Local project environment have precendence over values set in this file, so that local project can
override values for customization.
