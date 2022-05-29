# 🦁 mufasa

## Deprecation notice
Since this tool is still in development and there are better options out there, I've decided to archive this repo and stop further development on this project. If you want to manage your AWS (and Azure) sessions with ease, you can use [Leapp](https://www.leapp.cloud/).

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
