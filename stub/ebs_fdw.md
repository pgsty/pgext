## Usage

Sources:

- [Pinned upstream example](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/README.md)
- [Pinned FDW implementation](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/src/lib.rs)
- [Pinned Cargo manifest](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/Cargo.toml)
- [Pinned extension control file](https://github.com/optionfactory/olddog/blob/c4195bea2d90e00b98e1d346c5af480e4a5a0d0e/ebs_fdw/ebs_fdw.control)

ebs_fdw is a read-only experimental foreign data wrapper that calls the AWS EC2 DescribeVolumes API and exposes five EBS volume attributes. The catalog and Cargo manifest identify version 0.0.0, while the wrapper macro advertises 0.1.0; verify the actual package artifact rather than relying on one version string.

### Foreign server and table

The upstream build does not create the foreign-data-wrapper catalog object, so follow CREATE EXTENSION with the explicit handler and validator wiring:

```sql
CREATE EXTENSION ebs_fdw;

CREATE FOREIGN DATA WRAPPER ebs
  HANDLER ebs_fdw_handler
  VALIDATOR ebs_fdw_validator;

CREATE SERVER aws_ebs
  FOREIGN DATA WRAPPER ebs
  OPTIONS (region 'eu-west-3');

CREATE FOREIGN TABLE ebs_volumes (
  id text,
  name text,
  type text,
  size integer,
  encrypted boolean
) SERVER aws_ebs;

SELECT id, name, type, size, encrypted
FROM ebs_volumes;
```

The PostgreSQL server process obtains AWS credentials from the AWS SDK default provider chain. There is no user-mapping credential option in the reviewed code. The region server option is mandatory and a missing value causes a Rust panic.

### AWS, planner, and compatibility limits

Every scan calls DescribeVolumes with no filters. Although the source contains an unused equality-filter helper, begin_scan sets filters to none and ignores PostgreSQL qualifications, sort keys, and LIMIT. A selective query therefore still paginates every EBS volume visible to the AWS identity. Only id, name from the Name tag, volume type, size, and encryption state are mapped; an unrecognized foreign column returns NULL.

The wrapper creates a single-thread Tokio runtime and blocks the PostgreSQL backend while AWS requests and pagination run. SDK errors are converted to panics. Set statement timeouts, measure API throttling and latency, and do not hold application locks around scans. Cache results locally for repeated reporting.

Database roles allowed to use this server can make calls with the operating identity of the database host. Grant USAGE only to a dedicated inventory role, scope its IAM permissions to DescribeVolumes, restrict outbound network access, and audit activity in AWS rather than treating PostgreSQL query logs as the authoritative API trail.

The pinned source uses old AWS SDK releases, pgrx 0.9.7, supabase-wrappers 0.1.15, and feature flags only for PostgreSQL 11 through 15. The repository is a multi-extension experimental workspace with sparse documentation. Compile against the exact target major and libc, then test credential refresh, pagination, an empty account, throttling, permission denial, cancellation, and backend recovery before non-lab use.
