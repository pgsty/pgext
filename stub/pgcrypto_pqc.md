## Usage

Sources:

- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.0.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.0.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.0--1.1.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.0--1.1.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.1--1.2.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.1--1.2.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.2--1.3.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.2--1.3.sql)
- [contrib/pgcrypto_pqc/sql/pgcrypto_pqc.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/sql/pgcrypto_pqc.sql)
- [contrib/pgcrypto_pqc/meson.build](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/meson.build)
- [README.md](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/README.md)
- [contrib/pgcrypto_pqc/pgcrypto_pqc.control](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc.control)

`pgcrypto_pqc` is the SQL cryptography extension bundled with the FortressQL fork of PostgreSQL 17. Its control version is 1.3, reached through the supplied base SQL and upgrade scripts. It exposes KEM, signature, encryption and algorithm-discovery functions; it does not by itself enable the fork’s TLS, TDE or WAL features.

### Core Workflow

```sql
CREATE EXTENSION pgcrypto_pqc;
SELECT name, type, nist_level FROM pqc_algorithms() ORDER BY name;
SELECT length(public_key) AS public_key_bytes,
       length(secret_key) AS secret_key_bytes
FROM pqc_kem_keygen('ML-KEM-768');
```

### Operational Boundaries

Use a FortressQL build configured with post-quantum support and liboqs/OpenSSL. This catalog entry does not assert compatibility with an unmodified PostgreSQL server. Installation requires a superuser; the control file is relocatable and does not require extension-specific preloading.

`pqc_kem_keygen`, `pqc_kem_encapsulate` and `pqc_kem_decapsulate` implement key encapsulation. `pqc_sig_keygen`, `pqc_sign` and `pqc_verify` expose signing. `pqc_encrypt` and `pqc_decrypt` implement the documented hybrid encryption interface; `pqc_algorithms` lists available algorithms. Later SQL revisions add audit, column-encryption, escrow metadata and reporting objects tied to FortressQL.

Secret keys are returned to the SQL caller and handled on the server. Protect their storage, transport, logs and execution permissions; keep algorithms consistent across paired operations. The presence of an algorithm or a reporting function is not evidence of an independent security or compliance certification. Review the fork’s operational documentation before using the wider encryption and audit features.
