## Usage

Sources:

- [TencentDB for PostgreSQL extension version matrix](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB: Creating an extension in an instance](https://cloud.tencent.com/document/product/409/121744)
- [TencentDB example showing the preloaded module](https://cloud.tencent.com/document/product/409/111199)

`tencentdb_pwdcheck` is TencentDB for PostgreSQL's provider-only strong-password authentication extension. The current support matrix lists version `1.0` for PostgreSQL 10 through 17 and no PostgreSQL 18 package. Tencent does not publish its source, control file, password-policy settings, or SQL API.

Check the exact instance's catalog before enabling it:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'tencentdb_pwdcheck';
```

TencentDB documentation shows `tencentdb_pwdcheck` in `shared_preload_libraries`. In the managed console, add the module to that parameter if it is not already present, save the parameter change, and allow the instance to restart. Then create the extension in each database where its extension object is required:

```sql
CREATE EXTENSION tencentdb_pwdcheck;
```

Changing `shared_preload_libraries` restarts the managed instance, so schedule the operation and ensure clients reconnect automatically. Because the provider does not document individual policy knobs publicly, use the console and Tencent support information for the selected engine revision instead of assuming community `passwordcheck` settings apply.
