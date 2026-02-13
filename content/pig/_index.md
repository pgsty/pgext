---
title: PIG CLI
description: The Missing Package Manager for PostgreSQL Extensions
icon: PiggyBank
full: true
cascade:
  type: docs
breadcrumbs: false
comments: false
weight: 100
---

— ***Postgres Install Genius, the missing extension package manager for PostgreSQL ecosystem***

{{< cards >}}
{{< card link="/pig/intro"   title="Introduction" subtitle="Why we need a package manager" icon="sparkles" >}}
{{< card link="/pig/start"   title="Get Started"  subtitle="Tutorial and examples"         icon="play" >}}
{{< card link="/pig/install" title="Installation" subtitle="Install in different ways" icon="save" >}}
{{< /cards >}}



## Quick Start

[Install](/pig/install) `pig` with a single command

{{< tabs items="Default,Mirror" >}}
{{< tab >}}
```bash tab="Default"
curl -fsSL https://repo.pigsty.io/pig | bash     # via Cloudflare
```
{{< /tab >}}
{{< tab >}}
```bash tab="Mirror"
curl -fsSL https://repo.pigsty.cc/pig | bash     # via China Mirror
```
{{< /tab >}}
{{< /tabs >}}


Then it's ready to use, assume you want to install the [**`pg_duckdb`**](/e/pg_duckdb/) extension:

```bash
$ pig repo set                  # overwrite Linux, Pigsty + PGDG repo in one-pass!
$ pig install pg18              # install PostgreSQL 18 kernel package from official PGDG repo
$ pig install pg_duckdb -v 18   # install pg_duckdb extension for PG 18 from pigsty repo
$ pig install -y postgis timescaledb  # install some extension for active postgres installation
$ pig install -y vector         # use both pkg and ext name to install extensions
```


## CLI Usage

Check sub-commands [documentation](/pig/cmd) with `pig help <command>`

{{< cards cols="4" >}}
{{< card link="/pig/cmd/repo"  title="pig repo"  subtitle="Manage software repositories" icon="library" >}}
{{< card link="/pig/cmd/ext"   title="pig ext"   subtitle="Manage postgres extensions"   icon="cube" >}}
{{< card link="/pig/cmd/build" title="pig build" subtitle="Build extension from source"  icon="view-grid" >}}
{{< card link="/pig/cmd/sty"   title="pig sty"   subtitle="Manage pigsty installation"   icon="cloud-download" >}}
{{< /cards >}}

{{< cards cols="4" >}}
{{< card link="/pig/cmd/pg"    title="pig pg"    subtitle="Manage local PostgreSQL"      icon="database" >}}
{{< card link="/pig/cmd/pt"    title="pig pt"    subtitle="Manage Patroni HA cluster"    icon="refresh" >}}
{{< card link="/pig/cmd/pb"    title="pig pb"    subtitle="Manage pgBackRest backup"     icon="archive" >}}
{{< card link="/pig/cmd/pitr"  title="pig pitr"  subtitle="Orchestrated PITR recovery"   icon="clock" >}}
{{< /cards >}}



## About

The `pig` CLI is developed by [Vonng](https://blog.vonng.com/en/), and open-sourced under the [Apache License 2.0](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme).

You can also check the [pigsty](https://pgsty.com) project, which makes it even smoother to deliver all these [extensions](/list) in an IaC way:

{{< cards cols=4 >}}
{{< card link="https://github.com/pgsty/pgext"  title="PGEXT"  icon="github" subtitle="This Website and Extension Catalog" >}}
{{< card link="https://github.com/pgsty/pig"    title="PIG"    icon="github" subtitle="The PostgreSQL Package Manager" >}}
{{< card link="https://github.com/pgsty/pigsty" title="PIGSTY" icon="github" subtitle="The Battery-Included PostgreSQL Distribution" >}}
{{< /cards >}}

