## Usage

Sources:

- [Tencent Cloud PostgreSQL: using starrocks_fdw](https://cloud.tencent.com/document/product/409/111197)

`starrocks_fdw` is a Tencent Cloud PostgreSQL provider extension that exposes StarRocks tables as PostgreSQL foreign tables for hot/cold data separation. It is compatible with the `mysql_fdw` usage model but includes Tencent Cloud connection-routing options; this is a managed-provider feature, not a portable upstream extension workflow.

### Core Workflow

The StarRocks service must already be reachable from the Tencent Cloud PostgreSQL instance. Create the extension, a foreign server, a user mapping, and a foreign table that names the StarRocks database and table.

```sql
CREATE EXTENSION starrocks_fdw;

CREATE SERVER server_sr
  FOREIGN DATA WRAPPER starrocks_fdw
  OPTIONS (
    host '10.0.0.10', port '9030',
    instanceid 'ins-example', access_type '2',
    region 'ap-guangzhou', uin '100000000000',
    own_uin '100000000000', vpcid 'vpc-example',
    subnetid 'subnet-example'
  );

CREATE USER MAPPING FOR dbadmin
  SERVER server_sr
  OPTIONS (username 'root', password 'replace-me');

CREATE FOREIGN TABLE sr_m (
  id integer NOT NULL,
  str integer
)
SERVER server_sr
OPTIONS (dbname 'my_database', table_name 'sr_m1');

SELECT * FROM sr_m;
```

Tencent Cloud's example also attaches a foreign table as a partition of a local partitioned table, allowing local and StarRocks-backed partitions to be queried through one parent.

### Provider and Security Boundaries

The official document states support only for Tencent Cloud PostgreSQL 13 builds at `v13.14_r1.17` or later. Availability, server-option values, and network routing are controlled by the Tencent Cloud service; confirm them for the specific instance rather than assuming community PostgreSQL compatibility.

Credentials reside in the PostgreSQL user mapping, while server options contain cloud account and network identifiers. Restrict ownership and catalog visibility, use a dedicated least-privilege StarRocks account, and keep PostgreSQL and StarRocks column types aligned. Remote reads depend synchronously on StarRocks and the VPC path, so include remote latency and outages in query planning and monitoring.
