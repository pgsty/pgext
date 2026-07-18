## Usage

Sources:

- [Upstream README](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/README)
- [Extension control file](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/pgsmtp.control)
- [PL/Python implementation](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/pgsmtp--0.1.1.sql)
- [Distribution metadata](https://github.com/asotolongo/pgsmtp/blob/babd72278eb2f904ec2942909e52062894852668/META.json)

`pgsmtp` sends email from PostgreSQL through Python's SMTP client. Version `0.1.1` creates a `pgsmtp.user_smtp_data` table for sender/server/password records and provides `pgsmtp.pg_smtp_mail()` plus the attachment-capable `pgsmtp.pg_smtp_mail_attach()`. It requires the untrusted `plpython3u` language and superuser installation.

### Configure a sender and send mail

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION pgsmtp;

INSERT INTO pgsmtp.user_smtp_data
  (smtp_user, smtp_server, smtp_port, smtp_pass)
VALUES
  ('sender@example.com', 'smtp.example.com', 465, 'replace-this-secret');

SELECT pgsmtp.pg_smtp_mail(
  'sender@example.com',
  'receiver@example.com',
  ARRAY[]::text[],
  'Status',
  'Job completed'
);
```

The implementation uses `SMTP_SSL`, stores passwords as plaintext, and performs network I/O inside the database backend. More seriously, it concatenates the caller-supplied sender into a PL/Python SQL query, creating SQL-injection risk. `pgsmtp.pg_smtp_mail_attach()` also reads arbitrary server-side paths as the PostgreSQL operating-system account.

Upstream warns that bugs remain and describes the extension as administrator-only. Do not grant its schema, credential table, or function execution to application roles. Prefer an external mail queue; if legacy use is unavoidable, parameterize the lookup, encrypt or externalize credentials, restrict egress and file access, add timeouts/auditing, and test SMTP failure behavior before deployment.
