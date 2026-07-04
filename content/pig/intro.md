---
title: Introduction
description: Why would we need yet another package manager? especially for Postgres extensions?
weight: 100
icon: CircleHelp
breadcrumbs: false
---

Have you ever struggled with installing or upgrading PostgreSQL extensions? Digging through outdated documentation, cryptic configuration scripts, or searching GitHub for forks and patches?
Postgres's rich extension ecosystem also means complex deployment processes, especially across multiple distributions and architectures. PIG can solve these headaches for you.

This is exactly why **Pig** was created. Developed in Go, Pig is dedicated to one-stop management of Postgres and its [531](/list) extensions.
Whether it's TimescaleDB, Citus, PGVector, 30+ Rust extensions, or all the components needed to self-host Supabase, Pig's unified CLI makes everything accessible.
It completely eliminates source compilation and messy repositories, directly providing version-aligned RPM/DEB packages that perfectly support Debian, Ubuntu, RedHat, and other mainstream distributions on both x86 and Arm architectures, no guessing, no hassle.

Pig isn't reinventing the wheel; it fully leverages native system package managers (APT, YUM, DNF) and strictly follows [PGDG official](/repo/pgdg/) packaging standards for seamless integration.
You do not need to choose between "the standard way" and "shortcuts". Pig respects existing repositories, follows OS best practices, and coexists harmoniously with existing repositories and packages.
If your Linux system and PostgreSQL major version are not in the [supported list](#linux-compatibility), you can use [`pig build`](/pig/cmd/build/) to compile extensions for your specific combination.

Want to supercharge your Postgres and escape the hassle? Visit the [PIG official documentation](/pig/) for guides, and check out the extensive [extension list](/list),
turning your local Postgres database into an all-capable multi-modal data platform with one click.
If [Postgres's future is unmatched extensibility](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4), then Pig is the magic lamp that helps you unlock it. After all, no one ever complains about "too many extensions".

## Automation-Friendly

PIG's command system is automation-ready out of the box: consistent argument conventions, stable output behavior, `--plan` previews, and confirmation flows for high-risk operations to reduce mistakes.


> [ANNOUNCE pig: The Postgres Extension Wizard](https://www.postgresql.org/about/news/announce-pig-the-postgres-extension-wizard-2988/)


--------

## Linux Compatibility

PIG and the Pigsty extension repository support the following Linux distribution and PostgreSQL version combinations:

| OS Code        | Vendor | Major |  Minor  | Full Name         | PG Versions |  Notes   |
|:---------------|:-------|:-----:|:-------:|:------------------|:------------|:--------:|
| `el7.x86_64`   | EL     |   7   |   7.9   | CentOS 7 x86      | 13-15       |   EOL    |
| `el8.x86_64`   | EL     |   8   |  8.10   | RockyLinux 8 x86  | 14-18       | Near EOL |
| `el8.aarch64`  | EL     |   8   |  8.10   | RockyLinux 8 ARM  | 14-18       | Near EOL |
| `el9.x86_64`   | EL     |   9   |   9.7   | RockyLinux 9 x86  | 14-18       |    ✅     |
| `el9.aarch64`  | EL     |   9   |   9.7   | RockyLinux 9 ARM  | 14-18       |    ✅     |
| `el10.x86_64`  | EL     |  10   |  10.1   | RockyLinux 10 x86 | 14-18       |    ✅     |
| `el10.aarch64` | EL     |  10   |  10.1   | RockyLinux 10 ARM | 14-18       |    ✅     |
| `d11.x86_64`   | Debian |  11   |  11.11  | Debian 11 x86     | 14-18       |   EOL    |
| `d11.aarch64`  | Debian |  11   |  11.11  | Debian 11 ARM     | 14-18       |   EOL    |
| `d12.x86_64`   | Debian |  12   |  12.13  | Debian 12 x86     | 14-18       |    ✅     |
| `d12.aarch64`  | Debian |  12   |  12.13  | Debian 12 ARM     | 14-18       |    ✅     |
| `d13.x86_64`   | Debian |  13   |  13.3   | Debian 13 x86     | 14-18       |    ✅     |
| `d13.aarch64`  | Debian |  13   |  13.3   | Debian 13 ARM     | 14-18       |    ✅     |
| `u22.x86_64`   | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 x86  | 14-18       |    ✅     |
| `u22.aarch64`  | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 ARM  | 14-18       |    ✅     |
| `u24.x86_64`   | Ubuntu |  24   | 24.04.4 | Ubuntu 24.04 x86  | 14-18       |    ✅     |
| `u24.aarch64`  | Ubuntu |  24   | 24.04.4 | Ubuntu 24.04 ARM  | 14-18       |    ✅     |
| `u26.x86_64`   | Ubuntu |  26   | 26.04.0 | Ubuntu 26.04 x86  | 14-18       |    ✅     |
| `u26.aarch64`  | Ubuntu |  26   | 26.04.0 | Ubuntu 26.04 ARM  | 14-18       |    ✅     |
{.full-width}

**Notes:**

- **EL** refers to RHEL-compatible distributions, including RHEL, CentOS, RockyLinux, AlmaLinux, OracleLinux, etc.
- **EOL** indicates the operating system has reached or is about to reach end of support; upgrading to a newer version is recommended
- **✅** indicates full support; recommended for use
- PG versions 14-18 means support for PostgreSQL 14, 15, 16, 17, and 18 major versions
