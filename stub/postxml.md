## Usage

Sources:

- [Official documentation](https://github.com/albaike/postxml/blob/cc25905bf2fc0e77b65235ac10f67794ab4e61cd/README.md)
- [Extension control file](https://github.com/albaike/postxml/blob/cc25905bf2fc0e77b65235ac10f67794ab4e61cd/postxml.control)
- [Official repository](https://github.com/albaike/postxml)

`postxml` XML and XSLT transformation helpers for building hypermedia documents in PostgreSQL.

### Enablement

Install the files for the intended server, then create `postxml` in the target database:

```sql
CREATE EXTENSION postxml;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT postxml.xml_docs_equal('<a/>'::xml, '<a/>'::xml);
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `postxml.init_xsd` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.read_text` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.xml_docs_equal` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.xsd_schemas` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `postxml.xsd_validate` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.xsd_validate_name` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.xsl_transform` | FUNCTION | Callable function from the reviewed install surface. |
| `postxml.xsl_transform_name` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 16; do not infer unlisted majors.
