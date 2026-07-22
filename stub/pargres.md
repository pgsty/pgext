## Usage

Sources:

- [Official README](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/README.md)
- [Extension control file](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres.control)
- [Extension SQL](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres--0.1.sql)
- [Required PostgreSQL core patch](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres_core_changes.diff)

`pargres` is a 2018 research prototype for parallel query execution across shared-nothing PostgreSQL nodes. It is not a drop-in extension: the checked source requires a custom PostgreSQL core patch and explicitly provides neither a global snapshot nor distributed write commit.

### Experimental Setup

Build a disposable PostgreSQL tree with the supplied core patch, install matching binaries on every node, preload the module, and configure the same ordered host and port lists. The source exposes these settings:

```conf
shared_preload_libraries = 'pargres'
pargres.node = 1
pargres.nnodes = 2
pargres.hosts = 'node1,node2'
pargres.ports = '5432,5432'
pargres.eports = '1000,1000'
```

```sql
CREATE EXTENSION pargres;
SELECT * FROM relsfrag;
```

`relsfrag` records a relation name, fragmentation attribute number, and fragmentation function identifier. The SQL also exposes `isLocalValue(text,integer)` and the internal coordination helper `set_query_id(integer,integer)`. Planning and execution hooks add custom exchange nodes and use ordinary PostgreSQL connections plus separate exchange sockets.

### Prototype Limitations

Only read queries are within the stated design. Without a cluster-wide snapshot, nodes can observe different database states; without distributed commit, writes cannot be made atomic and must remain sequential outside this prototype. Host connections are assembled from host and port only and exchange traffic uses raw sockets, so authentication, encryption, firewalling, and network failure handling are deployment responsibilities.

The required patch changes hash-join and nested-loop executor behavior and targets a historical server tree. Do not apply it to a supported production PostgreSQL build without a full code audit. There are no maintained upgrade, failover, recovery, security, or compatibility guarantees. Restrict `pargres` to reproducible research environments and validate the patched server, plan correctness, node ordering, port reachability, and result consistency for every experiment.
