


## Usage

> [omni_var: Scoped variables](https://docs.omnigres.org/omni_var/variables/)

The `omni_var` extension provides typed variables with transaction, session, and statement scopes.

### Set Variables

```sql
SELECT omni_var.set('my_var', true);              -- transaction scope
SELECT omni_var.set_session('my_var', true);       -- session scope
SELECT omni_var.set_statement('my_var', true);     -- statement scope
```

For explicit typing:

```sql
SELECT omni_var.set('text_var', 'value'::text);
```

### Get Variables

```sql
SELECT omni_var.get('my_var', false);              -- transaction scope, default=false
SELECT omni_var.get_session('my_var', false);       -- session scope
SELECT omni_var.get_statement('my_var', false);     -- statement scope
```

The default value serves dual purposes: returned when the variable is not found, and its type specifies the expected return type.

### Behavior

- If a variable is explicitly set to `null`, `get` returns `null` (not the default)
- Mismatched type specifications raise errors: `ERROR: type mismatch DETAIL: expected integer, got boolean`
- Transaction-scoped variables are bound by the enclosing transaction lifetime
- Session variables persist throughout the session duration
