## Usage

Sources:

- [1.4.2 README](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/README.sslutils)
- [1.4.2 SQL API](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils--1.4.2.sql)
- [1.4.2 control file](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils.control)
- [File-access checks](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils.c)
- [1.4.1 to 1.4.2 upgrade](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils--1.4.1--1.4.2.sql)
- [PEM agent grants](https://github.com/EnterpriseDB/sslutils/blob/v1.4.2/sslutils-pemagent.sql.in)

`sslutils` provides SSL certificate-generation and revocation helpers for Postgres Enterprise Manager and other controlled administration workflows. Some functions read or modify files on the database server.

### Install and Inspect

```sql
CREATE EXTENSION sslutils;
SELECT sslutils_version();
SELECT openssl_get_crt_expiry_date('server.crt');
```

A superuser must install this non-relocatable extension. No preload or PostgreSQL restart is required to call the functions; deploying generated certificates is a separate server configuration operation.

### Important Functions

| Function | Purpose |
|---|---|
| `openssl_rsa_generate_key(bits)` | Return an RSA private key as text |
| `openssl_rsa_key_to_csr(key, cn, country, state, location, unit, email)` | Return a CSR |
| `openssl_csr_to_crt(csr, ca_cert_path, private_key_path, days)` | Sign a CSR using server-side certificate and key files |
| `openssl_rsa_generate_crl(ca_cert_path, ca_key_path, days)` | Return a certificate revocation list |
| `openssl_is_crt_expire_on(cert_path, at_time)` | Check certificate expiry; returns 1, -1 or 0 |
| `openssl_get_crt_expiry_date(cert_path)` | Return the expiry timestamp |
| `openssl_revoke_certificate(cert_pem, crl_path, days)` | Revoke a certificate and regenerate the CRL |

For self-signed generation, the certificate-path argument to the signing function is NULL and the key path selects the signing key. The optional validity defaults to 3650 days. File paths refer to the PostgreSQL server, not the SQL client's machine.

### Protect Keys and Server Files

In 1.4.2, all seven `openssl_*` functions use `SECURITY DEFINER` with a fixed `search_path` of `pg_catalog, pg_temp`. Their owner, `sslutils_owner`, is a cluster-wide `INHERIT`, `NOLOGIN` role. It holds `pg_read_server_files` and `pg_write_server_files` on PostgreSQL 11 and later, but must have no members, administrative attributes or other role memberships. Installation rejects an unsafe existing owner role.

`PUBLIC` has no execution privilege. Installation or upgrade grants `EXECUTE` to `pem_agent` if that role already exists; existing explicit function grants remain. For an agent created later, a superuser runs the installed `sslutils-pemagent.sql` in the extension's schema. An agent using a non-public schema also needs `USAGE` on that schema, without `CREATE`. Callers need the relevant function privileges, not direct server-file role membership; never grant them membership in `sslutils_owner`.

Grant access only to callers trusted to inspect certificates and use CA keys within the permitted paths. Confinement to the data directory does not separate agents or CAs from each other.

Private keys returned as SQL text can enter query results, logs, shell history or application traces. Protect their handling and storage. This extension does not provide SSL session-inspection getters or automatically configure PostgreSQL's TLS settings.

### Upgrade and Existing Grants

Upgrade an existing installation as a superuser:

```sql
ALTER EXTENSION sslutils UPDATE TO '1.4.2';
```

After migrating callers, remove obsolete broad file-role grants from agents:

```sql
REVOKE pg_read_server_files, pg_write_server_files FROM pem_agent;
```

Review direct and inherited grants to individual logins too, including `pg_execute_server_program` and superuser privileges. An extension upgrade within one database does not revoke these cluster-wide grants automatically. The shared `sslutils_owner` role also survives extension removal; remove it only after uninstalling from all databases and revoking its file-role memberships.

### File-Access Boundaries

Certificate inspection and CA signing paths must resolve within the server data directory. Revocation additionally requires `sslutils.revoke_certificate_crl_paths`, a comma-separated allowlist of CRL output directories configured by a superuser in postgresql.conf and activated by reload. In 1.4.2, checks resolve paths and enforce directory boundaries; a new output file is checked through its existing parent directory, and an existing output file is checked too.

Use dedicated CRL output directories under administrator control. Do not allow the whole data directory or directories containing configuration files: an authorized caller chooses the output filename within an allowed directory. Keep certificate and private-key paths under administrator control as well.

The revocation implementation accepts PEM certificate text as its first argument, despite older README and SQL comments describing a certificate path. It reads the server-side CA files `ca_certificate.crt` and `ca_key.key`, and creates or appends to `revoke_cert.db`. Prepare the CA files before using it. Do not infer that a broad server-file role bypasses these extension-specific checks.
