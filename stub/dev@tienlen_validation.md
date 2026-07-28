## Usage

Sources:

- [Official database.dev package page](https://database.dev/dev/tienlen_validation)

`dev@tienlen_validation` — Tien Len validation functions for game server. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "dev@tienlen_validation";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `can_beat(picked_cards TEXT[], played_cards TEXT[], ignore_rules BOOLEAN DEFAULT FALSE)` is an extension function and returns `TABLE`.
- `contains_2(cards TEXT[])` is an extension function and returns `BOOLEAN`.
- `get_card_face_rank(card TEXT)` is an extension function and returns `INTEGER`.
- `get_card_suit_rank(card TEXT)` is an extension function and returns `INTEGER`.
- `get_face_and_suit(card TEXT)` is an extension function and returns `TABLE`.
- `get_hand_info(cards TEXT[])` is an extension function and returns `TABLE`.
- `get_rank(card TEXT)` is an extension function and returns `INTEGER`.
- `get_total_ranks(cards TEXT[])` is an extension function and returns `INTEGER`.
- `is_double_sequence(cards TEXT[])` is an extension function and returns `BOOLEAN`.
- `is_four_of_a_kind(cards TEXT[])` is an extension function and returns `BOOLEAN`.
- `is_pair(cards TEXT[])` is an extension function and returns `BOOLEAN`.
- `is_sequence(cards TEXT[])` is an extension function and returns `BOOLEAN`.
- `is_three_of_a_kind(cards TEXT[])` is an extension function and returns `BOOLEAN`.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
