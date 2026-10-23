## Usage

Sources:

- [Official source](https://docs.opengauss.org/en/docs/7.0.0/extension_reference/extension_reference/server/gms_xmlgen.html)
- [Official source](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlgen/gms_xmlgen.control)
- [Official source](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlgen/gms_xmlgen--1.0.sql)

`gms_xmlgen` 1.0 converts query results into XML on openGauss. It is specific to that kernel and its Oracle-compatible A-mode databases, requires libxml, and is unavailable in lightweight builds. This does not establish compatibility with stock PostgreSQL.

### Core Workflow

With the matching openGauss extension files installed and sufficient installation privileges:

```sql
CREATE EXTENSION gms_xmlgen;
SELECT gms_xmlgen.getxml('SELECT 1 AS id');
```

The API returns CLOB or XMLTYPE values. `gms_xmlgen.newcontext` creates a reusable query context, `gms_xmlgen.getxml` fetches XML, and `gms_xmlgen.closecontext` releases the context. Limit and skip rows with `gms_xmlgen.setmaxrows` and `gms_xmlgen.setskiprows`. Row tags, rowset tags and null handling are configurable through the corresponding context helpers.

### Compatibility and Maintenance

XMLTYPE interfaces require `behavior_compat_options = 'bind_procedure_searchpath'`. The required XMLTYPE support was introduced in openGauss 7.0.0-RC1. This is SQL enablement, not a preload requirement.

Review query privileges and avoid passing untrusted SQL strings to the generator. Release contexts promptly. The vendor requires removing this extension before affected kernel upgrade or rollback procedures because of XMLTYPE dependencies; inventory dependent objects and rehearse restoration before following that procedure on a populated database.
