# Contributing

Contributions should be made via pull requests. Pull requests will be reviewed
by two or more maintainers and merged when acceptable.

The goal of the Compose Specification is to be the simplest cloud and platform
agnostic way of defining container based applications. A developer should be
able to develop and debug an application on their local system before
confidently deploying it to a production platform– cloud or otherwise. The
format will be portable across container platforms by providing a high level
abstraction for how the containers are built, deployed, networked, and
configured.

When proposing features as part of the Compose Specification, changes should be
focused towards features which enable developers as part of their
[inner loop](https://docs.microsoft.com/en-us/dotnet/architecture/containerized-lifecycle/design-develop-containerized-apps/docker-apps-inner-loop-workflow)
and not focused on operator controls.
Some features may benefit the container ecosystem, however, they may not be
appropriate for a first class feature in Compose.

## Successful Changes

We ask that contributors read the [Compose Vision](VISION.md) to ensure that
proposed changes are aligned with the objectives of the Compose project.

To help maintainers understand what user or developer problem needs to be
solved, the first step to a contribution is usually submitting an issue. A well
written issue is one that clearly outlines the developer or user problem that
needs to be solved along with a list of requirements for resolution of the
problem. If there are multiple possible solutions to the problem, these can be
outlined in the issue. Once consensus is reached on how to resolve the issue, a
pull request can be created.

Pull requests that propose minor changes or improvements may be submitted
without an associated issue or discussion.

For large or high impact changes, contributors can reach out to maintainers
before starting work. This will ensure that contributors and maintainers are
aligned and increase the chance that the change is accepted.

## Commit Messages

Commit messages should follow best practices and explain the context of the
problem and how it was solved-- including any caveats or follow up changes
required. They should tell the story of the change and provide readers an
understanding of what led to it.

[How to Write a Git Commit Message](https://cbea.ms/git-commit/)
provides a good guide for how to do so.

In practice, the best approach to maintaining a nice commit message is to
leverage a `git add -p` and `git commit --amend` to formulate a solid
change set. This allows one to piece together a change, as information becomes
available.

If you squash a series of commits, don't just submit that. Re-write the commit
message, as if the series of commits was a single stroke of brilliance.

That said, there is no requirement to have a single commit for a pull request,
as long as each commit tells the story. For example, if there is a feature that
requires a package, it might make sense to have the package in a separate commit
then have a subsequent commit that uses it.

Remember, you're telling part of the story with the commit message. Don't make
your chapter weird.

## Sign your work

The sign-off is a simple line at the end of the explanation for the patch. Your
signature certifies that you wrote the patch or otherwise have the right to pass
it on as an open-source patch. The rules are pretty simple: if you can certify
the below (from [developercertificate.org](https://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe.smith@email.com>

Use your real name (sorry, no pseudonyms or anonymous contributions.)

If you set your `user.name` and `user.email` git configs, you can sign your
commit automatically with `git commit -s`.


