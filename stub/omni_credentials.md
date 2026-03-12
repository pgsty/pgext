


## Usage

> [omni_credentials: Application credential management](https://docs.omnigres.org/omni_credentials/credentials/)

The `omni_credentials` extension provides secure, encrypted credential storage with access control. It is a templated extension.

### Setup

```sql
SELECT omni_credentials.instantiate(
    schema  => 'omni_credentials',
    env_var => 'OMNI_CREDENTIALS_MASTER_PASSWORD'
);
```

### Credentials View

The primary interface is the `credentials` view:

| Column      | Type            | Purpose                                      |
|:------------|:----------------|:---------------------------------------------|
| `name`      | text            | Credential identifier                        |
| `value`     | bytea           | Encrypted credential data                    |
| `kind`      | credential_kind | Type: `api_key`, `api_secret`, `password`, etc. |
| `principal` | regrole         | PostgreSQL role owning the credential        |
| `scope`     | jsonb           | Resource constraints (`{"all": true}` for universal) |

Query and update the view directly; changes propagate automatically. Row-Level Security ensures only roles with grants on the `principal` can access credentials.

### File Store (Development)

```sql
SELECT omni_credentials.instantiate_file_store(filename, schema);
-- Reload credentials from file:
SELECT credential_file_store_reload(filename);
```

Imports existing file records into the encrypted store and exports missing credentials from the table to the file. Auto-updates the file on credential changes.
