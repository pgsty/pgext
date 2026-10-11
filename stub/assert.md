## Usage

Sources:

- [docs/assert.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/assert.rst)
- [assert/assert--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/assert/assert--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [assert/assert.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/assert/assert.control)

`assert` is Oliphant’s SQL assertion toolkit. Successful checks return normally; failed checks raise an exception that can fail a test transaction.

### Core Workflow

```sql
CREATE EXTENSION "assert";
SELECT assert_equals(1, 1);
SELECT assert_is_not_null(1);
SELECT assert_raises('UTA08', 'SELECT assert_equals(1, 2)');
```

### Operational Boundaries

These are legacy pure-SQL/PLpgSQL utilities originally documented for PostgreSQL 9.x. Current-major compatibility is unverified. Installation uses the control defaults and needs administrator privileges; no shared library or preload is required. `assert_equals` and `assert_not_equals` do not reject NULL arguments; combine them with `assert_is_null` or `assert_is_not_null` when required. Object checks include `assert_table_exists`, `assert_column_exists`, `assert_function_exists` and `assert_trigger_exists`. `assert_raises` executes caller-supplied SQL and verifies its SQLSTATE, so use trusted test statements.
