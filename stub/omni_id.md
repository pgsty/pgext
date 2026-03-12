


## Usage

> [omni_id: Identity types](https://docs.omnigres.org/omni_id/identity_type/)

The `omni_id` extension introduces custom identity types that prevent comparison errors between different ID types, catching bugs at query time.

### Create an Identity Type

```sql
CREATE EXTENSION omni_id;
SELECT identity_type('user_id');
```

This creates a `user_id` type backed by `bigint` (default) with an auto-created sequence and constructor function.

### Parameters

| Parameter          | Default   | Description                                      |
|:-------------------|:----------|:-------------------------------------------------|
| `type`             | `bigint`  | Base type (`int`, `smallint`, `uuid`)            |
| `sequence`         | `<type>_seq` | Sequence identifier                           |
| `create_sequence`  | `true`    | Auto-create sequence                             |
| `increment`        | `1`       | Sequence step value                              |
| `cache`            | --        | Pre-allocate sequence numbers                    |
| `cycle`            | --        | Enable wraparound at limits                      |
| `constructor`      | --        | Constructor function name                        |
| `operator_schema`  | `public`  | Schema for operators                             |

### Auto-Generated Functions

- `user_id(1)` -- Constructor
- `user_id_nextval()` -- Get next value
- `user_id_currval()` -- Get current value
- `user_id_setval(user_id, bool)` -- Set value
