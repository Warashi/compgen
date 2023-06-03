# relaycompmul

The relaycompmul is a linter for GraphQL schemas.
This linter outout errors when fields that follow [relay cursor connections specificaiton](https://relay.dev/graphql/connections.htm) and do not have `@complexity` directive or lack `mul` arguments of `@complexity`.

## Quick Start
### installation

```sh
$ go install github.com/Warashi/compgen/linter/relaycompmul/cmd/relaycompmul@latest
```

The relaycompmul command has two flags, schema and query which will be parsed and analyzed by relaycompmul's Analyzer.

### Usage
```sh
$ relaycompmul -schema="server/graphql/schema/**/*.graphql" -query="client/**/*.graphql"
```

The default value of schema is "schema/*/**.graphql" and query is query/*/**.graphql.

