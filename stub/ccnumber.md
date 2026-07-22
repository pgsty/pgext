## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/README.md)
- [Version 1.0 SQL objects](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber--1.0.sql)
- [Backend implementation](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber.c)
- [Comparator implementation](https://github.com/tvondra/ccnumber/blob/ec41b073c9d601c164be9798b5e126eda45a41fa/ccnumber-comparator.c)

`ccnumber` is an experimental encrypted byte-string type. PostgreSQL keeps ciphertexts while a separate TCP comparator decrypts and compares them, enabling equality, ordering, B-tree and hash operator classes, and `min`/`max`. Upstream explicitly describes the project as a minimally functional, inefficient, generally untested proof of concept; do not use it for payment data or other secrets.

### Core Workflow

The project has no SQL plaintext-encryption function. Generate ciphertext with the separately built upstream tool, start `ccnumber-comparator`, and then provide the ciphertext as bytea-style hex input:

```sql
CREATE EXTENSION ccnumber;
CREATE TABLE payment_token (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number ccnumber UNIQUE
);

SHOW ccnumber.comparator_host;
SHOW ccnumber.comparator_port;
SHOW ccnumber.optimize_remote_calls;
```

`ccnumber.comparator_host` defaults to `127.0.0.1`, `ccnumber.comparator_port` to `9999`, and `ccnumber.optimize_remote_calls` to `true`. Every comparison can depend on the external service; an unavailable comparator makes relevant queries fail.

### Security and Operational Boundaries

The reviewed implementation hard-codes the encryption key in the comparator. Its TCP protocol has neither authentication nor transport encryption, so the process acts as a remotely accessible decryption/comparison oracle unless it is strictly isolated. Ordering leaks order; the optimization and hash index leak equality-related information; ciphertexts also add storage overhead.

Both the backend and comparator copy attacker-influenced values into fixed-size buffers without adequate length checks. Arbitrary or oversized `ccnumber` input can therefore corrupt memory. Treat this code only as a design study in a disposable environment, not as a security control or production extension.
