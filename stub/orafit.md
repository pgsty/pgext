## Usage

Sources:

- [Official source](https://github.com/ltslab/orafit/blob/v1.0.3/README.md)
- [Official source](https://github.com/ltslab/orafit/releases/download/v1.0.3/orafit-jdbc-1.0.3-runtime.tar.gz)
- [Official source](https://github.com/ltslab/orafit/blob/v1.0.3/LICENSE)
- [Official source](https://github.com/ltslab/orafit/blob/v1.0.3/COMPATIBILITY.md)
- [Official source](https://github.com/ltslab/orafit/releases/tag/v1.0.3)

orafit is the server-side SQL runtime for the Orafit JDBC compatibility layer. The 1.0.3 distribution's native control and install SQL remain 1.0.0; its release fixes are in the JDBC driver and do not require reinstalling an unchanged runtime. PostgreSQL 17 is the primary target, 16 is supported and 18 receives compatibility checks only.

### Enablement

```sql
CREATE EXTENSION orafit;
```

The native runtime installs nonrelocatable objects into orafit and needs no preload. On managed services, use the official transactional plain-SQL installer instead. Keep the matching Java 17+ Orafit driver, its declared JSqlParser dependency and a separately supplied pgJDBC on the application classpath; SQL helpers alone do not translate application SQL.

```properties
driver=io.github.orafit.jdbc.OrafitDriver
url=jdbc:orafit://localhost:5432/app
```

### Compatibility boundary

Oracle 18c XE is the compatibility reference. Supported subsets and explicit rejections are listed in COMPATIBILITY.md; accepted parsing is not proof of equivalent results. The driver defaults to autosave=always and cleanupSavepoints=true to preserve statement-error behavior, unless connection properties override them. Version 1.0.3 fixes unsafe operator and outer-join rewrites; re-run the analyzer because previously mistranslated statements may now be rejected.

After an Ora2Pg migration, orafit.apply_byte_length_semantics('public') adds BYTE-length constraints only when the source columns and source/target encodings justify that policy. Review it before loading or accepting writes. The runtime does not perform schema/data migration or provide general Oracle emulation.
