## Usage

Sources:

- [Huawei Cloud RDS for PostgreSQL supported extensions](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0045.html)

`obs_fdw` version `1.0.0` is a Huawei Cloud provider extension. The current Huawei RDS for PostgreSQL support matrix lists it for PostgreSQL 12–17 (and not for 9.5–11), but the cited public page does not document its foreign-server/table options, credentials, formats, read/write semantics, or privileges.

### Provider-safe discovery

Check the exact RDS instance before attempting installation, then inventory the installed wrapper rather than guessing options:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions WHERE name = 'obs_fdw';
CREATE EXTENSION obs_fdw;
SELECT fdwname, fdwhandler::regproc, fdwvalidator::regproc
FROM pg_foreign_data_wrapper WHERE fdwname = 'obs_fdw';
```

Treat availability as controlled by the RDS engine version, latest minor release, region, and provider rollout. This is not evidence that the extension is portable to self-managed PostgreSQL. Obtain the instance-specific Huawei option and credential contract through the console/support channel before creating servers or mappings; use a scoped OBS identity and never infer secret option names from unrelated FDWs.
