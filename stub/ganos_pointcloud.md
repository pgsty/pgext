## Usage

Sources:

- [Alibaba Cloud extension support matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [Alibaba Cloud point cloud model documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/point-cloud-model)

`ganos_pointcloud` is an Alibaba Cloud GanosBase extension for storing, compressing, inspecting, and processing point-cloud data in ApsaraDB RDS for PostgreSQL. It is a provider-specific component, not a portable community package; availability and the exact version depend on the RDS PostgreSQL engine generation and edition.

### Enable the Extensions

On a supported RDS instance, create the core extension in the intended schema. Add the geometry bridge when spatial conversion and analysis are needed:

```sql
CREATE EXTENSION ganos_pointcloud WITH SCHEMA public CASCADE;
CREATE EXTENSION ganos_pointcloud_geometry WITH SCHEMA public CASCADE;
```

`CASCADE` may install prerequisites available on the service. Confirm the resulting extension list and versions after creation, because the support matrix contains different Ganos releases for different engine versions.

### Data Model

Point formats are registered in `pointcloud_formats`. Each format defines dimensions, data types, scale/offset rules, compression, and an SRID through XML. Point values use `pcpoint`, while collections use `pcpatch`; all values reference a format identifier so the binary payload can be interpreted correctly. `ganos_pointcloud_geometry` connects these values to Ganos geometry operations.

Register and validate the format before loading a data set, then keep its dimension order, scale, offset, and SRID stable. Mixing values with an incorrect format identifier can silently reinterpret coordinates or attributes.

### Operational Notes

Check the Alibaba Cloud matrix for the target engine minor version before deployment or upgrade. Point-cloud sets can be large, so test compression choice, ingest throughput, indexes, spatial conversion, backup size, and query memory on representative data. Extension creation and upgrades follow Alibaba Cloud's privilege and maintenance model; do not assume that a self-managed build or APIs from the community `pointcloud` project are interchangeable with the installed Ganos release.
