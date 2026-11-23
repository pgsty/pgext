## Usage

Sources:

- [README.md](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/README.md)
- [pg_wget.control](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget.control)
- [pg_wget--1.0.sql](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget--1.0.sql)
- [pg_wget.c](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/pg_wget.c)
- [License](https://github.com/mkgrgis/pg_wget/blob/a3b9ff378de69583e0adc225f63b5e3bad570c8c/License)

`pg_wget` retrieves an HTTP or HTTPS response as PostgreSQL text by starting the server's wget program. The extension's SQL version is 1.0. It is intended for small text downloads on Linux or a POSIX system; the reviewed revision does not establish a tested PostgreSQL major-version matrix.

### Retrieve Text

The server needs the external wget executable on its process search path. A superuser installs the extension; it is relocatable and requires no shared preload or restart.

```sql
CREATE EXTENSION pg_wget;
SELECT pg_wget('https://www.postgresql.org/');
```

The sole function, `pg_wget(text)`, returns text, is VOLATILE and STRICT, and returns NULL for a NULL argument. Only URLs beginning with `http://` or `https://` are accepted. A nonzero wget exit causes a database error. The implementation executes wget directly with a separate argument for the URL, without a shell.

### Access and Resource Limits

By PostgreSQL's normal function defaults, execution is initially available to PUBLIC. Restrict it to an existing trusted role before allowing ordinary users to connect:

```sql
REVOKE EXECUTE ON FUNCTION pg_wget(text) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_wget(text) TO downloader;
```

Replace the example role with the intended existing database role. Requests originate from the database server and can reach addresses available to its operating-system account. The extension checks the URL scheme but supplies no destination allowlist. Network access and the external executable must be controlled outside the extension.

The implementation collects the entire response in backend memory. It passes a ten-second timeout option to wget; that is not a guaranteed total statement deadline or a response-size cap. Use small text responses, avoid untrusted destinations, and do not treat this function as a binary or streaming HTTP client. Its external network effects are not undone by a transaction rollback.
