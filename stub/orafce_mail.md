## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/README.md)
- [Official extension control file (orafce_mail.control)](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/orafce_mail.control)
- [Official extension SQL (orafce_mail--1.2.sql)](https://api.pgxn.org/src/orafce_mail/orafce_mail-1.2.3/orafce_mail--1.2.sql)

`orafce_mail` — An extension Orafce should be installed before. Use it when porting or emulating the corresponding database API. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION orafce_mail;

set orafce_mail.smtp_server_url to 'smtps://smtp.gmail.com:465';
set orafce_mail.smtp_server_userpwd to 'pavel.stehule@gmail.com:yourgoogleapppassword';

call utl_mail.send(sender => 'pavel.stehule@gmail.com',
                   recipients => 'pavel.stehule@gmail.com',
                   subject => 'ahoj, nazdar, žlutý kůň',
                   message => e'test, \nžlutý kůň');

do $$
declare
  myimage bytea = (select img from foo limit 1);
begin
  call utl_mail.send_attach_raw(sender => 'pavel.stehule@gmail.com',
                                recipients => 'pavel.stehule@gmail.com',
                                subject => 'mail with picture',
                                message => 'I am sending some picture',
                                attachment => myimage,
                                att_mime_type => 'image/png',
                                att_filename => 'screenshot.png');
end
$$;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dbms_mail.send` is an extension procedure.
- `utl_mail.send` is an extension procedure.
- `utl_mail.send_attach_raw` is an extension procedure.
- `utl_mail.send_attach_varchar2` is an extension procedure.
- `dbms_mail` is a schema created by the extension.
- `utl_mail` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- Install the confirmed extension dependencies first: `orafce`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
