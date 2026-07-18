## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/README.md)
- [Version 1.0 SQL objects](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber--1.0.sql)
- [Extension control file](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber.control)

`ccnumber` is a proof-of-concept encrypted data type for card numbers or other sensitive byte strings. Comparisons are sent over TCP to a separate `ccnumber-comparator` process, allowing B-tree and hash indexes, ordering, equality, and `min`/`max` without decrypting values inside PostgreSQL.

```sql
CREATE EXTENSION ccnumber;
CREATE TABLE payment_token (number ccnumber PRIMARY KEY);

SHOW ccnumber.comparator_host;
SHOW ccnumber.comparator_port;
```

Start the separately built comparator before evaluating values. Its endpoint is selected with `ccnumber.comparator_host` and `ccnumber.comparator_port`; `ccnumber.optimize_remote_calls` controls a hash-based optimization.

### Security Status

Upstream explicitly calls this project minimally functional, inefficient, generally untested, and not production-ready. It uses `libsodium`, but the encryption key is hard-coded, TCP transport is not authenticated, ordering leaks information, and each ciphertext carries storage overhead. Use it only to study the trusted-component design, never for real payment data or secrets.
