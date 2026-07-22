## Usage

Sources:

- [Official extension SQL](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring--1.0.0.sql)
- [Official C implementation](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.c)
- [Official extension control file](https://github.com/tvondra/randomstring/blob/ae5259706484d66eb4dac63aaaea6ba745f7ffcd/randomstring.control)

`randomstring` version `1.0.0` provides simple pseudo-random text and byte generators for test data. It uses the C library `random()` generator, so neither function is suitable for passwords, tokens, keys, salts, or other security-sensitive material.

### Core Workflow

Both functions require a strictly positive output length:

```sql
CREATE EXTENSION randomstring;

SELECT random_string(24);
SELECT encode(random_bytea(16), 'hex');
```

`random_string(integer)` returns text, and `random_bytea(integer)` returns binary data. Zero and negative lengths raise an error.

### Distribution and Safety Caveats

The reviewed implementation indexes only the first 62 bytes of its declared text alphabet. Actual text output can contain a space, lowercase and uppercase letters, and digits 0 through 8; digit 9 and the listed punctuation are unreachable. Binary generation uses a modulus of 255, so byte value 255 is never emitted.

There is no extension-level seeding or uniqueness guarantee. Values can repeat and their distribution does not cover the advertised alphabets fully. Use these functions only for non-sensitive fixtures where those biases are acceptable; use PostgreSQL cryptographic facilities or an application CSPRNG for security material.
