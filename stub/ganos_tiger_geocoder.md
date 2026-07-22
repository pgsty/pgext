## Usage

Sources:

- [Alibaba Cloud supported-extension matrix](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)
- [Alibaba Cloud extension management](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/manage-plug-ins)
- [Official Tiger Geocoder documentation](https://postgis.net/docs/Extras.html#Tiger_Geocoder)
- [Official geocode function reference](https://postgis.net/docs/Geocode.html)

`ganos_tiger_geocoder` is an Alibaba Cloud GanosBase extension for U.S. Census Bureau TIGER data, forward geocoding, and reverse geocoding. It is a managed-service component, not an extension that can be installed on community PostgreSQL from an upstream source repository.

### Core Workflow

On ApsaraDB RDS for PostgreSQL, confirm that the instance engine version exposes 7.0, make the target database owned by a privileged account, and install `ganos_tiger_geocoder` through the Alibaba Cloud extension-management workflow. GanosBase and PostGIS extensions cannot be installed in the same schema.

The geocoder needs TIGER lookup and state data before address matching is useful. After the provider-supported loader has populated the data and the `tiger` schema is on `search_path`, a forward-geocoding query follows the Tiger interface:

```sql
SET search_path = public, tiger;

SELECT pprint_addy(addy) AS normalized_address,
       rating,
       ST_AsText(geomout) AS point
FROM geocode('1 Devonshire Place, Boston, MA 02109', 5);
```

Important user-facing objects include `geocode` for ranked address candidates, `reverse_geocode` for addresses near a point, `normalize_address` and `pprint_addy` for address handling, and loader or index-maintenance helpers in the `tiger` schema. Lower `rating` values indicate better forward-geocoding matches.

### Managed-Service Boundaries

Availability and version depend on the Alibaba Cloud product, region, PostgreSQL major, and minor engine release; query `pg_available_extensions` on the actual instance instead of assuming catalog availability. The service documentation requires a privileged database owner for this extension. Loading national or state TIGER data consumes substantial storage and maintenance time, and geocoding quality depends on the loaded Census vintage and local data quality.

The linked PostGIS documentation describes the Tiger SQL model to which Alibaba Cloud points, but Alibaba controls the Ganos build and lifecycle. Follow the provider console and support documentation for installation, data loading, upgrades, backups, and removal; do not substitute community PostGIS files or expect `ganos_tiger_geocoder` to work on self-managed PostgreSQL.
