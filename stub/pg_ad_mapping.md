## Usage

Sources:

- [AWS guide to AD security-group access control](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AD.Security.Groups.html)
- [AWS Aurora PostgreSQL extension versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`pg_ad_mapping` 1.0 is an Amazon Aurora PostgreSQL managed extension that maps AWS Directory Service for Microsoft Active Directory security groups to PostgreSQL login roles. It is not a portable community package: authentication, binaries, upgrades, and operational controls are supplied by Aurora.

### Core Workflow

First configure Kerberos authentication and domain membership for the Aurora cluster. Add `pg_ad_mapping` to `shared_preload_libraries` in a custom DB cluster parameter group, reboot the writer, and perform setup as an `rds_superuser`:

```sql
SHOW shared_preload_libraries;
CREATE EXTENSION pg_ad_mapping;

ALTER ROLE "accounts-role" LOGIN;
GRANT CONNECT ON DATABASE "accounts-db" TO "accounts-role";

SELECT pgadmap_set_mapping(
  'accounts-group',
  'accounts-role',
  'S-1-5-21-3168537779-1985441202-1799118680-1612',
  10
);

SELECT * FROM pgadmap_read_mapping();
```

The SID must identify the AD security group. When a user belongs to several mapped groups, the mapping with the highest weight wins; if weights tie, the most recently added mapping wins.

### Important Objects

- `pgadmap_set_mapping(ad_group, db_role, ad_group_sid, weight)` creates a mapping. It is restricted to the primary instance and an `rds_superuser`.
- `pgadmap_read_mapping()` lists the SID, database role, weight, and AD group for all mappings.
- `pgadmap_reset_mapping(ad_group_sid, db_role, weight)` removes one exact mapping; calling `pgadmap_reset_mapping()` with no arguments removes every mapping. Reset is restricted to the primary instance and an `rds_superuser`.

### Authentication and Role Boundaries

The mapped database role must have `LOGIN` and `CONNECT` to the target database. Do not grant `rds_ad` to a group-mapped role; AWS documents `rds_ad` for individually provisioned AD users instead. Group authentication is triggered with a domain-qualified principal, and a user individually provisioned with `rds_ad` does not also log in through group mapping.

### Operational Notes

AWS documents AD group access from Aurora PostgreSQL 14.10 and 15.5, while the live extension matrix determines the exact engine releases and extension version available to a cluster. A preload change requires a writer reboot. Rules are managed per cluster, and administrative writes must run on the primary.

Microsoft Entra ID is not supported. AWS also documents a migration hazard for clusters that had Kerberos enabled before AD group support: after adding the preload library, logins can fail and the engine can potentially crash. For affected existing instances, disable and re-enable Kerberos as directed by AWS; instances created after April 2025 are not affected by that workaround requirement.
