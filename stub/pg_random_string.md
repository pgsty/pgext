## Usage

Sources:

- [Upstream README](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/README.md)
- [Version 0.0.1 install SQL](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/pg_random_string--0.0.1.sql)
- [C implementation](https://github.com/meiyifei/pg_random_string/blob/d4ce8fb8c5835549503c16a3118ca71f31bcc0c0/pg_random_string.c)

pg_random_string version 0.0.1 exposes random_string(integer), which attempts to return an alphanumeric string of the requested length. The implementation is a small proof of concept and is unsafe as published.

### Demonstration only

Use the function only in a disposable, isolated test server after reviewing or repairing the C code:

```sql
CREATE EXTENSION pg_random_string;
SELECT random_string(16);
```

Do not use its output for passwords, tokens, salts, identifiers that require unpredictability, or any other security purpose.

### Caveats

- The function seeds the process-global rand generator from the current time on every call. Calls within the same second can repeat, and the generator is not cryptographically secure.
- The requested integer length is truncated to a short integer. Negative or large inputs are not validated.
- The C code allocates exactly the requested byte count but does not append a terminating zero before calling pstrdup. This can read beyond the allocation, return corrupted data, disclose adjacent memory, or crash a backend.
- The allocation uses malloc without a corresponding free and bypasses PostgreSQL memory contexts.
- Upstream publishes no tests, license, release history, or current PostgreSQL compatibility statement. Replace this implementation with maintained, reviewed SQL or cryptographic primitives rather than deploying it.
