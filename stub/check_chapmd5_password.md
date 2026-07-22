## Usage

Sources:

- [Official README](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/README.md)
- [Version 0.0.1 extension SQL](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password--0.0.1.sql)
- [C implementation](https://github.com/sshutdownow/pg_check-chap-md5/blob/5371628350775b7d51ac097bee24913705856fb1/check_chapmd5_password.c)
- [RFC 1994: PPP Challenge Handshake Authentication Protocol](https://www.rfc-editor.org/rfc/rfc1994)

`check_chapmd5_password` verifies an RFC 1994 CHAP-MD5 response against a challenge and a cleartext secret. The package is named `pg_check_chap_md5`, but the canonical PostgreSQL extension and SQL function are both named `check_chapmd5_password`.

### Core Workflow

```sql
CREATE EXTENSION check_chapmd5_password;

SELECT check_chapmd5_password(
    '00777f2a3f6a2e661947b520c6777e0b25',
    '45c915d82d67257209048420a31292d3',
    'password'
);
```

The first argument is the hexadecimal CHAP response, including its one-byte identifier; the second is the hexadecimal challenge; the third is the cleartext shared secret. The function returns `true` when the response matches and `false` otherwise.

### Object Index

- `check_chapmd5_password(chap_password text, chap_challenge text, clear_password text) RETURNS boolean` performs the verification.

### Operational Notes

Version `0.0.1` is relocatable and declares no extension dependency or preload requirement. The function is `STRICT` and `IMMUTABLE`.

Treat the third argument as a credential: restrict `EXECUTE`, avoid exposing calls through query text or monitoring, and do not enable `DEBUG2` logging because the implementation logs the cleartext password at that level. Inputs are decoded by a permissive hexadecimal parser; malformed characters produce warnings rather than robust validation. Validate response length, even-length hexadecimal encoding, and challenge size before calling the function. CHAP-MD5 itself is a legacy protocol and does not replace transport security or modern password storage.
