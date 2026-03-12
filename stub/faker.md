


## Usage

> [faker: Wrapper for the Faker Python library](https://github.com/anpandu/postgresql_faker)

`faker` is a PostgreSQL extension that wraps Python's Faker library, providing functions to generate realistic fake data directly in SQL queries. It requires `plpython3u`.

```sql
CREATE EXTENSION faker;
```

### Generate Fake Data

```sql
SELECT faker.name();           -- 'John Smith'
SELECT faker.first_name();     -- 'Jane'
SELECT faker.last_name();      -- 'Doe'
SELECT faker.email();          -- 'jane.doe@example.com'
SELECT faker.address();        -- '123 Main St, Anytown, US 12345'
SELECT faker.company();        -- 'Smith LLC'
SELECT faker.phone_number();   -- '(555) 123-4567'
SELECT faker.text();           -- random paragraph of text
SELECT faker.city();           -- 'Portland'
SELECT faker.country();        -- 'United States'
```

Note: `faker.date()` and `faker.time()` are **not available** because `date` and `time` are reserved PostgreSQL keywords. Use `faker.date_between()` or `faker.date_this_century()` instead.

### Populate Tables with Fake Data

```sql
INSERT INTO users (name, email, address, created_at)
SELECT
  faker.name(),
  faker.email(),
  faker.address(),
  faker.date_this_century()::timestamp
FROM generate_series(1, 1000);
```

### Localized Fake Data

Locale is set per session, not per function call:

```sql
SELECT faker.faker('de_DE');   -- set locale for this session
SELECT faker.name();           -- returns a German name
SELECT faker.address();        -- returns a German address
```

### Unique Values

Use the `unique_` prefix to guarantee unique values within a session:

```sql
SELECT faker.unique_name();
SELECT faker.unique_email();
```

### Discover All Functions

```sql
SELECT faker._functions();     -- list all 500+ available functions
```

All faker functions return `text`. Cast explicitly for other types.

### Common Faker Providers

| Function | Description |
|----------|-------------|
| `faker.name()` | Full name |
| `faker.first_name()` | First name |
| `faker.last_name()` | Last name |
| `faker.email()` | Email address |
| `faker.company_email()` | Company email |
| `faker.phone_number()` | Phone number |
| `faker.address()` | Full address |
| `faker.city()` | City name |
| `faker.country()` | Country name |
| `faker.company()` | Company name |
| `faker.text()` | Random text |
| `faker.date_this_century()` | Random date |
| `faker.ssn()` | Social security number |
| `faker.ean()` | EAN barcode |
