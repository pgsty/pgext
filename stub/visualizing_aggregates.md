## Usage

Sources:

- [Upstream README](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/README.md)
- [Rust extension entry point](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/src/lib.rs)
- [Graph renderer](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/src/graph.rs)
- [Version and PostgreSQL feature matrix](https://github.com/shortdiv/pg-visualize-aggregates/blob/659e6de3f404d8a43cdc92e41d7ddb9b74654521/Cargo.toml)

visualizing_aggregates is an unfinished pgrx prototype intended to render aggregate data and create a CodePen visualization. Version 0.0.0 exposes hello_visualizing_aggregates() and intends to expose draw_graph(text). It is not usable from the reviewed upstream commit without source repairs.

### Intended smoke test

After repairing and successfully building the extension, its only side-effect-free SQL entry point can be checked with:

```sql
CREATE EXTENSION visualizing_aggregates;
SELECT hello_visualizing_aggregates();
```

The intended draw_graph function ignores its dataset argument, queries route_name from a hard-coded climbs table, logs the selected value, renders fixed markup, and performs a blocking outbound HTTP POST to CodePen. It does not yet visualize a caller-selected aggregate.

### Caveats

- The reviewed commit does not compile: lib.rs includes a nonexistent bar_graph_template.html file, while the tree contains svg_template.html, and it calls draw_svg with three arguments although graph.rs defines a one-argument method.
- The function unwraps database, template, serialization, and network results in backend code. Missing tables, empty results, rendering failures, or remote-service failures can raise errors or panic.
- Calling draw_graph initiates outbound network traffic from a PostgreSQL backend. It can block a backend, depends on third-party availability, and crosses a security boundary; restrict execute privilege and database-host egress.
- The source logs the hard-coded query and fetched route name. Treat server logs and external requests as possible data-disclosure paths.
- The control file requires superuser and marks the extension untrusted. Upstream provides no operational documentation, release, license, meaningful tests beyond the greeting, or production compatibility claim.
