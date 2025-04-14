# D★Flag

A simple but fast feature flag server that loads feature flag data from a static file.

## Why?

Because feature flags are powerful and you need a solution with a very low barrier to entry. D★Flag is a single executable with no dependencies — not even a database.

Pro:

- It is fast. Data is served from memory, after all. [wrk](https://github.com/wg/wrk) says it can handle 70k+ rps on my MacBook Air.

- It has no infrastructure dependencies.

- It requires minimal resources (a few megabytes of memory at most).

- If changes are done through a version control system, you get an audit log for free -- your version control commit list. If you use pull request, you get a review step for free too.

Con:

- Toggling a feature flag requires redeployment (but hopefully that is quick enough)

- It does not provide advanced use cases like A/B testing or scheduled toggles.

## Defining feature flags

Feature flag definitions reside in `data/feature_flags.hcl`. It contains entries like these:

```hcl
feature_flag "payments_v2" {
  team_responsible = "Payments team <payments@example.com>"
  contact_person = "Denis Defreyne <denis@example.com>"

  environment "staging" {
    enabled = true
  }

  environment "production" {
    enabled = false
  }
}
```

## Running

The built `dflag` executable runs like any other server:

```
dflag
```

Configuration is done through the environment, and accepts the following variables:

- `PORT`: The port to listen on (default `3000`)

- `FEATURE_FLAGS_DATA_FILE_PATH`: The path to the data file containing feature flags (default: `data/feature_flags.hcl`)

## Checking

To verify the configuration and the feature flag data without starting the server, pass the `-check` flag:

```
dflag -check
```

This will errors if there is a misconfiguration.

## API

To check whether a given feature flag is enabled or not in a given environment, mke a GET request to <code>/feature-flags/<var>FLAGNAME</var>?env=<var>ENVNAME</var></code>, replacing <var>FLAGNAME</var> with the name of the feature flag, and <var>ENVNAME</var> with the name of the environment:

```
GET /feature-flags/payments_v2?env=staging

HTTP/1.1 200 OK
Content-Type: application/json
{"enabled":true}
```

There is also the `/healthz` endpoint for health checks.
