# Hookie

A simple server that turns webhook calls into command execution. This enables simple and effective
automation scenario's to be scripted in bash script.

## Example

Here is an simple example that pull new content from a git repository on every webhook call:

```
hookie -address "0.0.0.0:80" sh -c 'cd /var/www/site && git fetch --all && git reset --hard origin/master'

```
## Installation

Official binary distributions are available for Windows, Mac OS and Linux in both 32 bit (386) and 64 bit (amd64).

To install hookie you download the [latest version](https://github.com/pjvds/hookie/releases/latest) for your platform and extract the archive to a directory in your `$PATH`.

### Linux (64bit)

```
curl -O https://github.com/pjvds/hookie/releases/download/v1.1/hookie-v1.1-linux-amd64.tar.gz
tar -C /usr/local/bin -xzf hookie-v1.1-linux-amd64.tar.gz
```

### Linux (32bit)

```
curl -O https://github.com/pjvds/hookie/releases/download/v1.1/hookie-v1.1-linux-386.tar.gz
tar -C /usr/local/bin -xzf hookie-v1.1-linux-386.tar.gz
```

### Mac OS (64bit)

```
curl -O https://github.com/pjvds/hookie/releases/download/v1.1/hookie-v1.1-darwin-amd64.tar.gz
tar -C /usr/local/bin -xzf hookie-v1.1-darwin-amd64.tar.gz
```

### Mac OS (32bit)

```
curl -O https://github.com/pjvds/hookie/releases/download/v1.1/hookie-v1.1-darwin-386.tar.gz
tar -C /usr/local/bin -xzf hookie-v1.1-darwin-386.tar.gz
```

### Mac OS (via Homebrew)

Alternative you can use homebrew to easily install hookie on Mac OS:

```
brew install pjvds/homebrew-tools/hookie
```

## Github signature validation

If you are using hookie to listen for webhook calls send from Github you can increase the security
by validating the origin of the payload. You can add a secret token to the configuration of your
webhook at Github and use the `-github-secret` argument to configure hookie to validate the signature.

You'll need to set up your secret token at Github and configure hookie with it.

Navigate to the repository where you're setting up your webhook.

1. Navigate to the repository where you're setting up your webhook.
2. Fill out the Secret textbox. Use a random string with high entropy, for example an GUID from
   [www.uuidgenerator.net](https://www.uuidgenerator.net/) like `d846f12d-e46e-4d24-bea7-36979223bb4a`.
3. Click Update Webhook.
4. Start hookie with the secret token argument: `hookie -secret-token='d846f12d-e46e-4d24-bea7-36979223bb4a' script.sh`
