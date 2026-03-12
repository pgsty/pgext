


## Usage

> [omni_types: Advanced types](https://docs.omnigres.org/omni_types/function_signature_types/)

The `omni_types` extension provides advanced type utilities, including function signature types that capture complete function signatures and allow direct invocation.

### Define a Function Signature Type

**Explicit definition:**

```sql
SELECT omni_types.function_signature_type('sig', 'text', 'int');
-- Creates type 'sig' for functions accepting text and returning int
```

**From existing function:**

```sql
SELECT omni_types.function_signature_type_of('sig', 'length(text)');
```

### Cast Functions to the Type

```sql
SELECT 'length'::sig;
```

For non-failing validation:

```sql
SELECT sig_conforming_function('length');  -- returns NULL if no match
```

### Invoke Typed Functions

```sql
SELECT call_sig('length', 'hello');  -- returns 5
```

The `call_<TYPE>` function is auto-generated for each signature type and executes the referenced function with the provided arguments.
