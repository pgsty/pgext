## Usage

Sources:

- [Official README](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/README.md)
- [Usage documentation](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/doc/variant.md)
- [Extension SQL](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/sql/variant.sql)
- [C implementation](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/src/variant.c)
- [Extension control file](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/variant.control)

`variant` 1.0.1 defines `variant.variant`, a container that stores a PostgreSQL value in its native representation together with the original type and type modifier. It can support heterogeneous settings or attribute values, but it couples stored rows to the underlying type OIDs, binary representation, I/O functions, and dynamically generated casts. Prefer explicit typed columns or JSON for portable application data.

### Register and Use a Variant

Upstream discourages the disabled default variant because some PostgreSQL paths do not preserve type modifiers. Register a named modifier and use it on the column:

```sql
CREATE EXTENSION variant;

SELECT variant.register('setting');

CREATE TABLE app_setting (
    setting_name text PRIMARY KEY,
    setting_value variant.variant(setting) NOT NULL
);

INSERT INTO app_setting VALUES
    ('retry_count', 3::integer),
    ('label', 'primary'::text),
    ('window', '((0,0),(1,1))'::box);

SELECT setting_name,
       variant.original_type(setting_value),
       variant.text_out(setting_value)
FROM app_setting;
```

Ordinary casts are the primary input/output interface. `variant.text_in(text, text)` parses the explicit textual representation for a named registered modifier. `variant.original_type(variant.variant)` identifies the stored base type, while `variant.text_out(variant.variant)` returns its text form.

### Casts, Comparisons, and Indexing

Installation calls `variant.create_casts()` to generate functions and casts for existing non-composite, non-domain, non-pseudo types. Run it again after installing a new supported data type, and inspect `variant.missing_casts` and `variant.variant_casts`. Dynamic catalog-wide cast creation is privileged DDL; restrict the function to extension administrators.

The extension provides comparison operators across variants and a hash operator class based on the separate image-equality operator `*=`. Test semantic equality, binary image equality, ordering across different stored types, collation, NaN values, and hash-index behavior for every intended type rather than assuming the underlying type's normal operator class is preserved.

### Type Lifecycle and Portability

Variant stores native binary data, not a self-describing portable encoding. Upstream warns that dropping a type used by variant rows can make those values unrecoverable; the extension does not create a dependency for each stored row or reliably restrict registered modifiers to approved types. Inventory stored base types before dropping or upgrading any extension, domain, enum, or user-defined type.

Binary formats and type OIDs can change across PostgreSQL major versions, extension upgrades, architectures, or restore order. Prove logical dump/restore into an empty cluster and conversion back to each original type before adopting it. Do not use raw filesystem copies or logical replication into a catalog with different type identities. The reviewed project is old preview-era native code with no current compatibility matrix; pin the exact build, bound value sizes, test TOAST and malformed text input, and keep a migration path to explicit columns or a stable serialized format.
