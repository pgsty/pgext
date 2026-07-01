
## Usage

Sources: [README](https://github.com/pmpetit/pglinter/blob/2.0.0/README.md), [how-to](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/how-to/README.md), [examples](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/examples/README.md), [rules](https://github.com/pmpetit/pglinter/blob/2.0.0/docs/rules/README.md), [2.0.0 release](https://github.com/pmpetit/pglinter/releases/tag/2.0.0)

pglinter analyzes a PostgreSQL database for potential issues, performance problems, and best practice violations. Current user docs expose findings through `pglinter.get_violations()`, which returns enabled-rule violations as rows that can be filtered or joined to `pg_identify_object()`.

### Run Checks

```sql
SELECT * FROM pglinter.get_violations();
SELECT * FROM pglinter.get_violations() WHERE rule_code = 'B001';

SELECT
  rule_code,
  (pg_identify_object(classid, objid, objsubid)).type AS object_type,
  (pg_identify_object(classid, objid, objsubid)).schema AS object_schema,
  (pg_identify_object(classid, objid, objsubid)).name AS object_name,
  (pg_identify_object(classid, objid, objsubid)).identity AS object_identity
FROM pglinter.get_violations();
```

### Rule Management

```sql
SELECT pglinter.show_rules();                -- Show all rules and their status
SELECT pglinter.explain_rule('B001');        -- Get rule details and suggested fixes
SELECT pglinter.enable_rule('B001');         -- Enable a specific rule
SELECT pglinter.disable_rule('B001');        -- Disable a specific rule
SELECT pglinter.is_rule_enabled('B001');     -- Check if a rule is enabled
SELECT pglinter.enable_all_rules();
SELECT pglinter.disable_all_rules();
SELECT pglinter.show_rule_queries('B001');   -- Inspect the rule query
SELECT pglinter.list_rules();                -- Return a formatted rule list
```

### Rule Import And Export

```sql
SELECT pglinter.export_rules_to_yaml();                -- Export rules to YAML
SELECT pglinter.import_rules_from_yaml('yaml...');     -- Import rules from YAML
SELECT pglinter.export_rules_to_file('/path/to/rules.yaml');
SELECT pglinter.import_rules_from_file('/path/to/rules.yaml');
SELECT pglinter.export_rulemessages_to_yaml();
SELECT pglinter.import_rule_messages_from_yaml('yaml...');
```

### Rule Families

**Base (B-series):** B001 tables without PK, B002 redundant indexes, B003 missing FK indexes, B004 unused indexes, B005 uppercase names, B006 unused tables, B007 cross-schema FKs, B008 FK type mismatches, B009 shared trigger functions, B010 reserved keywords, B011 multiple owners per schema, B012 composite primary keys with more than four columns, B013 row-by-row trigger processing without a `WHERE` clause.

**Cluster (C-series):** C002 insecure pg_hba.conf entries, C003 MD5 password encryption.

**Schema (S-series):** S001 no default role grants, S002 env prefixes/suffixes, S003 unsecured public schema, S004 system role ownership, S005 multiple owners per schema.

### Caveats

Version `2.0.0` removes the older `check()` and `check_rule()` functions; `get_violations()` is now the only check API. The current build uses `pgrx` 0.18.1.

The upstream `1.1.2` release added `B013`. The main README remains partially stale compared with the docs and exported functions, so this stub uses `get_violations()` and omits removed `check()`/`check_rule()` examples.
