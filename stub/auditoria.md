## Usage

Sources:

- [Official extension control file (auditoria.control)](https://github.com/zsnails/postgres-audit-log/blob/44368b5696450699618b11c26bfcc2ff40af9bda/auditoria.control)
- [Official extension SQL (auditoria--0.0.1.sql)](https://github.com/zsnails/postgres-audit-log/blob/44368b5696450699618b11c26bfcc2ff40af9bda/auditoria--0.0.1.sql)

`auditoria` — PostgreSQL auditing extension. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION auditoria;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `auditoria.registrar_actividad(tabla_afectada TEXT, tipo_operacion TEXT, datos_anteriores JSONB, datos_nuevos JSONB)` is an extension function and returns `VOID`.
- `auditoria.registrar_en_tablas_nuevas()` is an extension function and returns `event_trigger`.
- `auditoria.trig_registrar_actividad()` is an extension function and returns `TRIGGER`.
- `auditoria.registrar_auditores` is an extension procedure.
- `auditoria.registrar_en_todo_lado` is an extension procedure.
- `auditoria.registro_auditoria` is a table installed or managed by the extension.
- `auditoria` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
