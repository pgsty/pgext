# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

pgext is a PostgreSQL Extension Metadata Manager that manages PostgreSQL extension metadata using PostgreSQL itself. It fetches, parses, and maintains metadata about PostgreSQL extensions from various package repositories (APT, YUM, PGXN).

## Architecture

The project uses a PostgreSQL database schema (`pgext`) to store and manage extension metadata. Key components:

- **Database Schema**: Located in `db/schema.sql` and `cli/meta/assets/pgext-schema.sql`. Contains tables for extensions, repositories, packages, and availability matrices.
- **CLI Tool**: Cobra-based CLI in `cmd/` that provides commands for managing the metadata lifecycle
- **Meta Package**: Core logic in `cli/meta/` handles database operations, fetching, parsing, and processing metadata

## Common Development Commands

```bash
# Build the binary
go build -o pgext

# Run the tool
./pgext --help

# Database operations (requires PostgreSQL with semver extension)
./pgext init                    # Initialize pgext schema
./pgext load ext                # Load extension catalog from CSV
./pgext fetch                   # Fetch repository metadata
./pgext parse                   # Parse repository data into packages
./pgext recap                   # Generate availability matrix
./pgext reload                  # Complete reload cycle (fetch + parse + recap)
./pgext status                  # Show metadata status

# Development with custom database
./pgext -d "postgres:///vonng" init
./pgext -d "host=localhost dbname=mydb" load extension

# Makefile shortcuts (Hugo/documentation related - not for pgext itself)
make dump                       # Export data from PostgreSQL to CSV files
make load                       # Import CSV files to PostgreSQL
```

## Database Connection

The tool accepts database connections via:
- `-d` flag: Can be a database name, PostgreSQL URL, or key=value connection string
- `PGURL` environment variable
- pg_service.conf entries

The test environment will use the 'postgres:///' URL on a local PostgreSQL, you can use that database for testing.


## Key Workflows

1. **Initial Setup**: `init` → `load extension/repository/category` → `fetch` → `parse` → `recap`
2. **Update Cycle**: `fetch` (with --force to re-download) → `parse` → `recap`
3. **Quick Reload**: `reload` command combines fetch + parse + recap

## Data Flow

1. **Extensions**: CSV catalog loaded via `load` command
2. **Repositories**: Metadata fetched from APT/YUM repos via `fetch`
3. **Packages**: Repository data parsed into package records via `parse`
4. **Availability**: Matrix generated from packages via `recap`

## Important Tables

The schema is defined in `db/schema.sql`

- `pgext.status`: Singleton status tracking table, fixed one record with id = 0
- `pgext.pg` : config table, one row per PostgreSQL major version, only `active` pg version need to be considered (`pgext.active_pg`)
- `pgext.os` : config table, one row per OS/distribution, only `active` os need to be considered (`pgext.active_os`)
- `pgext.extension`: Core extension catalog, details info about postgres extension
- `pgext.repository`: extension repository definitions, reference os
- `pgext.repo_data`: raw metadata fetched from repositories, wait to be parsed into pgext.apt and pgext.dnf
- `pgext.apt`: parsed apt package metadata
- `pgext.dnf`: parsed dnf/yum package metadata
- `pgext.bin`: filtered, processed, sorted and merged apt/dnf packages of postgres extensions, wait to be recap (id generated, order mattered) 
- `pgext.pkg`: the final yield, availability matrix of postgres extensions

One `pkg` (extension package) may contain multiple `ext` (extension), but there's a leading (primary) extension in a package.
One `pkg` can be versioned for multiple `pg` (PostgreSQL major versions) , like the `pgvector` can be versioned for `pgvector_17`, `postgresql-16-pgvector`.



## Parallel Processing

Commands support concurrent workers via `-p/--parallel` flag (default: 8) for operations like fetching and parsing.


