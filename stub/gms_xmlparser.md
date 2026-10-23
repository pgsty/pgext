## Usage

Sources:

- [Official source](https://docs.opengauss.org/en/docs/latest/extension_reference/extension_reference/server/gms_xmlparser.html)
- [Official source](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlparser/gms_xmlparser.control)
- [Official source](https://github.com/opengauss-mirror/openGauss-server/blob/41fe2237f5adb65405a1f76b893f527a457b1bec/contrib/gms_xmlparser/gms_xmlparser--1.0.sql)

`gms_xmlparser` 1.0 provides XML parsing and document handles for openGauss Oracle-compatibility workflows. It depends on libxml and is unavailable in lightweight builds. The SQL uses openGauss-specific types and is not a stock PostgreSQL compatibility claim.

### Enablement and Objects

```sql
CREATE EXTENSION gms_xmlparser;
```

The extension creates `gms_xmlparser` and `gms_xmldom` schemas. `gms_xmlparser.newparser` allocates a parser. `gms_xmlparser.parse`, `gms_xmlparser.parsebuffer` and `gms_xmlparser.parseclob` accept XML input; `gms_xmlparser.getdocument` returns the parsed document. `gms_xmlparser.freeparser` and `gms_xmldom.freedocument` release the respective resources.

### Resource and File Boundaries

Parser and document values are internal handles, not user-constructible scalar values. Allocate, use and release them within the same session workflow to avoid retaining session memory.

The single-string parsing interface tries XML first and interprets invalid XML as a local file path. Treat input validation and database-host file permissions as part of the security boundary; do not expose arbitrary input as a file-reading endpoint. Install with the privileges required by the openGauss deployment. No preload is documented. Remove dependent objects before dropping the extension.
