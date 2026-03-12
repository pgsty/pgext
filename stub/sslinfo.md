

## Usage

> [sslinfo: SSL certificate information functions](https://www.postgresql.org/docs/current/sslinfo.html)

sslinfo provides functions to access information about the SSL certificate used in the current database connection.

### Functions

```sql
-- Check if current connection uses SSL
SELECT ssl_is_used();

-- SSL/TLS protocol version (TLSv1.2, TLSv1.3, etc.)
SELECT ssl_version();

-- Cipher name (e.g., DHE-RSA-AES256-SHA)
SELECT ssl_cipher();

-- Check if client presented a valid certificate
SELECT ssl_client_cert_present();

-- Client certificate serial number
SELECT ssl_client_serial();

-- Client certificate subject (full DN)
SELECT ssl_client_dn();
-- e.g., /CN=Somebody /C=Some country/O=Some organization

-- Certificate issuer (full DN)
SELECT ssl_issuer_dn();

-- Specific field from client certificate subject
SELECT ssl_client_dn_field('CN');
SELECT ssl_client_dn_field('O');

-- Specific field from certificate issuer
SELECT ssl_issuer_field('CN');

-- Client certificate extensions
SELECT * FROM ssl_extension_info();
-- Returns: name, value, critical
```

### Notes

- Most functions return NULL if the connection does not use SSL
- Requires PostgreSQL compiled with OpenSSL support
- The combination of `ssl_client_serial()` and `ssl_issuer_dn()` uniquely identifies a certificate
