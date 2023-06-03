# relaycompmul

The relaycompmul is a linter for GraphQL schemas.
This linter outout errors when fields that follow [relay cursor connections specificaiton](https://relay.dev/graphql/connections.htm) and do not have `@complexity` directive or lack `mul` arguments of `@complexity`.
