## Usage

Sources:

- [Official README](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/README.md)
- [Extension control file](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.control)
- [Extension SQL](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail--0.0.2.sql)
- [Mail implementation](https://github.com/sshutdownow/pg_sendmail/blob/f8c355d3ffcfc31302694d85c0740b2017c35f3f/pg_sendmail.c)

`pg_sendmail` sends a message by piping generated mail text to `/usr/sbin/sendmail` from a PostgreSQL backend. The pinned 0.0.2 implementation has a command-injection vulnerability in address handling, so it must not be exposed to untrusted input or users and should not be deployed without a code fix.

### Installed Surface

If auditing an existing installation, contain it before any functional test:

```sql
CREATE EXTENSION pg_sendmail;

REVOKE ALL ON FUNCTION mail(text, text, text, text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION sendmail(text, text, text, text) FROM PUBLIC;
```

`mail(mailfrom,rcptto,subject,msg_body,headers)` is the C entry point. `sendmail(mailfrom,rcptto,subject,msg_body)` is a PL/pgSQL wrapper that encodes the subject and supplies MIME headers. Both return a Boolean from the local sendmail invocation; that value does not prove final delivery.

### Critical Security and Transaction Caveats

The C code constructs a shell command by inserting `mailfrom` and `rcptto` directly into a string passed to `popen`. These values are neither quoted nor safely passed as an argument vector, so crafted input can execute operating-system commands as the PostgreSQL service account. Headers and subject data can also alter message structure. Revoking `PUBLIC` is only containment, not a complete fix; replace shell execution with a non-shell API and strict address/header validation before enabling the feature.

Sending mail is synchronous, can block a database backend, and is an external side effect that PostgreSQL rollback cannot undo. Queue application notifications transactionally and let a separate least-privilege worker deliver them instead. The SQL script also changes ownership of `sendmail(text,text,text,text)` to a hard-coded `postgres` role, which can make installation fail or create an unintended owner. Review package code, role ownership, executable path, MTA policy, timeouts, audit logging, and failure handling on the exact host.
