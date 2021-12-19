# auth0
auth0 cli

## setup

create a config file `~/.auth0/config.yaml` with the following content
```yaml
default:
  domain: "<domain>"
  clientId: "<client-id>"
  clientSecret: "<client-secret>"
```

Where `default` is a profile. This allows to store multiple clients. Profile can be selected by `profile` or `p` flag.

## run

```
auth0 list <user|log>
Flags:
  -h, --help             help for auth0
  -p, --profile string   auth0 profile (default "default")
```
