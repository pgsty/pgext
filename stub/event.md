## Usage

Sources:

- [Official upstream README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/event/README.md)
- [Official extension control file (event.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/event/event.control)

`event` — Now imagine running the following SQL to change the data:. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION event;

insert into person (name, score) values ('Don Pablo', 14);
update person set name='Sandy Jones', score=score+3 where id=3;
delete from person where id=4;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.5.0`.
- Install the confirmed extension dependencies first: `meta`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
