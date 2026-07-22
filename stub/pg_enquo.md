## Usage

Sources:

- [Official README](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/README.md)
- [Official data-type documentation](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/doc/data_types/README.md)
- [Official extension control file](https://github.com/enquo/pg_enquo/blob/9ae68398212302f877631784131ee93e30badf96/pg_enquo.control)

`pg_enquo` provides PostgreSQL types and operators for storing and comparing ciphertexts produced by an Enquo client library. Encryption and query-token construction happen in the application; the extension never receives the plaintext or encryption key as part of its normal query workflow.

### Core Workflow

Create the extension and model columns with the encrypted counterpart of the required PostgreSQL type:

```sql
CREATE EXTENSION pg_enquo;

CREATE TABLE private_records (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  secret_number enquo_bigint NOT NULL,
  secret_date enquo_date,
  secret_text enquo_text
);
```

Generate storage ciphertexts and query ciphertexts with an official Enquo client library, then bind them as parameters rather than constructing ciphertext JSON in SQL:

```sql
INSERT INTO private_records(secret_number, secret_date, secret_text)
VALUES (
  :'number_ciphertext'::enquo_bigint,
  :'date_ciphertext'::enquo_date,
  :'text_ciphertext'::enquo_text
);

SELECT id
FROM private_records
WHERE secret_text = :'text_query_ciphertext'::enquo_text;
```

The client must use the same root key, field identity, schema context, data type, and storage options for stored and query ciphertexts.

### Supported Type Surface

- `enquo_bigint` stores encrypted signed 64-bit integers and supports comparisons. Its documented default representation is roughly 500 bytes per value.
- `enquo_date` stores encrypted dates with years from `-32768` through `32767` and supports comparisons. Its documented default representation is roughly 250 bytes.
- `enquo_text` stores arbitrary UTF-8 text and supports equality/inequality. Its size starts around 110 bytes for an empty value and grows in encrypted blocks.
- Client option `no_query` reduces storage when values never need querying, at the cost of removing query tokens.

### Security Trade-offs

- Default security should be the baseline. Client option `enable_reduced_security_operations` adds deterministic/order information needed for indexes or `ORDER BY`, but leaks equality and, for ordered types, the ordering of the column; this can enable inference attacks.
- `enquo_bigint` can use B-tree or hash indexes and ordering only with the required reduced-security material. `enquo_date` similarly enables B-tree/order operations; `enquo_text` enables hash indexing. Match the index to the client storage options or operations will fail.
- Ciphertext does not hide table shape, row count, nullness, access patterns, or query frequency. Define the threat model, rotate and back up root keys, and prevent logs/traces from recording plaintext before adopting the scheme.
- Version `0.0` is an early interface. Pin the client library, `enquo-core`, extension build, schema field identities, and storage options as one compatibility unit; test restore, key loss, schema changes, re-encryption, index rebuilds, malformed ciphertext, and upgrades.
