## Usage

Sources:

- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/README.md)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl.control)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.5.sql)
- [Official source](https://api.pgxn.org/src/pg_curl/pg_curl-2.5.0/pg_curl--2.4--2.5.sql)
- [Official source](https://api.pgxn.org/dist/pg_curl/2.5.0/META.json)

`pg_curl` exposes libcurl transfers as SQL functions. Upstream release 2.5.0 uses SQL extension version 2.5. Requests run in the database backend with the server's network and filesystem access; installing the extension requires an administrator. Ordinary calls need no preload.

### A Request in One Transaction

```sql
CREATE EXTENSION pg_curl;
BEGIN;
SELECT curl_easy_reset();
SELECT curl_easy_setopt_url('https://www.postgresql.org/');
SELECT curl_easy_perform();
SELECT convert_from(curl_easy_getinfo_data_in(), 'UTF8');
COMMIT;
```

`curl_easy_reset` clears the handle, `curl_easy_setopt_url` chooses the destination, and `curl_easy_perform` performs the transfer. `curl_easy_getinfo_data_in` returns response bytes and `curl_easy_getinfo_header_in` returns headers. `pg_curl.transaction` defaults to true: handles are released at transaction end, so keep a multi-statement request in one transaction. Network/file side effects are not rolled back with SQL.

### Access and Whitelists

A superuser grants allowed prefixes to an existing application role:

```sql
ALTER ROLE app_user SET pg_curl.whitelist = 'https://api.example.com/';
```

Reconnect to apply a role default. With an unset `pg_curl.whitelist`, superusers are unrestricted and other roles are denied. A configured whitelist also restricts superusers. Only matching HTTP, HTTPS and file URLs are allowed while it applies; FTP, SMTP and other protocols are denied. HTTP paths match at segment boundaries, not arbitrary string prefixes.

Local-file options are checked too. File paths are resolved and must already exist, including an output cookie-jar file. Redirect targets are checked; restricted redirect following requires libcurl 7.80.0 or later. Explicit proxy and DNS-over-HTTPS options are denied while restricted, although the server's proxy environment can still apply. Limit allowed hosts and paths deliberately; stored credentials and server-local files remain sensitive.

### Other Operations and Upgrade

`curl_postfield_append`, `curl_easy_setopt_postfields` and `curl_header_append` construct form or JSON requests. `curl_mime_data` and `curl_mime_file` construct MIME parts; 2.5 adds a `head text[]` overload for multiple headers. Header array elements cannot be NULL or contain carriage returns/newlines. The MIME data overload accepts text or binary data.

Check `pg_extension.extversion` before upgrading. This distribution includes the 2.4-to-2.5 update script; it does not supply a 2.4.1-to-2.5 edge, so do not assume an existing installation reporting 2.4.1 can use an automatic update. For an installation actually at SQL 2.4, after installing matching files:

```sql
ALTER EXTENSION pg_curl UPDATE TO '2.5';
```

Use the new whitelist rules before granting execution to application roles. Review libcurl's supported protocols and use the matching library on all database hosts.
