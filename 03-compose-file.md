## Compose file

The Compose file is a [YAML](http://yaml.org/) file defining:
- [Version](04-version-and-name.md) (Optional)
- [Services](05-services.md) (Required)
- [Networks](06-networks.md)
- [Volumes](07-volumes.md)
- [Configs](08-configs.md) 
- [Secrets](09-secrets.md)

The default path for a Compose file is `compose.yaml` that is placed in the working directory.
Compose also supports `docker-compose.yaml` and `docker-compose.yml` for backwards compatibility of earlier versions.
If both files exist, Compose prefers the canonical `compose.yaml`.

You can use [fragments](10-fragments.md) and [extensions](11-extension.md) to keep your Compose file efficient and easy to maintain.

Multiple Compose files can be [merged](13-merge.md) together to define the application model. The combination of YAML files must be implemented by appending/overriding YAML elements based on the Compose file order set by the user. 
Simple attributes and maps get overridden by the highest order Compose file, lists get merged by appending. Relative
paths must be resolved based on the first Compose file's parent folder, whenever complimentary files being
merged are hosted in other folders. As some Compose file elements can both be expressed as single strings or complex objects, merges must apply to
the expanded form.

You can also use [`include`](14-include.md) in your Compose file if you want your Compose application to be dependent on another Compose application, managed by a different team or shared with others.
