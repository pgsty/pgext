## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/README.md)
- [Utility implementation](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/src/utility_lib.rs)
- [Networking implementation](https://github.com/isdaniel/rust_pg_extensions/blob/f25df80c4ce06bf81bb7e3a58fc5087ea85c64ac/src/networking_lib.rs)

`all_in_one_lib` is a pgrx experiment combining hash helpers, AES-GCM encrypt/decrypt functions, server network discovery, a logical-decoding output plugin, and example foreign-data wrappers in one superuser-only extension.

```sql
CREATE EXTENSION all_in_one_lib;
SELECT compute_hash('payload', 'sha256');
SELECT get_server_hostname(), get_server_ip();
```

Do not use `data_encrypt` for real confidentiality. The reviewed source uses one hard-coded nonce for every AES-256-GCM operation; nonce reuse under the same key breaks GCM security. It also accepts raw keys from SQL, where values can leak through queries, logs, or monitoring. Existing ciphertext needs a separately designed migration, not continued use.

The network functions expose host identity and non-loopback addresses. The decoder implements only a narrow set of tuple types and contains unfinished paths, while the FDWs are demonstrations. Revoke all functions from `PUBLIC`, grant only individually after review, and do not install this mixed experimental library on production. Prefer focused, maintained extensions and external key management for each requirement.
