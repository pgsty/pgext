## Usage

Sources:

- [Official README](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/README.md)
- [Version 1.0 SQL API](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid--1.0.sql)
- [PostgreSQL entry points](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/pg_xid.c)
- [Identifier generation and encoding](https://github.com/iCyberon/pg_xid/blob/61214ecd1a5e469771580eb8a7d320d632a7d6d1/xid.c)

`pg_xid` 1.0 generates 12-byte identifiers with the classic MongoDB ObjectID layout: a four-byte seconds timestamp, three-byte hostname-derived value, two-byte process ID, and three-byte per-process counter. It needs no central sequence and is roughly ordered by creation time. `xid()` returns raw `bytea`; `xid_encoded()` returns a 20-character custom Base32 string.

### Generate Identifiers

```sql
CREATE EXTENSION pg_xid;

SELECT encode(xid(), 'hex') AS raw_hex;
SELECT xid_encoded() AS compact_text;
```

Choose one representation for a column and keep it consistent:

```sql
CREATE TABLE event (
  id bytea PRIMARY KEY DEFAULT xid(),
  payload jsonb NOT NULL
);
```

The binary form is 12 bytes; its ordinary hex rendering is 24 characters. The encoded function uses upstream's ordered alphabet `0-9a-v` and produces 20 characters, not MongoDB's usual hexadecimal text.

### Operational and Security Boundaries

The leading timestamp makes values only approximately time ordered: all IDs generated in the same second are ordered by host/process/counter components, not by a global clock. Clock changes, hostname reuse, containers with identical hostnames, PID reuse, counter wrap, and cloned environments all affect collision assumptions. A 24-bit counter wraps after 16,777,216 IDs in one second per initialized host/process context.

These values are not secrets and are not cryptographic random tokens. They expose creation time plus host- and process-derived bits. Do not use them for authentication, capability URLs, or information-hiding identifiers.

Upstream provides no modern PostgreSQL compatibility matrix or upgrade series. The C code depends on OpenSSL MD5 for the hostname component and PostgreSQL APIs for strong random initialization. Build and concurrency-test it on each target architecture and PostgreSQL major, and validate dump/restore plus byte/text ordering before using it as a distributed primary key.
