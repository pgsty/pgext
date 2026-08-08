## Usage

Sources:

- [pg_mentat v1.5.7 README](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/README.md)
- [pg_mentat v1.5.7 control file](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/pg_mentat.control)
- [pg_mentat v1.5.6 to v1.5.7 upgrade SQL](https://codeberg.org/gregburd/pg_mentat/src/tag/v1.5.7/pg_mentat/sql/pg_mentat--1.5.6--1.5.7.sql)
- [Pigsty package matrix](https://pgext.cloud/ext/pg_mentat)

`pg_mentat` implements a Datomic-compatible data model and Datalog query engine inside PostgreSQL. It stores immutable facts as typed datoms and exposes schema transactions, Datalog queries, pull expressions, time travel, transaction history, and permanent excision through SQL functions. Use it for applications that need this model; it is not a transparent replacement for relational tables or SQL.

### Install and Define a Schema

```sql
CREATE EXTENSION pg_mentat;

SELECT mentat.t('[
  {:db/ident       :person/name
   :db/valueType   :db.type/string
   :db/cardinality :db.cardinality/one}
  {:db/ident       :person/age
   :db/valueType   :db.type/long
   :db/cardinality :db.cardinality/one}
]');
```

The recommended convenience aliases live in schema `mentat`. Schema must be transacted before facts use the new attributes.

### Transact and Query Data

```sql
SELECT mentat.t('[
  {:person/name "Alice" :person/age 30}
  {:person/name "Bob"   :person/age 25}
]');

SELECT mentat.q('
  [:find ?name ?age
   :where [?e :person/name ?name]
          [?e :person/age ?age]
          [(> ?age 28)]]
');
```

`mentat.t(edn)` applies an ACID transaction and returns its transaction report. `mentat.q(query, inputs)` compiles a Datalog query to PostgreSQL execution. Use EDN parameters and input bindings rather than interpolating application strings into a query.

### Pull, History, and What-If Transactions

```sql
SELECT mentat.pull('[*]', 10001);
SELECT mentat.log('default', 1000001, 1000010);
SELECT mentat.diff('default', 1000003, 1000007);

SELECT mentat.mentat_with('[
  {:person/name "Alice" :person/age 31}
]');
```

`mentat.pull` returns entity-shaped JSON. `mentat.log` and `mentat.diff` expose transaction history, and `mentat.mentat_with` evaluates a transaction without persisting it. Queries can also be evaluated as of or since a transaction by using the documented database arguments.

Permanent excision is intentionally separate from normal immutable history:

```sql
SELECT mentat.mentat_excise('default', 10042, NULL);
```

Review the target entity and backups before excision; it permanently removes datoms and is intended for requirements such as privacy erasure.

### Important Objects

- `mentat.t(edn)`: transact schema or data.
- `mentat.q(query, inputs)`: execute Datalog.
- `mentat.pull(pattern, eid)` and `mentat.pull_many(pattern, eids)`: entity-shaped reads.
- `mentat.entity(eid)` and `mentat.schema()`: inspect an entity or current schema.
- `mentat.log(...)` and `mentat.diff(...)`: inspect transaction history.
- `mentat.stats()`, `mentat.storage()`, and `mentat.cache_stats()`: operational inspection.
- `mentat.subscribe(...)`: reactive query notifications through PostgreSQL `LISTEN`/`NOTIFY`.

The extension stores typed datoms in narrow tables under schema `mentat`, including reference, integer, string, boolean, floating-point, instant, keyword, UUID, and byte values.

### Requirements and Caveats

- Upstream v1.5.7 supports PostgreSQL 13-18. Current Pigsty packages target PostgreSQL 14-18 and are rebuilt with pgrx 0.19.1; upstream's tagged source declares pgrx 0.17. Treat the packaged binary as the compatibility boundary.
- The extension is not relocatable and does not require `shared_preload_libraries`.
- The optional `mentatd` HTTP/Datomic-wire daemon is an upstream companion program and is not included in the Pigsty `pg_mentat` package. SQL use of the extension does not require it.
- Datalog compilation, pull recursion, full-text attributes, subscriptions, and history can have very different cost profiles. Inspect generated SQL with the documented explain helper and benchmark representative data.
- Excision bypasses the normal immutable-history model. Restrict privileges and audit its use.
