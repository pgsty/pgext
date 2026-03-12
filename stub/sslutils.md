


## Usage

> [sslutils: Manage SSL certificates through SQL](https://github.com/EnterpriseDB/sslutils)

`sslutils` is a PostgreSQL extension for managing SSL certificates through SQL commands. It provides functions to generate, inspect, and manage SSL/TLS certificates directly within the database.

```sql
CREATE EXTENSION sslutils;
```

### Functions

The extension provides SQL functions for SSL certificate management:

| Function | Description |
|----------|-------------|
| `openssl_rsa_generate_key(bits int)` | Generate an RSA private key |
| `openssl_rsa_key_to_csr(key text, cn text, ...)` | Generate a Certificate Signing Request |
| `openssl_csr_to_crt(csr text, ca_key text, ca_crt text)` | Sign a CSR to produce a certificate |
| `openssl_rsa_generate_crl(ca_key text, ca_crt text)` | Generate a Certificate Revocation List |
| `ssl_is_init_fn()` | Check if SSL is initialized |
| `ssl_get_cipher_fn()` | Get current SSL cipher |
| `ssl_get_version_fn()` | Get current SSL version |

### Typical Workflow

```sql
-- Generate a CA private key
SELECT openssl_rsa_generate_key(2048);

-- Create a self-signed CA certificate
-- Generate server key and CSR
-- Sign the CSR with the CA
```

This extension is useful for automating SSL certificate provisioning in managed PostgreSQL environments.
