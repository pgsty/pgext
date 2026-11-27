## Usage

Sources:

- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/README.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/pg_bao.control)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/pg_bao--0.0.1.sql)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/pg_extension/main.c)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/docs/src/tutorial/1_pg_setup.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/docs/src/tutorial/2_bao_setup.md)
- [Official source](https://github.com/LearnedSystems/BaoForPostgreSQL/blob/1c15600a6d0a5bf09b08405702889eb1783579ca/LICENSE)

`pg_bao` is the Bao learned-query-optimizer prototype. Its control version is 0.0.1; the accompanying SQL script is empty. The official tutorial demonstrates PostgreSQL 12.3. No support for current PostgreSQL majors is implied by that historical example.

### Activate the Hook

After installing a matching library in a research cluster, preserve existing preload entries, add the module and restart PostgreSQL. Start the separate Bao model server before using it:

```ini
shared_preload_libraries = 'pg_bao'
```

```sql
SHOW enable_bao;
SET bao_host = 'localhost';
SET bao_port = 9381;
SET enable_bao = on;
EXPLAIN SELECT count(*) FROM existing_table;
SET enable_bao = off;
```

Replace the example table with an existing relation. Activation is through preload and session settings, not through database SQL objects. `enable_bao` defaults off. The separate `enable_bao_selection` and `enable_bao_rewards` settings control plan selection and latency-feedback collection when it is enabled. The Bao server maintains and retrains the model.

### Observe Behavior

`EXPLAIN` includes Bao's prediction and recommended hint when enabled. An untrained server may report no prediction. Training needs representative completed workloads and sufficient observations; compare total planning plus execution time to the ordinary optimizer. The implementation considers multiple hint alternatives sequentially, which adds planning cost.

### Security and Maintenance

The model server provides neither authentication nor encryption. Restrict it to a trusted network and intended clients. Although predicate literals are omitted from its plan representation, relation names, plan shapes and execution measurements still leave PostgreSQL. Upstream marks the code as a prototype not audited for production security.

Preserve training/model state separately from PostgreSQL backups. Removing preload takes a server restart; disabling the session setting is the immediate way to stop using its optimizer path in that session. The project supplies no SQL migration chain or verified modern-major compatibility matrix. The implementation is licensed under AGPLv3.
