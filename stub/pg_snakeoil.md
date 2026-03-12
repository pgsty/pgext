


## Usage

> [pg_snakeoil: ClamAV antivirus scanning for PostgreSQL data](https://github.com/credativ/pg_snakeoil)

`pg_snakeoil` provides ClamAV virus scanning of data stored in PostgreSQL without interfering with normal database operations.

```sql
CREATE EXTENSION pg_snakeoil;
```

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `so_is_infected(text)` | `bool` | Check if text data matches a virus signature |
| `so_is_infected(bytea)` | `bool` | Check if bytea data matches a virus signature |
| `so_virus_name(text)` | `text` | Return virus name if infected, empty string otherwise |
| `so_virus_name(bytea)` | `text` | Return virus name if infected, NULL otherwise |
| `so_update_signatures()` | `bool` | Reload virus signatures, true if changed |

### Ad-hoc Scanning

```sql
SELECT so_is_infected('Not a virus!');
-- f

SELECT so_is_infected('X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*');
-- t

SELECT so_virus_name('X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*');
-- Eicar-Test-Signature
```

### On-Access Protection with Domains

```sql
CREATE DOMAIN safe_text AS text CHECK (NOT so_is_infected(value));
CREATE TABLE t1 (safe safe_text);

INSERT INTO t1 VALUES ('This text is safe!');
-- INSERT

INSERT INTO t1 VALUES('X5O!P%@AP...');
-- NOTICE: Virus found: Eicar-Test-Signature
-- ERROR: value for domain safe_text violates check constraint "safe_text_check"
```

### On-Access Protection with Triggers

```sql
CREATE OR REPLACE FUNCTION check_virus() RETURNS trigger AS $$
BEGIN
    IF so_is_infected(NEW.content) THEN
        RAISE EXCEPTION 'Virus detected: %', so_virus_name(NEW.content);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER virus_check BEFORE INSERT OR UPDATE ON uploads
    FOR EACH ROW EXECUTE FUNCTION check_virus();
```
