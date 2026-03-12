


## Usage

> [omni_os: Operating system integration](https://docs.omnigres.org/omni_os/intro/)

The `omni_os` extension provides access to operating system facilities. Intended for superuser use.

### Environment Variables

The `omni_os.env` view exposes system environment variables:

```sql
SELECT * FROM omni_os.env;
```

| Column     | Type |
|:-----------|:-----|
| `variable` | text |
| `value`    | text |

Works similarly to the `env` command-line utility, returning all environment variables visible to the PostgreSQL process.
