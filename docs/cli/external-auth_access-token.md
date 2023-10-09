<!-- DO NOT EDIT | GENERATED CONTENT -->

# external-auth access-token

Print auth for an external provider

## Usage

```console
coder external-auth access-token [flags] <provider>
```

## Description

```console
Print an access-token for an external auth provider. The access-token will be validated and sent to stdout with exit code 0. If a valid access-token cannot be obtained, the URL to authenticate will be sent to stdout with exit code 1
  - Ensure that the user is authenticated with GitHub before cloning.:

     $ #!/usr/bin/env sh

OUTPUT=$(coder external-auth access-token github)
if [ $? -eq 0 ]; then
  echo "Authenticated with GitHub"
else
  echo "Please authenticate with GitHub:"
  echo $OUTPUT
fi

```

## Options

### --s

|      |                   |
| ---- | ----------------- |
| Type | <code>bool</code> |

Do not print the URL or access token.