## What is this?
This is a support tool to use the GitHub API.

### GitHub API

GitHub provices an API to create/update/delete repositories. detail: https://developer.github.com/v3/
To use this API, you need 
- To get access_token
- Call API with access_token

This tool supports to use GitHub API.

## Get Access Token

At first, you have to get access token.

### How to get

1. Create GitHub OAuth Application at [here](https://github.com/settings/applications/new).

GitHub APIv3 OAuth sequence is folliwing:

Authorization callback URL: Your authorization server

Aboute sequence, please see sequence.pu

Authorization callback URL is your authorization server's URI.
And you can get access_token by using code which is in callback URL's query.

Please see [Authorizing OAuth Apps](https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/) to see more detail.

### Support tool: Authorization callback URL server

This repository provides a server to get access_token.
#### How to use

1. Install go 1.13

```
go get golang.org/dl/go1.13
~/go/bin/go1.13 download
```

2. build

```
cd server
make
```

3. run

```
#please set client_key and secret from  https://github.com/settings/developers.
export SERVER_PORT=port
export SERVER_CERT=cert_path
export SERVER_KEY=key_path
export CLIENT_KEY=client_key
export CLIENT_SECRET=secret

 ./server/auth_server
```

4. Call `https://github.com/login/oauth/authorize?client_id=$CLIENT_ID&scope=repo`

You can get result as 

```
{"access_token":"yyyy"}
```

## Call API with access_token

```
cd repos
./create_repos.sh testrepos yyyy
```

```
$ curl -H "Authorization: bearer yyyy" -X POST \
    -d '{"name":"repos_name","auto_init":true or false}' \
    https://api.github.com/user/repos
```

## Reference

[GitHub API v3 sample](https://qiita.com/ngs/items/34e51186a485c705ffdb)

