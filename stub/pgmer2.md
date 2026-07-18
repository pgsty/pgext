## Usage

Sources:

- [Connector README at the reviewed commit](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/README.md)
- [SQL function implementation](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/src/lib.rs)
- [RPC client implementation](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/src/rpc.rs)
- [Cargo metadata](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/Cargo.toml)

`pgmer2` connects SQL sessions to a separately running MeritRank graph-ranking service. Set `MERITRANK_SERVICE_URL` in the PostgreSQL server environment before backend startup; the default is `tcp://127.0.0.1:8080`.

```sql
CREATE EXTENSION pgmer2;
SELECT mr_service_url();
SELECT * FROM mr_put_edge('alice', 'bob', 1.0, '', -1);
SELECT * FROM mr_node_score('alice', 'bob', '');
```

The API manages graph contexts and edges, calculates scores, lists nodes and neighbors, and supports bulk loading through `mr_bulk_load_edges`. The service is external state: deploy, secure, monitor, and back it up independently of PostgreSQL.

Current transport is unencrypted TCP with a custom binary protocol and no authentication shown in the connector. Several read functions that perform network calls are declared `IMMUTABLE`, which can permit PostgreSQL to reuse or fold results despite changing remote state. Restrict execution, isolate the service network, and do not rely on ordinary SQL volatility or transaction semantics for remote changes.
