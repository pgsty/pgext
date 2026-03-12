

## Usage

> [Address Standardizer: Address parsing and standardization for PostGIS](https://github.com/postgis/postgis)

The Address Standardizer is a PostGIS extension that parses a single-line address string into a structured form using configurable lexicon, grammar, and rules tables. It is a more flexible alternative to the built-in `normalize_address` function in the TIGER geocoder.

- [Address Standardizer Reference](https://postgis.net/docs/Extras.html#Address_Standardizer)

### Setup

```sql
CREATE EXTENSION address_standardizer;
```

--------

## Standardizing Addresses

The core function takes an address string and three table references (lex, gaz, rules):

```sql
SELECT *
FROM standardize_address(
    'us_lex',        -- lexicon table
    'us_gaz',        -- gazetteer table
    'us_rules',      -- rules table
    '1600 Pennsylvania Ave NW, Washington, DC 20500'
);
```

The result contains structured fields:

| Field | Description |
|-------|-------------|
| `building` | Building name or identifier |
| `house_num` | Street number |
| `predir` | Prefix direction (N, S, E, W) |
| `qual` | Qualifier |
| `pretype` | Prefix type |
| `name` | Street name |
| `suftype` | Suffix type (St, Ave, Blvd) |
| `sufdir` | Suffix direction |
| `ruralroute` | Rural route |
| `extra` | Extra information |
| `city` | City name |
| `state` | State |
| `country` | Country |
| `postcode` | ZIP/postal code |
| `box` | PO Box |
| `unit` | Unit/apartment number |

--------

## Lexicon, Gazetteer, and Rules Tables

The standardizer is driven by three user-configurable tables:

**Lexicon (lex)** -- Maps input tokens to standardized forms and token classes:

```sql
CREATE TABLE us_lex (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**Gazetteer (gaz)** -- Maps place names (cities, states) to standard forms:

```sql
CREATE TABLE us_gaz (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**Rules (rules)** -- Defines grammar rules for parsing addresses:

```sql
CREATE TABLE us_rules (
    id serial PRIMARY KEY,
    rule text
);
```

For US addresses, the `address_standardizer_data_us` extension provides pre-built data for these tables.
