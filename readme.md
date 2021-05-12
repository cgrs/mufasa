# 🦁 mufasa
<img src="docs/icon.png" height=100 />

AWS CLI MFA-enabled credentials bootstrapper.

It automates the process of obtaining a temporary access token using AWS STS, required for MFA-enabled accounts.

> *This tool is still in early development, expect breaking changes and other disasters.*

# Getting started

```sh
go install github.com/cgrs/mufasa
cp .mufasa.example.yaml ~/.mufasa.yaml
# provide totp_token in YAML
mufasa >> ~/.aws/credentials
aws --profile mfa <service>
```

# TODO

- [ ] support for flags and interactive prompt
- [ ] update credentials automatically

Logo based on:
> Lion by Icons Producer from the Noun Project