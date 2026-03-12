

## Usage

> [pg_utl_smtp: Oracle UTL_SMTP compatibility extension for PostgreSQL](https://github.com/hexacluster/pg_utl_smtp)

### Enabling

```sql
CREATE EXTENSION plperlu;
CREATE EXTENSION pg_utl_smtp;
```

### Sending an Email

```sql
DO $$
DECLARE
    c utl_smtp.connection;
BEGIN
    c := utl_smtp.open_connection('smtp.example.com', 25);
    CALL utl_smtp.ehlo(c, 'mydomain.com');
    CALL utl_smtp.mail(c, 'sender@example.com');
    CALL utl_smtp.rcpt(c, 'recipient@example.com');
    CALL utl_smtp.open_data(c);
    CALL utl_smtp.write_data(c, 'From: sender@example.com' || E'\r\n');
    CALL utl_smtp.write_data(c, 'To: recipient@example.com' || E'\r\n');
    CALL utl_smtp.write_data(c, 'Subject: Test Email' || E'\r\n');
    CALL utl_smtp.write_data(c, E'\r\n');
    CALL utl_smtp.write_data(c, 'Hello from PostgreSQL!');
    CALL utl_smtp.close_data(c);
    CALL utl_smtp.quit(c);
END;
$$;
```

### Procedures

- **OPEN_CONNECTION(host, port, tx_timeout, ...)** - Opens a connection to an SMTP server. Returns a `utl_smtp.connection` type. Supports SSL/TLS via `secure_connection_before_smtp`.
- **EHLO(c, domain)** / **HELO(c, domain)** - Performs initial SMTP handshake.
- **MAIL(c, sender)** - Initiates a mail transaction.
- **RCPT(c, recipient)** - Specifies e-mail recipient. Call multiple times for multiple recipients.
- **OPEN_DATA(c)** - Sends the DATA command to begin message body.
- **WRITE_DATA(c, data)** - Writes a portion of the message body.
- **WRITE_RAW_DATA(c, data)** - Writes raw data to the message body.
- **CLOSE_DATA(c)** - Closes the data session.
- **QUIT(c)** - Terminates the SMTP session and disconnects.

### Connection Type

```sql
-- utl_smtp.connection composite type
(host varchar(255), port integer, tx_timeout integer,
 private_tcp_con integer, private_state integer)
```

### Notes

- Requires the Perl `Net::SMTP` module installed on the system
- Use `E'\r\n'` for line breaks instead of `utl_tcp.crlf`
- The `wallet_path` and `wallet_password` parameters are not used
