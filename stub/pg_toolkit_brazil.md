## Usage

Sources:

- [Official documentation](https://github.com/fabriziomello/pg_toolkit_brazil/blob/25aa4d0472582d17af8b84e8c80cf8a1db8b8943/README.md)
- [Extension control file](https://github.com/fabriziomello/pg_toolkit_brazil/blob/25aa4d0472582d17af8b84e8c80cf8a1db8b8943/pg_toolkit_brazil.control)
- [Official repository](https://github.com/fabriziomello/pg_toolkit_brazil)

`pg_toolkit_brazil` Brazilian CPF and CNPJ types, validation, formatting, casts, and operators.

### Enablement

Install the files for the intended server, then create `pg_toolkit_brazil` in the target database:

```sql
CREATE EXTENSION pg_toolkit_brazil;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE TABLE pessoa (id bigint GENERATED ALWAYS AS IDENTITY, cpf cpf);
INSERT INTO pessoa (cpf) VALUES (cpf '40100276300');
SELECT id, cpf FROM pessoa;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `cnpj` | TYPE | User-facing data type created by the extension. |
| `cnpj_cmp` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_eq` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_ge` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_gt` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_le` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_lt` | FUNCTION | Callable function from the reviewed install surface. |
| `cnpj_ne` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
