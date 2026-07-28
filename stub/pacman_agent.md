## Usage

Sources:

- [Official upstream README](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/README.md)
- [Official extension control file (pacman_agent.control)](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/pacman_agent.control)
- [Official extension SQL (pacman_agent--0.1.0.sql)](https://github.com/polkiloo/pacman/blob/c64a68e650bcc9fcc079ee3487a57004a3690720/postgresql/pacman_agent/sql/pacman_agent--0.1.0.sql)

`pacman_agent` — This directory contains the initial PACMAN PostgreSQL background-worker extension scaffold. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pacman_agent;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
