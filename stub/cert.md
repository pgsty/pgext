## Usage

Sources:

- [Official extension documentation](https://github.com/beargiles/pg-cert/blob/555b021861e0c8a52b8fcfb08e732945e56f5d82/doc/cert.md)
- [Official extension SQL](https://github.com/beargiles/pg-cert/blob/555b021861e0c8a52b8fcfb08e732945e56f5d82/sql/cert.sql)
- [Official PGXN distribution](https://pgxn.org/dist/cert/)

`cert` 0.1.1 adds PostgreSQL types for OpenSSL-format X.509 certificates, certificate requests, and DSA parameters, plus functions that extract certificate metadata. Upstream describes the implementation as preliminary and not intended for production use.

### Core Workflow

The install SQL requires the `bignum` extension. Store PEM certificates in a `cert` column, then use the accessor functions in ordinary queries:

```sql
CREATE EXTENSION bignum;
CREATE EXTENSION cert;

CREATE TABLE certificates (
  certificate cert NOT NULL
);

SELECT get_version(certificate),
       get_serial_number(certificate),
       get_not_before(certificate),
       get_not_after(certificate),
       get_issuer(certificate),
       get_subject(certificate)
FROM certificates;
```

Values are supplied through the `cert` input function as complete OpenSSL PEM certificates. Invalid or unsupported certificate text is rejected by the C input routine.

### Important Objects

- `cert` stores an X.509 certificate; `cert_req` stores a certificate request.
- `basic_cert_info` is a composite of serial number, validity interval, issuer, and subject; `get_basic_info(cert)` returns it.
- `get_version(cert)`, `get_serial_number(cert)`, `get_not_before(cert)`, `get_not_after(cert)`, `get_issuer(cert)`, and `get_subject(cert)` expose basic certificate fields.
- `get_signature(cert)`, `get_issuer_hash(cert)`, `get_iands_hash(cert)`, `get_common_name(cert)`, `get_subject_hash(cert)`, and `get_subject_keyid_hash(cert)` expose additional text properties.
- `dsa_params`, `size(dsa_params)`, and `generate_dsa_params(integer)` support DSA parameter storage and generation.

### Operational Notes

Only OpenSSL-format input is supported. The SQL comments note that issuer and subject currently return text rather than a distinguished-name type, and several public-key operations remain unimplemented. The C module links against OpenSSL and the install script creates `bignum`; verify both dependencies on the target PostgreSQL build. Do not treat stored certificates as automatically validated trust objects: the documented API parses and inspects them, but upstream does not document chain validation or revocation checking.
