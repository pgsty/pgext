## Usage

Sources:

- [Official README](https://github.com/asio/password/blob/master/README)
- [Extension SQL](https://github.com/asio/password/blob/master/sql/password--0.0.1.sql)
- [C implementation](https://github.com/asio/password/blob/master/src/password.c)

`password` 0.0.1 is an obsolete prototype for a custom password type backed by bundled Blowfish crypt code. Source review found multiple memory-safety and type-handling defects in the hashing, typmod, and comparison paths. Do not deploy it for authentication or store real credentials with it.

### Intended SQL Surface

The following illustrates the upstream interface for audit and compatibility analysis only, not a production recommendation:

```sql
CREATE EXTENSION password;

CREATE TABLE password_demo (
  account text PRIMARY KEY,
  secret password
);
```

The type intends to hash clear input with a `$2a$` Blowfish setting, accepts strings already beginning with `$2a$` as stored hashes, and exposes an implicit cast back to text. Its `=` operators treat the left operand as a stored password and the right operand as entered clear text; this is not normal symmetric equality.

### Security and Correctness Caveats

The C source returns a pointer to stack storage from its hashing helper, misuses integer pointers in typmod handling, copies only pointer-sized bytes in one comparison path, and does not consistently preserve or check cryptographic error results. These defects can cause crashes, incorrect comparisons, or undefined behavior. The comparison is not documented as constant-time, and the bundled implementation targets the old `$2a$` variant.

### Safer Boundary

Keep password verification in a maintained authentication component, or use a currently supported and independently reviewed PostgreSQL mechanism such as `pgcrypto` only under an explicit security design. Existing databases containing this custom type should be treated as a migration and incident-review problem: export hashes only after controlled validation, revoke access, and retire the extension rather than trusting its comparisons.
