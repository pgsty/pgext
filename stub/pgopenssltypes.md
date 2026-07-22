## Usage

Sources:

- [Official extension documentation](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/doc/pgopenssltypes.md)
- [Official version 0.0.1 SQL](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/sql/pgopenssltypes--0.0.1.sql)
- [Official function reference](https://github.com/beargiles/pgopenssltypes/blob/7f89bf03d6c0d7ee7a1fbac03276490d575b0037/FUNCTIONS.md)

`pgopenssltypes` adds PostgreSQL types for OpenSSL key, certificate, request, revocation-list, and cryptographic-message structures, plus digest and X.509 accessor functions. It is intended for certificate-authority or certificate-inventory prototypes that need to parse and query PEM objects inside PostgreSQL.

### Core Workflow

```sql
CREATE EXTENSION pgopenssltypes;

CREATE TABLE certificates (
    id bigserial PRIMARY KEY,
    certificate x509 NOT NULL
);

SELECT x509_get_subject_name(certificate),
       x509_get_issuer_name(certificate),
       x509_get_not_before(certificate),
       x509_get_not_after(certificate)
FROM certificates;

SELECT dgst_sha256('message to hash');
```

Important types include `X509`, `X509_NAME`, `PKEY`, `PKCS8`, `PKCS7`, `X509_REQ`, `X509_CRL`, and arbitrary-precision `BN`. `RSA`, `DSA`, and `DSA_PARAMS` are retained for historical formats but upstream advises against them because private keys are stored unencrypted. `PKCS12` is declared but documented as not implemented.

X.509 accessors expose version, serial number, validity bounds, subject, issuer, public key, aliases, hashes, and key identifiers. `x509_check_private_key(X509, PKEY)` checks whether a key matches a certificate. Digest availability depends on the linked OpenSSL build; MD4, MD5, SHA-1, and other legacy algorithms must not be used for new security designs.

### Security and Maintenance Notes

Most values use PEM text input/output and can contain private key material. PostgreSQL tables, WAL, replicas, logs, dumps, and backups may therefore hold secrets; use strict privileges and external encryption/key-management controls. SQL-visible parsing is not a substitute for certificate-path validation or a managed HSM. The upstream documentation explicitly says that support is unavailable, and version `0.0.1` is old; validate it against the exact PostgreSQL and OpenSSL ABI and audit parsers before exposing them to untrusted inputs.
