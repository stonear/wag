# Fork of `wag`

This repository is a fork of the original [wag](https://github.com/Clever/wag). The motivation behind forking and modifying this project stems from a few key challenges I faced while maintaining a service that depends on `wag` and OpenTelemetry.

## Motivation

I maintain a legacy service that relies on [wag](https://github.com/Clever/wag) to automatically generate a server based on a `swagger.yml` specification. However, our CISO team consistently raised concerns that the version of OpenTelemetry used in the project was outdated and needed to be upgraded.

### Challenges

Upgrading OpenTelemetry is not a straightforward process due to the following reasons:

- **Breaking Changes**: OpenTelemetry frequently introduces breaking changes in new versions, making it difficult to upgrade without causing issues in our service.
- **Dependency Version Mismatches**: Each direct dependency in the project has its own indirect dependencies on OpenTelemetry, and these dependencies often require different versions of OpenTelemetry. This version mismatch results in complex dependency resolution issues.

As a result, upgrading OpenTelemetry while maintaining compatibility across our dependencies has become increasingly difficult and time-consuming.

### Solution

To resolve these challenges, I decided to fork this repository and modify it to prevent it from generating unnecessary dependencies and changes that our team does not need. By doing this, I can have more control over the versions of OpenTelemetry used in our project, making it easier to maintain and upgrade the service without encountering the aforementioned issues.

## Changes Made

- Custom adjustments to dependencies to avoid unnecessary OpenTelemetry dependencies.
- Other tweaks to streamline the integration with the existing service requirements.

## Purpose of the Fork

It is not intended to be a long-term solution for the wider community, but rather a temporary measure to meet our internal needs.

## Usage

Please refer to the original [Clever/wag documentation](https://github.com/Clever/wag) for how to use this repository. However, note that this fork might have minor differences due to the custom modifications.

If you are facing similar issues, feel free to open an issue or contribute to the repository.
