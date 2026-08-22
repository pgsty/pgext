---
title: About
linkTitle: About
description: What PGEXT.CLOUD is, who maintains it, and where the source lives
weight: 950
icon: fa-solid fa-circle-info
sidebar_enabled: false
breadcrumb: false
---

**PGEXT.CLOUD** is a catalog of the PostgreSQL extension ecosystem, and the
public face of the repository that builds a large part of it into native Linux
packages.

It exists because the ecosystem's greatest strength — anyone can extend the
database — is also what makes it hard to use. Extensions live on GitHub, on
PGXN, on vendor sites, in university FTP directories and in tarballs attached
to a decade-old mailing-list post. Finding out whether one exists, whether it
still builds, what it needs at runtime, and whether anyone has packaged it for
your distribution used to be an afternoon's work per extension.

## What it provides

{{< cards cols=3 >}}
{{< card link="/list" title="Catalog" icon="clipboard-list" subtitle="2,241 catalogued extensions, indexed by category, language and licence" />}}
{{< card link="/repo" title="Repository" icon="warehouse" subtitle="576 of them built as signed RPM and DEB across 16 Linux targets" />}}
{{< card link="/pig" title="Package manager" icon="cash" subtitle="One CLI that installs the kernel and its extensions by canonical name" />}}
{{< /cards >}}

The three layers are usable independently. The catalog is worth reading even if
you never install anything from the repository, and the repository works with a
plain `apt` or `dnf` if you would rather not add another CLI.

## Design commitments

- **PGDG-compliant.** The PostgreSQL kernel comes from the official
  [PGDG](https://www.postgresql.org/download/linux/) repositories, unmodified.
  These packages follow PGDG naming, paths and conventions so that the two
  compose rather than compete.
- **Linux-native.** Real RPM and DEB packages with real dependencies — not a
  private package format, not a container, not a build script.
- **Reproducible.** Every package is built from a public specification
  ([RPM](https://github.com/pgsty/rpm), [DEB](https://github.com/pgsty/deb))
  against a pinned source revision.
- **Both architectures.** `x86_64` and `aarch64` are built to parity; the
  [availability matrix](/os/matrix) shows exactly where they differ.
- **Free.** The catalog, the build specifications, the CLI and the repository
  are open source and free to use, with no registration and no quota.

## Maintainers

PGEXT.CLOUD is maintained by [**PGSTY**](https://github.com/pgsty) /
[**Vonng**](https://vonng.com/en/) (rh@vonng.com), and released under the
[**Apache License 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme).

| GitHub repository | What it holds |
|:------------------|:--------------|
| [github.com/pgsty](https://github.com/pgsty) | The **PGSTY** organization |
| [github.com/pgsty/pgext](https://github.com/pgsty/pgext) | This site, the extension metadata, and the tooling that generates both |
| [github.com/pgsty/pigsty](https://github.com/pgsty/pigsty) | The PostgreSQL distribution built on top of it |
| [github.com/pgsty/pig](https://github.com/pgsty/pig) | The `pig` package manager |
| [github.com/pgsty/rpm](https://github.com/pgsty/rpm) | RPM build specifications |
| [github.com/pgsty/deb](https://github.com/pgsty/deb) | DEB build specifications |
| [github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | Infrastructure package build specifications |

## Contributing

The most useful contribution is a correction. If an extension is missing, a
version is stale, a licence is wrong, or a package fails to install, open an
issue on [pgsty/pgext](https://github.com/pgsty/pgext/issues) — a large share
of the 576 packaged extensions arrived exactly that way.

Further reading: [***PostgreSQL is eating the Database World***](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4).
