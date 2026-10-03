## Usage

Sources:

- [1.4.1 README](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/README.sslutils)
- [1.4.1 SQL API](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils--1.4.1.sql)
- [1.4.1 control file](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.control)
- [File-access checks](https://github.com/EnterpriseDB/sslutils/blob/v1.4.1/sslutils.c)

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

Certificate-inspection functions require membership in `pg_read_server_files` on PostgreSQL 11 and later; revocation checks both that role and `pg_write_server_files`. These roles provide broad server-file access, so grant them only to trusted administrators and restrict function execution as needed.

Private keys returned as SQL text can enter query results, logs, shell history or application traces. Protect their handling and storage. This extension does not provide SSL session-inspection getters or automatically configure PostgreSQL's TLS settings.

### File-Access Changes in 1.4.1

Certificate inspection and CA signing paths are checked against the server data directory; place the input files there and use server-relative paths where suitable. Revocation requires `sslutils.revoke_certificate_crl_paths`, a comma-separated allowlist of CRL output path prefixes configured in postgresql.conf and activated by reload. The revocation check compares lexical string prefixes without resolving canonical directory containment; configure narrow prefixes and protect the output location.

The revocation implementation accepts PEM certificate text as its first argument, despite older README and SQL comments describing a certificate path. It reads the server-side CA files `ca_certificate.crt` and `ca_key.key`, and creates or appends to `revoke_cert.db`. Prepare the CA files before using it. Do not infer that a broad server-file role bypasses these extension-specific checks.
