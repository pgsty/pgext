## Usage

Sources:

- [Official README](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/README.md)
- [Extension control file](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer.control)
- [Extension SQL implementation](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer--0.0.1.sql)

gdpr_anatomizer is a research prototype implementing the Anatomy privacy-publication algorithm with PL/Python. It separates quasi-identifiers from a sensitive attribute into linked tables. Upstream explicitly describes it as inefficient for large data and intended for learning and debugging; it does not itself establish GDPR compliance or a privacy guarantee.

### Core Workflow

The extension requires the untrusted `plpython3u` language. Inspect its installation SQL before use: installation calls initialization functions that drop and recreate unqualified example and similarity tables. Run it only in a disposable database, then analyze a copied dataset and generate anatomy tables:

```sql
CREATE EXTENSION gdpr_anatomizer CASCADE;

SELECT analyze_dataset('public.input_copy', 'sensitive_value');

SELECT anatomy(
  'public.input_copy',
  'sensitive_value',
  ARRAY['age_band', 'region'],
  3
);
```

The default output names are `qi_table` and `sa_table`. Verify the generated equivalence classes and whether suppression meets the application's threat model before exposing any result.

### Important Objects

- `analyze_dataset` reports the sensitive-value distribution and candidate diversity levels.
- `anatomy` constructs the quasi-identifier and sensitive-attribute tables; arguments also control schema, output names, and reference columns.
- `apply_sa_tag` and `remove_sa_tags` manage tags used for sensitive attributes.
- `set_del_eq_class_trigger` installs deletion handling for equivalence classes.
- `init_csv_dataset` performs a server-side file import and recreates a demonstration table.

### Security and Data-Safety Notes

PL/Python is untrusted and its code runs with database-server process capabilities. This prototype loads much of the dataset into memory, uses randomized grouping, and builds dynamic SQL by string formatting, so untrusted identifiers or values can cause SQL injection. Several helpers drop or create fixed, unqualified tables, and `init_csv_dataset` reads a server-local path. Use a restricted disposable database, inspect every generated table, and never treat the output as independently validated anonymization.
