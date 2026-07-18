## Usage

Sources:

- [Alibaba Cloud RDS for PostgreSQL extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_geometry_sfcgal` is cataloged as a Ganos 7.0 component offered by Alibaba Cloud RDS for PostgreSQL. It is provider-managed software rather than a portable extension with a public source tree, control file, or SQL API captured by this review.

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'ganos_geometry_sfcgal';
```

Use the query only to discover what the current RDS engine exposes. Availability can vary by RDS engine version, edition, region, and maintenance policy. Enable the component only through the provider-supported workflow shown for the exact instance; do not copy installation commands or function names from unrelated SFCGAL projects.

The public matrix confirms availability but does not document a standalone function surface in the reviewed material. Inspect `pg_available_extension_versions`, provider documentation exposed in the console, and a disposable RDS instance before depending on it. Plan migrations around the fact that this component may not exist outside Alibaba Cloud.
