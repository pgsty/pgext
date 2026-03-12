

## Usage

> [pg_xenophile: internationalization (i18n) and localization (l10n) utilities](https://github.com/bigsmoke/pg_xenophile)

The `pg_xenophile` extension provides i18n/l10n infrastructure including reference data for countries, languages, and currencies, plus automated localization table management.

```sql
CREATE EXTENSION pg_xenophile CASCADE;
```

All objects reside in the `xeno` schema (non-relocatable).

### Reference Tables

- **`xeno.country`**: ISO 3166-1 codes with calling codes and currencies
- **`xeno.lang`**: ISO 639-1 language codes
- **`xeno.currency`**: ISO 4217 currency codes with symbols
- **`xeno.country_subdivision`**: ISO 3166-2 subdivision codes
- **`xeno.eu_country`**: EU membership tracking
- **`xeno.country_postal_code_pattern`**: Postal code validation patterns

### Localization Tables

The extension auto-manages translation tables. Insert into `xeno.l10n_table` to register a translatable base table:

```sql
INSERT INTO xeno.l10n_table (base_table_schema, base_table_name)
VALUES ('public', 'products');
```

This automatically creates a companion `products_l10n` table and language-specific views.

### Convenience Views

- `xeno.country_l10n_en`: Country names in English
- `xeno.lang_l10n_en`: Language names in English
- `xeno.country_subdivision_l10n_en`: Subdivision names in English

### Configuration

```sql
SET pg_xenophile.base_lang_code = 'en';
SET pg_xenophile.user_lang_code = 'en';
SET pg_xenophile.target_lang_codes = '{nl,fr,de}';
```
