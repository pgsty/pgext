


## Usage

> [pg_pwhash: Advanced password hashing for PostgreSQL](https://github.com/cybertec-postgresql/pg_pwhash)

`pg_pwhash` provides modern adaptive password hashing algorithms including Argon2, scrypt, and yescrypt for PostgreSQL.

```sql
CREATE EXTENSION pg_pwhash;
```

### Supported Algorithms

| Identifier | Algorithm | Salt Pattern |
|------------|-----------|--------------|
| `argon2i` | Argon2i | `$argon2i$v=19$m=4096,t=3,p=1$<salt>` |
| `argon2d` | Argon2d | `$argon2d$v=19$m=4096,t=3,p=1$<salt>` |
| `argon2id` | Argon2id | `$argon2id$v=19$m=4096,t=3,p=1$<salt>` |
| `scrypt` | Scrypt | `$scrypt$ln=16,r=8,p=1$<salt>` |
| `$7$` | Scrypt (crypt) | `$7$BU<salt>` |
| `yescrypt` | yescrypt (crypt) | `$y$j9T$<salt>` |

### Core Functions

#### Generate Salt and Hash

```sql
-- Argon2id (recommended)
SELECT pwhash_crypt('password', pwhash_gen_salt('argon2id'));
-- $argon2id$v=19$m=4096,t=3,p=1$<salt>$<hash>

-- Scrypt
SELECT pwhash_crypt('password', pwhash_gen_salt('scrypt'));

-- Yescrypt
SELECT pwhash_crypt('password', pwhash_gen_salt('yescrypt'));
```

#### Verify Password

```sql
-- Hash matches if output equals stored hash
SELECT stored_hash = pwhash_crypt('entered_password', stored_hash) AS valid;
```

#### Direct Hashing Functions

```sql
SELECT pwhash_argon2('password', pwhash_gen_salt('argon2id'));
SELECT pwhash_scrypt('password', pwhash_gen_salt('scrypt'));
SELECT pwhash_yescrypt_crypt('password', pwhash_gen_salt('yescrypt'));
```

### Custom Salt Parameters

```sql
-- Argon2 with custom memory/time/parallelism
SELECT pwhash_gen_salt('argon2id', 'm=65536', 't=4', 'p=2');

-- Scrypt with custom parameters
SELECT pwhash_gen_salt('scrypt', 'ln=20', 'r=8', 'p=1');
```

### Configuration

| Parameter | Description |
|-----------|-------------|
| `pg_pwhash.argon2_default_backend` | Backend for Argon2: `libargon2` or `openssl` |
