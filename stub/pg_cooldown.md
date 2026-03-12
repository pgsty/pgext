


## Usage

> [pg_cooldown: remove buffered pages for specific relations](https://github.com/rbergm/pg_cooldown)

pg_cooldown is the complement to `pg_prewarm`: it removes all buffered pages of a specific relation from the shared buffer, useful for simulating cold-start scenarios in research and testing.

### Remove Pages from Shared Buffer

```sql
-- Remove all data pages of a table
SELECT pg_cooldown('my_relation');

-- Remove pages of a primary key index
SELECT pg_cooldown('my_relation_pkey');

-- Remove pages of any index
SELECT pg_cooldown('my_index_name');
```
