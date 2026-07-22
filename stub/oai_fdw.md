## Usage

Sources:

- [Official v1.13 README](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/README.md)
- [Official v1.13 control file](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/oai_fdw.control)
- [Official v1.13 changelog](https://github.com/jimjonesbr/oai_fdw/blob/v1.13/CHANGELOG.md)
- [Official OAI-PMH specification](https://www.openarchives.org/OAI/openarchivesprotocol.html)

`oai_fdw` 1.13 is a foreign data wrapper for harvesting metadata from OAI-PMH 2.0 repositories over HTTP or HTTPS. It maps OAI record headers and XML content to foreign-table columns, supports automatic table generation for repositories or sets, and can copy harvested records into local tables.

### Core Workflow

Create a server for the repository, then either define a foreign table explicitly or use `IMPORT FOREIGN SCHEMA`. Every table needs a metadata format such as `oai_dc`.

```sql
CREATE EXTENSION oai_fdw;
CREATE SCHEMA oai_data;

CREATE SERVER museum_oai
FOREIGN DATA WRAPPER oai_fdw
OPTIONS (url 'https://example.org/oai');

CREATE FOREIGN TABLE oai_data.records (
    id text OPTIONS (oai_node 'identifier'),
    content xml OPTIONS (oai_node 'content'),
    sets text[] OPTIONS (oai_node 'setspec'),
    datestamp timestamp OPTIONS (oai_node 'datestamp'),
    format text OPTIONS (oai_node 'metadataprefix'),
    status boolean OPTIONS (oai_node 'status')
)
SERVER museum_oai
OPTIONS (metadataprefix 'oai_dc');

SELECT id, datestamp
FROM oai_data.records
LIMIT 100;
```

Without `setspec`, a scan can harvest every set exposed by the repository. Use `from`, `until`, or `setspec` table options and an outer `LIMIT` where appropriate, but verify request behavior with the remote service.

### Import and Support Interfaces

- `IMPORT FOREIGN SCHEMA oai_repository` creates one table spanning the repository; `oai_sets` creates a table per OAI set. Both require the `metadataprefix` import option and support `LIMIT TO` or `EXCEPT`.
- Column option `oai_node` maps `identifier`, `setspec`, `datestamp`, `content`, `metadataprefix`, or `status` to compatible PostgreSQL types.
- Server options include `url`, `http_proxy`, `connect_timeout`, `connect_retry`, `request_redirect`, and `request_max_redirect`.
- `OAI_HarvestTable` copies a foreign-table harvest into a local table; OAI support functions expose repository metadata such as sets and formats.

Version 1.13 moves `proxy_user` and `proxy_password` to the user mapping. Repository `user` and `password` also belong there and use HTTP Basic Authentication.

### Network, Credentials, and Consistency

Queries perform synchronous remote requests and parse untrusted XML. Latency, pagination, server throttling, redirects, malformed responses, and repository changes affect results. Store credentials in narrowly scoped user mappings, restrict who may create or alter servers, and treat configurable URLs/proxies as outbound-network access with SSRF implications.

OAI-PMH is a harvesting protocol, not a transactional database API. Repeated scans can observe changing pages and deletions; local copies become stale. Record harvest windows and identifiers, make refreshes restartable, and reconcile deleted/status records explicitly. Version 1.13 requires matching libcurl/libxml2 build dependencies and should be tested against each target repository’s metadata format.
