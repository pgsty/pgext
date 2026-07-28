## Usage

Sources:

- [Official upstream README](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/ReadMe.md)
- [Official extension control file (kmea.control)](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/kmea.control)
- [Official extension SQL (kmea--1.0.sql)](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/kmea--1.0.sql)

`kmea` — KMEA[^1] is a PostgreSQL extension that supports various DNA data types, along with some operators. DNA sequences Kmers Qkmers. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION kmea;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `canonical(kmer)` is an extension function and returns `kmer`.
- `contains(qkmer, kmer)` is an extension function and returns `boolean`.
- `DNA(text)` is an extension function and returns `DNA`.
- `dna_in(cstring)` is an extension function and returns `DNA`.
- `dna_out(DNA)` is an extension function and returns `cstring`.
- `dna_recv(internal)` is an extension function and returns `DNA`.
- `dna_send(DNA)` is an extension function and returns `bytea`.
- `equals(kmer, kmer)` is an extension function and returns `boolean`.
- `generate_kmers(DNA, integer)` is an extension function and returns `SETOF`.
- `kmer(text)` is an extension function and returns `kmer`.
- `kmer_hash(kmer)` is an extension function and returns `integer`.
- `kmer_in(cstring)` is an extension function and returns `kmer`.
- `kmer_out(kmer)` is an extension function and returns `cstring`.
- `kmer_recv(internal)` is an extension function and returns `kmer`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
