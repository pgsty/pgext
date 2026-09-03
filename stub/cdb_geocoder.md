## Usage

Sources:

- [Official documentation](https://github.com/CartoDB/data-services/blob/11ecbff14d81b71af4cc5f87c503bd71084fcefc/geocoder/extension/README.md)
- [Extension control file](https://github.com/CartoDB/data-services/blob/11ecbff14d81b71af4cc5f87c503bd71084fcefc/geocoder/extension/cdb_geocoder.control)
- [Official repository](https://github.com/CartoDB/data-services)

`cdb_geocoder` Archived CARTO geocoder extension for administrative, postal, IP, and named-place lookups.

### Enablement

Install the files for the intended server, then create `cdb_geocoder` in the target database:

```sql
CREATE EXTENSION cdb_geocoder CASCADE;
```

The reviewed control or official workflow requires `cartodb`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
set statement_timeout = '20min';
INSERT INTO ip_address_locations (the_geom, network_start_ip) SELECT the_geom, ('::ffff:' || split_part(network, '/', 1))::inet FROM latest_ip_address_locations;
INSERT INTO ip_address_locations (the_geom, network_start_ip) SELECT the_geom, split_part(network, '/', 1)::inet FROM latest_ip6_address_locations;
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- Catalog lifecycle is archived; test upgrades, dump/restore, and server compatibility before production use.
