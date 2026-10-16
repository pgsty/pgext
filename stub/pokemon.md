## Usage

Sources:

- [README.md](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/README.md)
- [client/extension/pokemon--0.1.0.sql](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/pokemon--0.1.0.sql)
- [client/extension/Cargo.toml](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/Cargo.toml)
- [LICENSE](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/LICENSE)
- [client/extension/pokemon.control](https://github.com/stillinbeta/PostgresPlaysPokemon/blob/fcdc0e3043da4c4a9a9b6cb2b5d587f6cd2b8602/client/extension/pokemon.control)

`pokemon` 0.1.0 exposes emulator state through foreign tables. It is the PostgreSQL client of PostgresPlaysPokemon, not an emulator running inside PostgreSQL.

### Core Workflow

```sql
CREATE EXTENSION pokemon;
IMPORT FOREIGN SCHEMA red FROM SERVER red INTO public;
SELECT * FROM party;
```

### Operational Boundaries

The official environment tests PostgreSQL 10 and requires the separate PyBoy-based server, its ROM data and a reachable Unix socket for gRPC. Creating the extension registers `pokemon_fdw()`, the `pokemon` wrapper and server `red`. Importing schema `red` exposes `party`, `inventory` and `story`: party/story support SELECT and UPDATE, while inventory also supports INSERT. Writes manipulate emulator state. Installation defaults to superuser-only; this historic integration does not establish support for current PostgreSQL majors or a remotely configurable connection.
