## Usage

Sources:

- [Official README](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/README.md)
- [Extension implementation](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/src/lib.rs)
- [Cargo and PostgreSQL compatibility](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/Cargo.toml)

`pgmer1` 0.1.0 is a pgrx HTTP client for a separate MeritRank service. It can submit weighted graph edges and query ranking scores from SQL. The endpoint is hard-coded as `http://localhost:8000`; there is no GUC, per-call URL, authentication, or TLS configuration.

### Core Workflow

Run a compatible MeritRank service on the PostgreSQL server's loopback interface, then:

```sql
CREATE EXTENSION pgmer1;

SELECT mr_service_url();

SELECT mr_edge('alice', 'bob', 1.0);

SELECT * FROM mr_scores('alice');

SELECT mr_node_score('alice', 'bob');
```

`mr_edge(src, dest, weight)` sends an HTTP `PUT /edge` JSON request and returns the response's message field. `mr_scores(ego)` sends `GET /scores/{ego}` and returns rows `(node text, ego text, score double precision)`. `mr_node_score(ego, target)` calls `GET /node_score/{ego}/{target}` and returns one score. `mr_service_url()` reveals the fixed endpoint.

### Transaction and Failure Boundaries

Calls are synchronous inside the PostgreSQL backend. Slow or unavailable HTTP service calls hold the SQL session, and the implementation does not configure a request timeout. The GET paths unwrap transport setup errors, which can turn connectivity failures into a Rust panic rather than a controlled SQL error.

`mr_edge` changes an external service immediately and cannot participate in PostgreSQL rollback: retrying after an ambiguous timeout may duplicate a change unless the service makes the operation idempotent. Graph identifiers are inserted directly into URL path segments without explicit percent-encoding in this extension. Accept trusted identifiers only, restrict execution privileges, and isolate the loopback service from untrusted local processes.

The pinned Cargo features cover PostgreSQL 11–15 and use an old pgrx release; the project is cataloged as abandoned. Rebuild and test against the exact server and MeritRank API revision, and prefer application-side integration when authentication, observability, timeout control, or reliable delivery matters.
