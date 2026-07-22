## Usage

Sources:

- [AWS Aurora PostgreSQL extension support table](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [Amazon Aurora product page](https://aws.amazon.com/rds/aurora/)

`apgunit` was an AWS-internal Aurora PostgreSQL extension, not a portable community extension. AWS now describes it as no longer supported. Use this entry only to identify old Aurora clusters or dumps that mention it; do not design new database code around it.

### Historical Boundary

The AWS support table lists `apgunit` version `1.0` for old `Aurora PostgreSQL 9.6.3` and `Aurora PostgreSQL 9.6.6` releases, while later releases show it as unavailable. AWS does not publish a control file, SQL API, configuration reference, or migration interface for the component.

There is therefore no supported current enablement workflow to document. In particular, do not assume that `CREATE EXTENSION` is accepted on a modern Aurora cluster, that a similarly named package exists for self-managed PostgreSQL, or that `shared_preload_libraries` can restore it.

### Operational Guidance

When an inherited schema or migration references `apgunit`:

1. Identify the exact Aurora engine release on which it was used.
2. Ask AWS Support for the historical behavior and a supported replacement.
3. Remove or replace dependent objects before upgrading to an engine release where the extension is unavailable.
4. Test the resulting schema and application behavior on a separate target cluster.

Treat the catalog version as historical metadata, not evidence that `apgunit` can be installed or used today.
