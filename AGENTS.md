# Compose Specification — agent instructions

This repository is the Compose Specification: the normative markdown
chapters (`0*-*.md`, `build.md`, `deploy.md`, `develop.md`, `models.md`) plus
`spec.md`, which is **generated** from them — never edit `spec.md` directly,
regenerate it with `make spec` after changing any chapter.

## Version badge rule

Every attribute documented in the spec carries a badge stating the Docker
Compose release that introduced it, placed on its own line right after the
attribute heading:

```markdown
## some_attribute

[![Compose v2.30.0](https://img.shields.io/badge/compose-v2.30.0-blue?style=flat-square)](https://github.com/docker/compose/releases/v2.30.0)
```

**When documenting a new attribute, always add this badge.** The release
that will ship the attribute is usually unknown at the time the spec change
is written: use the `NEXT` placeholder, to be replaced with the actual
version once the next Docker Compose release is cut:

```markdown
[![Compose NEXT](https://img.shields.io/badge/compose-NEXT-blue?style=flat-square)](https://github.com/docker/compose/releases)
```

`NEXT` is greppable on purpose: at release time, maintainers sweep the repo
for `compose-NEXT` and pin the badges to the released version (including the
release link). Only the long-standing attributes inherited from the original
2021 schema import go without a badge — anything added since then must have
one.

## Checks before submitting

- `make spec` — regenerate `spec.md` (CI fails if it drifts).
- Every commit needs a DCO `Signed-off-by` trailer (`git commit -s`).
