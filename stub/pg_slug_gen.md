
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_slug_gen;
> SELECT gen_random_slug();
> SELECT gen_random_slug(13);
> ```
>
> Source: [PGXN release README](https://pgxn.org/dist/pg_slug_gen/1.0.0/)

`pg_slug_gen` generates timestamp-based slugs using cryptographic randomness. The PGXN release README describes it as a PostgreSQL extension that maps timestamp digits into letter buckets and inserts a hyphen in the middle, producing URL-friendly slugs.

## Function

The documented SQL function is:

```sql
gen_random_slug(slug_length int DEFAULT 16) RETURNS text
```

The README shows these interfaces:

```sql
gen_random_slug()      -- default: 16 (microseconds)
gen_random_slug(10)    -- seconds
gen_random_slug(13)    -- milliseconds
gen_random_slug(16)    -- microseconds
gen_random_slug(19)    -- nanoseconds
```

## Precision and Length

The release README maps precision to timestamp digits and maximum collision-free throughput:

- `10` digits for seconds, up to 1 insert per second
- `13` digits for milliseconds, up to 1,000 inserts per second
- `16` digits for microseconds, up to 1,000,000 inserts per second
- `19` digits for nanoseconds, up to 1 billion inserts per second

The slug includes a midpoint hyphen:

- seconds: `5-5` pattern, 11 characters total
- milliseconds: `6-7` pattern, 14 characters
- microseconds: `8-8` pattern, 17 characters
- nanoseconds: `9-10` pattern, 20 characters

## Examples

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(13);
SELECT gen_random_slug(16);
SELECT gen_random_slug(19);
```

As a table default:

```sql
CREATE TABLE products (
    id serial PRIMARY KEY,
    name text NOT NULL,
    slug text DEFAULT gen_random_slug() UNIQUE
);
```

## How It Works

The release README describes the algorithm as:

1. take the current timestamp at the chosen precision
2. map each digit to a QWERTY-based character bucket
3. choose one random character from that bucket using `pg_strong_random()`
4. insert a hyphen at the midpoint

The README also notes that same-timestamp collisions remain possible, but with microsecond precision the probability is stated as roughly 1 in 10 million.
