

## Usage

> [Address Standardizer Data US: US address data for the address_standardizer extension](https://github.com/postgis/postgis)

This extension provides pre-built US lexicon, gazetteer, and rules data for use with the `address_standardizer` extension. It includes tables with common US street types, directional abbreviations, state names, and grammar rules needed to parse US addresses.

- [Address Standardizer Reference](https://postgis.net/docs/Extras.html#Address_Standardizer)

### Setup

```sql
CREATE EXTENSION address_standardizer_data_us;
```

This creates the `us_lex`, `us_gaz`, and `us_rules` tables in the public schema, pre-populated with US address data.

--------

## Using with address_standardizer

Once installed, you can immediately standardize US addresses:

```sql
SELECT *
FROM standardize_address(
    'us_lex', 'us_gaz', 'us_rules',
    '123 Main Street NW, Apt 4B, Springfield, IL 62704'
);
```

The provided data covers common US address patterns including:

- Street types (Street, Avenue, Boulevard, Drive, Lane, Court, etc.)
- Directional prefixes and suffixes (North, South, N, S, NW, SE, etc.)
- State names and abbreviations
- Unit designators (Apt, Suite, Unit, etc.)
- Highway designators (US, State, County, Interstate, etc.)
