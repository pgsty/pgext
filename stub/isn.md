

## Usage

> [isn: ISBN, ISSN, EAN, UPC product number types](https://www.postgresql.org/docs/current/isn.html)

The `isn` extension provides data types for international product numbering standards with validation and check-digit verification.

```sql
CREATE EXTENSION isn;
```

### Data Types

| Type | Description |
|------|-------------|
| `EAN13` | European Article Numbers (always 13-digit) |
| `ISBN13` | International Standard Book Numbers (13-digit) |
| `ISMN13` | International Standard Music Numbers (13-digit) |
| `ISSN13` | International Standard Serial Numbers (13-digit) |
| `ISBN` | Book numbers (10-digit short format when possible) |
| `ISMN` | Music numbers (10-digit short format when possible) |
| `ISSN` | Serial numbers (10-digit short format when possible) |
| `UPC` | Universal Product Codes |

### Examples

```sql
SELECT isbn('978-0-393-04002-9');
SELECT isbn13('0901690546');
SELECT issn('1436-4522');

-- Create table with ISBN column
CREATE TABLE books (id isbn);
INSERT INTO books VALUES ('9780393040029');

-- Auto-calculate check digit with ?
INSERT INTO books VALUES ('220500896?');

-- Type casting
SELECT ean13('0901690546'::isbn);
```

### Weak Mode

Enable weak mode to accept invalid check digits (useful for OCR/data cleanup):

```sql
SET isn.weak TO true;
INSERT INTO books VALUES ('978-0-11-000533-4');
SET isn.weak TO false;

-- Find invalid entries
SELECT id FROM books WHERE NOT is_valid(id);

-- Fix invalid entries
UPDATE books SET id = make_valid(id) WHERE id = '2-205-00876-X!';
```

### Functions

| Function | Description |
|----------|-------------|
| `is_valid(isn)` | Check if value has valid check digit |
| `make_valid(isn)` | Clear the invalid-check-digit flag |
| `isn_weak()` | Return current weak mode status |
| `isn_weak(boolean)` | Set weak mode |

### Operators and Indexing

Standard comparison operators (`=`, `<`, `>`, `<=`, `>=`, `<>`) with B-tree and hash index support.
