## Usage

Sources:

- [Alibaba Cloud extension support matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_address_standardizer_data_us` version `7.0` is the Alibaba Cloud ApsaraDB RDS for PostgreSQL US address-standardization data plug-in based on the PAGC standard. It is a provider-supplied data component, not a catalog of independently documented SQL functions.

### Enablement

First verify that the extension and version are listed for the exact ApsaraDB RDS PostgreSQL engine version and edition. Where the instance permits SQL installation, the provider documents the standard extension mechanism:

```sql
CREATE EXTENSION ganos_address_standardizer_data_us;
```

Some managed extensions must instead be enabled through the service console. Follow the instance's provider workflow and confirm installation in `pg_extension`.

### Operational Boundary

The provider matrix lists `ganos_address_standardizer` separately from `ganos_address_standardizer_data_us`. Install and use the separately supported standardizer component for address-processing operations; the data plug-in itself has no standalone function surface documented by the provider page.

Availability changes across engine releases and editions. If the extension is absent, Alibaba Cloud advises upgrading the minor engine version or checking the applicable support column. Do not apply this workflow to self-managed PostgreSQL: the cited artifact and enablement contract are specific to Alibaba Cloud.
