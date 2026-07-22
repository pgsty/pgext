## Usage

Sources:

- [PGXN README for distribution 1.0.0](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/README.md)
- [Version 0.0.1 installation SQL](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/sql/pg_datatype_password--0.0.1.sql)
- [Extension control file](https://api.pgxn.org/src/pg_datatype_password/pg_datatype_password-1.0.0/pg_datatype_password.control)

`pg_datatype_password` provides a `password` type and plaintext comparison operators backed by `pgcrypto` Blowfish hashes. The reviewed PGXN distribution is `1.0.0`, but its installable extension version is `0.0.1`.

### Core Workflow

The type input function does not hash values. Every table that stores the type must use `t_encrypt_password` as a BEFORE trigger, or plaintext will be stored.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_datatype_password;

CREATE TABLE app_user (
  username text PRIMARY KEY,
  passwd password NOT NULL
);

CREATE TRIGGER app_user_encrypt_password
BEFORE INSERT OR UPDATE OF passwd ON app_user
FOR EACH ROW EXECUTE PROCEDURE t_encrypt_password('passwd', '10');

INSERT INTO app_user VALUES ('alice', 'secret');
SELECT username FROM app_user WHERE passwd = 'secret'::text;
```

The optional second trigger argument is the bcrypt cost; upstream defaults to `8`. Output displays the stored hash. Updating the password column hashes the new value, so assigning an existing hash will hash it again.

### Security and Caveats

Available comparisons are between `password` and `text` in either direction, using `=` and `<>`; there is no documented ordering or password index operator class. Restrict direct table access, avoid logging plaintext, and use current password policies and rate limiting around authentication.

The versioned SQL hard-codes several objects in `public` even though the control file says relocatable. Install and test in the expected schema. This is an old bcrypt-based design with an outdated default work factor; evaluate whether application-side password hashing with a maintained library better fits the threat model.
