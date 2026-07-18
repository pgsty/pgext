## Usage

Sources:

- [ApsaraDB RDS for PostgreSQL supported-extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_address_standardizer` is Alibaba Cloud's GanosBase address-standardization extension based on the PAGC rules. It is provider-only, and its available version depends on the ApsaraDB RDS PostgreSQL major, edition, and engine minor. The current Standard Edition matrix lists 7.0 for PostgreSQL 17, 6.9 for PostgreSQL 12 through 16, 6.3 for PostgreSQL 10 and 11, and no support for PostgreSQL 18.

### Verify the instance before installation

The service catalog is the authority for a particular instance. Check both the base extension and the optional US data package before creating either one:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name IN (
  'ganos_address_standardizer',
  'ganos_address_standardizer_data_us'
)
ORDER BY name;

CREATE EXTENSION ganos_address_standardizer;
CREATE EXTENSION ganos_address_standardizer_data_us;
```

Install the US data extension only when its rules and reference data match the address population. Inventory the installed functions and signatures on the instance rather than assuming that community PostGIS examples exactly match the Ganos build.

### Provider and data-quality boundaries

Version 7.0 in the catalog is not a universal version for all PostgreSQL majors. Service security restrictions can also prevent new extension creation on particular engine minors while leaving existing installations intact. Confirm availability again before restore, replica creation, major upgrade, or region migration.

Address standardization is parsing and normalization, not proof that an address exists or belongs to a person. Results depend on rules, reference data, language, abbreviations, and locale. Retain the original input, record the extension and data-package versions, test representative international and malformed values, and route ambiguous results for review. Backups and migrations remain tied to Alibaba Cloud's extension lifecycle and supported engine combinations.
