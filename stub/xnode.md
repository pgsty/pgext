## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/README)
- [Official extension control file (xnode.control)](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/src/xnode.control)
- [Official extension SQL (xnode.sql)](https://api.pgxn.org/src/pg_xnode/pg_xnode-0.7.2/src/sql/xnode.sql)

`xnode` — Implementation of XML using DOM. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION xnode;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add(doc, path, node, add_mode)` is an extension function and returns `doc`.
- `children(node)` is an extension function and returns `node[]`.
- `doc_in(cstring)` is an extension function and returns `doc`.
- `doc_out(doc)` is an extension function and returns `cstring`.
- `doc_to_node(doc)` is an extension function and returns `node`.
- `element(text, text[][2], node)` is an extension function and returns `node`.
- `fragment_sfunc(node, node)` is an extension function and returns `node`.
- `node(xnt, text[], record)` is an extension function and returns `node`.
- `node_debug_print(node)` is an extension function and returns `text`.
- `node_in(cstring)` is an extension function and returns `node`.
- `node_kind(node)` is an extension function and returns `text`.
- `node_out(node)` is an extension function and returns `cstring`.
- `node_to_doc(node)` is an extension function and returns `doc`.
- `path(path, doc)` is an extension function and returns `pathval`.

### Requirements and Caveats

- The reviewed control file declares default version `0.7.2`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
