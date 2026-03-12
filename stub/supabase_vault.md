


## Usage

> [supabase_vault: Encrypted secret storage for Supabase](https://github.com/supabase/vault)

Supabase Vault provides a `vault.secrets` table to store sensitive information (API keys, tokens, etc.) encrypted at rest. Decryption happens on the fly through the `vault.decrypted_secrets` view.

```sql
CREATE EXTENSION supabase_vault CASCADE;
```

### Storing Secrets

```sql
INSERT INTO vault.secrets (secret) VALUES ('s3kre3t_k3y') RETURNING *;

-- Or use the helper function:
SELECT vault.create_secret('another_s3kre3t');

-- With optional name and description:
SELECT vault.create_secret('my_secret', 'unique_name', 'This is the description');
```

### Reading Secrets

The `vault.secrets` table stores data encrypted. Use the `vault.decrypted_secrets` view to read decrypted values:

```sql
SELECT * FROM vault.decrypted_secrets ORDER BY created_at DESC LIMIT 3;
-- Includes a `decrypted_secret` column with the plaintext value
```

### Updating Secrets

```sql
SELECT vault.update_secret(
    '7095d222-efe5-4cd5-b5c6-5755b451e223',
    'n3w_upd@ted_s3kret',
    'updated_unique_name',
    'This is the updated description'
);
```

### Security Note

Turn off statement logging to prevent secrets from appearing in logs:

```sql
ALTER SYSTEM SET statement_log = 'none';
```
