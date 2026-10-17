---
title: "PIG 1.10.0 Documentation"
description: The Missing Package Manager for PostgreSQL Extensions
icon: fa-solid fa-piggy-bank
page_width: full
cascade:
  type: docs
breadcrumb: false
comments: false
weight: 100
---

Current stable version: **v{{< param pig_version >}}**. Read the [release notes](/release/pig/#v1100).

PIG provides commands to find, install, and build PostgreSQL extension packages. It also manages PostgreSQL, Patroni, pgBackRest, and Pigsty workflows. Package installation uses the system's `apt`, `dnf`, or `yum`.

## Start here

| Your task | Page |
|:---|:---|
| Install PostgreSQL and enable your first extension | [Getting started](/pig/start/) |
| Learn what PIG does and which platforms it supports | [Introduction](/pig/intro/) |
| Download, install, update, or remove PIG | [Installation](/pig/install/) |
| Look up commands, options, and automation features | [Command reference](/pig/cmd/) |

## Install PostgreSQL and extensions

On a Linux host with root or `sudo` access, run:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
sudo pig repo add pgsql --update
sudo pig install -v 18 pg18 vector
```

In mainland China, use `https://repo.pigsty.cc/pig` for the installer. These commands install PostgreSQL 18 and pgvector packages. Continue with [Getting started](/pig/start/) to start a test instance and enable the extension with `CREATE EXTENSION vector`.

## Command reference

| Task | Command |
|:---|:---|
| Configure APT/DNF repositories | [`pig repo`](/pig/cmd/repo/) |
| Query the catalog, install, or update extension packages | [`pig ext`](/pig/cmd/ext/), [`pig install`](/pig/cmd/#pig-install) |
| Build extension packages from source | [`pig build`](/pig/cmd/build/) |
| Manage local PostgreSQL instances and databases | [`pig pg`](/pig/cmd/pg/) |
| Manage Patroni clusters through patronictl | [`pig pt`](/pig/cmd/pt/) |
| Inspect or reload pg_exporter | [`pig pe`](/pig/cmd/pe/) |
| Back up and restore with pgBackRest | [`pig pb`](/pig/cmd/pb/) |
| Run point-in-time recovery | [`pig pitr`](/pig/cmd/pitr/) |
| Install and configure Pigsty, manage Grafana | [`pig sty`](/pig/cmd/sty/) |
| Read, edit, and validate Pigsty inventory | [`pig inventory`](/pig/cmd/inventory/) |
| Run Pigsty administration playbooks | [`pig do`](/pig/cmd/do/) |
| Collect host, database, and backup status | [`pig context`](/pig/cmd/#pig-context) |

Run `pig help <command>` for command help. For automation, use `pig -o json --help` to inspect command capabilities, then select the relevant output and plan options.

## Find more information

Press {{< kbd "Ctrl" "K" >}} ({{< kbd "⌘" "K" >}} on macOS) to search the site. Press {{< kbd "/" >}} for page actions. An AI assistant action sends the public page URL and a reading prompt to the selected service.

- [Design Records](https://pig.pgsty.com/design/): implementation decisions, trade-offs, and proposals.
- [Release notes](/release/pig/): changes and download files for each version.
- [Extension catalog](https://pigsty.io/ext/list/): extension descriptions and package information.
- [GitHub Issues](https://github.com/pgsty/pig/issues): report a problem or request a feature.

PIG is developed by [Vonng](https://vonng.com/en/) and released under [Apache 2.0](https://github.com/pgsty/pig/blob/main/LICENSE). [Pigsty](https://pigsty.io/) provides PostgreSQL deployment, monitoring, backups, and high availability.
