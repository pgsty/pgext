## Usage

Sources:

- [Official ApsaraDB RDS extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [Official ApsaraDB RDS extension-management guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/manage-plug-ins)

`ganos_geometry_topology` is the Alibaba Cloud GanosBase topology component for spatial geometry analysis on supported ApsaraDB RDS for PostgreSQL instances. It is provider-managed software rather than a portable community package; version, availability, dependencies, and SQL surface follow the selected RDS engine.

### Enable and Verify

Use the RDS extension-management page for the target database and a privileged RDS account. Check the exact available version before creation:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'ganos_geometry_topology';

CREATE EXTENSION ganos_geometry_topology;

SELECT extname, extversion, extnamespace::regnamespace
FROM pg_extension
WHERE extname = 'ganos_geometry_topology';
```

After installation, use the version-matched GanosBase topology documentation and inspect the installed objects rather than assuming that names or semantics from another topology package are identical:

```sql
SELECT n.nspname, p.proname, pg_get_function_identity_arguments(p.oid)
FROM pg_extension AS e
JOIN pg_depend AS d
  ON d.refclassid = 'pg_extension'::regclass
 AND d.refobjid = e.oid
 AND d.deptype = 'e'
JOIN pg_proc AS p
  ON d.classid = 'pg_proc'::regclass
 AND d.objid = p.oid
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE e.extname = 'ganos_geometry_topology'
ORDER BY 1, 2, 3;
```

Use `pg_depend` and the RDS console to inventory other object classes; do not guess an API from the extension name.

### Provider Boundaries

- The RDS support matrix is authoritative for each PostgreSQL major and engine minor. A catalog version such as `7.0` does not prove that the same version can be installed on every instance.
- Alibaba requires a privileged database account for `ganos_geometry_topology`. Grant application roles only the specific topology schemas, tables, and functions they need.
- The provider warns that GanosBase and PostGIS extensions cannot be installed in the same schema. Plan schemas before enabling either family and do not relocate objects by hand.
- Topology structures typically have dependent schemas, tables, and geometry objects. Treat removal or upgrade as a data migration: inventory dependencies, back up, rehearse on a clone, and follow the RDS service procedure.
- The public extension-list page confirms availability but does not document this version's complete SQL API. Keep usage code tied to the installed version's GanosBase documentation and verify results, coordinate reference systems, validity, precision, and index plans with representative spatial data.
