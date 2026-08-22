---
title: Download
linkTitle: Download
description: Get the pig package manager, add the repositories, install extensions
weight: 45
icon: fa-solid fa-download
sidebar_enabled: false
breadcrumb: false
page_width: wide
---

Three things are downloadable here, in the order you need them: the `pig`
package manager, the two software repositories it configures, and the
extension packages themselves.

## 1 · The package manager {#pig}

`pig` is a single static Go binary with no runtime dependencies. Pick whichever
channel fits how you manage the machine.

{{< download pig >}}

## 2 · The repositories {#repositories}

`pig` writes these for you (`pig repo add all -u`), but they are ordinary APT
and DNF repositories and can be added by hand.

{{< cards cols=4 >}}
{{< card link="/repo/pgsql" title="PGSQL Repo" icon="database" subtitle="576 packaged extensions, one repository per distribution codename" />}}
{{< card link="/repo/infra" title="INFRA Repo" icon="cube" subtitle="OS-generic infrastructure packages, including pig itself" />}}
{{< card link="/repo/pgdg" title="PGDG Repo" icon="cloud-download" subtitle="The upstream PostgreSQL kernel repository these are built against" />}}
{{< card link="/repo/gpg" title="GPG Keys" icon="key" subtitle="Signing keys, fingerprints, and how to verify a package" />}}
{{< /cards >}}

Full setup instructions, including the mainland-China mirror, are on the
[Repository](/repo) page.

## 3 · The extension packages {#packages}

Every extension page carries its own download table: one row per platform and
PostgreSQL major, with the exact filename, size, and a direct link into the
repository.

{{< cards cols=4 >}}
{{< card link="/e" title="Extension list" icon="clipboard-list" subtitle="All 576 packaged extensions, with per-platform download tables" />}}
{{< card link="/os" title="By platform" icon="server" subtitle="What is available on each of the 16 Linux targets" />}}
{{< card link="/os/matrix" title="Availability matrix" icon="view-grid" subtitle="Every package against every OS and PostgreSQL major" />}}
{{< card link="/matrix/pgext-global-matrix.csv" title="Matrix as CSV" icon="document-download" subtitle="The same grid as raw data, alongside a JSON export" />}}
{{< /cards >}}

## Verifying what you downloaded {#verify}

Repository packages are GPG-signed; `apt` and `dnf` check the signature for you
once the [key](/repo/gpg) is installed. Release artifacts on GitHub are covered
by the `checksums.txt` published with each tag:

```bash
curl -fsSLO https://github.com/pgsty/pig/releases/latest/download/checksums.txt
sha256sum -c checksums.txt --ignore-missing
```

## Source {#source}

Nothing here is built from anything you cannot read:

| Repository | What it builds |
|:-----------|:---------------|
| [pgsty/pig](https://github.com/pgsty/pig) | The package manager itself |
| [pgsty/rpm](https://github.com/pgsty/rpm) | Every RPM in the PGSQL repository |
| [pgsty/deb](https://github.com/pgsty/deb) | Every DEB in the PGSQL repository |
| [pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | The OS-generic infrastructure packages |
