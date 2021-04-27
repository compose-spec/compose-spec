# Release Process

## Goals

* Do not explicitly expose Compose specification versions to users
* Keep the process simple and light

## Process

### Requesting a release

To trigger the release process, the requester MUST open a pull request against
this repository with an updated [CHANGELOG.md](./CHANGELOG.md).
Anyone can request a release of the specification.

A section MUST be added to the file with the following format:

```markdown
## date (commit digest)

* item 1
* item n
```

The date:
* MUST be of the form `YYYY-MM-DD`
* MUST be the same as that of the commit referenced by the digest

The commit digest:
* MUST be of the same form as that given by the command `git rev-parse HEAD`
* MUST be a valid commit digest on the default branch (currently `master`)
* MUST be at the point that the release requester would like the release to be
  made

The list of changes:
* MUST be a bullet point list
* MUST list all impactful changes to the specification since the last release

### Approving a release

To approve a release, a simple majority of maintainers MUST approve the pull
request.
The maintainers are expected to check that the pull request requirements are
met.
Once a quorum has been reached, the pull request MUST be merged.

### Publishing a release

Once the pull request has been merged, a maintainer MUST create a tag at the
commit digest mentioned in the changelog and push it to this repository.
This tag represents a release of the Compose specification.

The version of the specification is then the date or commit digest of the
release.
Tool builders MAY reference this version.
