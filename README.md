Looking for the CLI docs? [Go Here](https://github.com/Nevermore-FMS/poesitory/blob/main/cli/poesitory/README.md)

# Poesitory

The Plugin repository for Nevermore FMS

## Running locally

Poesitory uses `docker-compose` for easy one command setup. Simply run `docker-compose up -d --build` in this directory with proper environment variables set (see below) to spin up a poesitory environment.

### Environment variables

| Environment Variables  | Purpose                                                                                                                                                              | Required |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|
| GITHUB_CLIENT_ID       | The Client ID used for GITHUB login                                                                                                                                  | Yes      |
| GITHUB_CLIENT_SECRET   | The Client Secret used for GITHUB login                                                                                                                              | Yes      |
| GITLAB_CLIENT_ID       | The Client ID used for GITLAB login                                                                                                                                  | Yes      |
| GITLAB_CLIENT_SECRET   | The Client Secret used for GITLAB login                                                                                                                              | Yes      |
| POESITORY_BASE_URI     | The hostname that Poesitory will be accessed on                                                                                                                      | Yes      |
| POESITORY_CDN_URI      | The hostname that Poesitory CDN will be accessed on                                                                                                                  | Yes      |
| POESITORY_SECRET       | A random string used as an internal secret                                                                                                                           | Yes      |
| POESITORY_DEV_INSECURE | Set to `true` when the CDN will be accessed with a self signed certificate                                                                                           | No       |
| POESITORY_CERTRESOLVER | Set to `cre` to use an ACME certificate resolver. If left unset, a self signed certificate will be generated (and you should set `POESITORY_DEV_INSECURE` to `true`) | No       |
| POESITORY_CASERVER     | ACME URL to use for certificate resolution. If left unset, will use https://acme-v02.api.letsencrypt.org/directory (Let's Encrypt)                                   | No       |
| POESITORY_CA_EMAIL     | Email to use with ACME certificate request                                                                                                                           | No       |

A full start-up command might look like this:

```bash
GITHUB_CLIENT_ID=exampleclientid \
GITHUB_CLIENT_SECRET=exampleclientsecret \
POESITORY_BASE_URI=poesitory.edgarallanohms.com \
POESITORY_CDN_URI=cdn.poesitory.edgarallanohms.com \
POESITORY_SECRET=examplesecret \
POESITORY_CERTRESOLVER=cre 
POESITORY_CA_EMAIL="frcteam5276@gmail.com" \
docker-compose up -d --build
```

## Identifiers

Plugins on poesitory use identifiers when pulling. Their structure is as follows:

```
pluginname[#channel][@version]
```

- If `channel` is omitted, it will default to STABLE
- If `version` is omitted, it will default to the latest version for the given channel
- `pluginname` is limited to lowercase letters and hyphens 
- `channel` is limited to uppercase letters and hyphens
- `version` must follow the format `major.minor.patch`. No element can be omitted.

### Examples

- `test-plugin#ALPHA@0.1.5` references test-plugin version 0.1.5 in the ALPHA channel
- `test-plugin@1.2.0` references test-plugin version 1.2.0 in the STABLE channel
- `test-plugin#BETA` references the latest version of test-plugin in the BETA channel
- `test-plugin` references the latest version of test-plugin in the STABLE channel

### Behaviour in Nevermore FMS.

It is important to understand that the behaviour of Nevermore FMS varies depending on how the identifier is specified, even if they currently point to the same version

- if `test-plugin` is specified, the Nevermore FMS will pull the latest version of test-plugin in the STABLE channel, and will auto update in the future
- if `test-plugin#BETA` is specified, Nevermore FMS will pull the latest version of test-plugin in the BETA channel, and will auto update in the future from the BETA channel
- if `test-plugin@1.2.0` is specified, Nevermore FMS will pull test-plugin version 1.2.0 in the STABLE channel and **WILL NOT** auto update

In the examples above `test-plugin` and `test-plugin@1.2.0` might currently reference the same plugin version, but the latter will cause Nevermore FMS to not auto update the plugin