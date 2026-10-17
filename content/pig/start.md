---
title: Get Started
description: Quick start guide for PIG, the PostgreSQL package manager
icon: fa-solid fa-play
weight: 200
breadcrumb: false
---

This tutorial installs PostgreSQL 18 and pgvector, starts a local test instance, and enables the extension in a database.

Use a [supported Linux test host](/pig/intro/#linux-compatibility) with internet access and `curl` installed. You need root or `sudo` access. Run the commands in order.

## Install PIG {#installation}

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
pig version
```

In mainland China, use `https://repo.pigsty.cc/pig` instead. The installer selects the RPM or DEB package for your host. See [Installation](/pig/install/) for other methods.

## Check the host {#check-environment}

```bash
pig status
```

Check the OS code, architecture, and detected PostgreSQL installations. This command also tests repository connectivity.

## Add repositories {#add-repositories}

Add PGDG and Pigsty PostgreSQL repositories, then refresh package metadata:

```bash
sudo pig repo add pgsql --update
```

For China mirrors, add `--mirror` to this command.

## Install PostgreSQL and pgvector {#install-pg}

```bash
sudo pig install -v 18 pg18 vector
```

Review the package manager's transaction and confirm it.

`pg18` selects the PostgreSQL 18 packages. `vector` selects pgvector. The `-v 18` option selects the extension's PostgreSQL major version.

Check the installed extension files:

```bash
pig ext status -v 18
```

## Start a test instance {#start-a-test-instance}

This tutorial uses `/pg/data-demo` and port `15432` to avoid the default cluster on port `5432`. Use an unused directory and port.

Create an empty data directory owned by the `postgres` operating-system user:

```bash
sudo install -d -m 700 -o postgres -g postgres /pg/data-demo
sudo pig pg init -v 18 -D /pg/data-demo
```

Start the instance on port `15432`:

```bash
sudo pig pg start -v 18 -D /pg/data-demo -O '-p 15432'
sudo pig pg status -v 18 -D /pg/data-demo
```

## Enable and test pgvector {#enable-and-test-pgvector}

Connect to the test instance:

```bash
sudo pig pg psql -v 18 -D /pg/data-demo
```

PIG reads the instance's connection information from the specified data directory. At the `psql` prompt, run:

```sql
CREATE EXTENSION vector;
SELECT extname, extversion FROM pg_extension WHERE extname = 'vector';
SELECT '[1,2,3]'::vector <-> '[4,5,6]'::vector AS distance;
```

The first query returns the installed pgvector version. The second returns a distance of approximately `5.196`. Exit `psql` with `\q`.

`CREATE EXTENSION` enables an extension in the connected database. Repeat it in each database that needs pgvector.

Stop the test instance when you finish:

```bash
sudo pig pg stop -v 18 -D /pg/data-demo
```

The data stays in `/pg/data-demo`. Run the start command again to reuse the instance. See [PostgreSQL management](/pig/cmd/pg/) for service configuration and other operations.

## Find more extensions {#list-extensions}

```bash
pig ext list                 # List the catalog
pig ext list duckdb          # Search by name or description
pig ext info pg_duckdb       # Show metadata and setup requirements
pig ext avail pg_duckdb      # Show package availability
```

Before installing another extension, use `info` to check its setup requirements and `avail` to check package availability.

## Select packages and versions {#alias-translation}

To use another PostgreSQL major version, replace `18` in the commands above.
See [extension installation](/pig/cmd/ext/#ext-add) for package groups, version selection, and existing installations.

## Inspect and update {#show-extensions}

`pig ext status -v 18` lists installed extension files. To list extensions enabled in a database, run `\dx` in `psql`.
See [extension updates](/pig/cmd/ext/#ext-update) when you need to upgrade a package and its database objects.
