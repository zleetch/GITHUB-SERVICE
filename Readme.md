# Github Service

> This repository is using Python 3. This repository also using Python library such as PyGithub (Github API v3), os (Operation System), and argparse (Command-line parsing module)

## Table of contents ##
- [Setup](#setup)
- [Login](#login)
- [Repository](#repository)
- [Collaborator](#collaborator)

## Setup
For making this aplication without define the username and password or access token. You must define GITHUB_USERNAME and GITHUB_PASSWORD or GITHUB_ACCESSTOKEN on your local

```
$ export GITHUB_USERNAME=${username}
$ export GITHUB_PASSWORD=${password}
$ export GITHUB_ACCESSTOKEN=${accessToken}
```

## Login
> If you not define the environment variable on setup

- Login with username and password
```
$ python3 github-service.py ${function} ${action} -u ${username} -p ${password} ${ARGS}
```

- Login with access-token
```
$ python3 github-service.py ${function} ${action} --access ${accessToken} ${ARGS}
```

## Repository
> This command if you already setup the GITHUB_USERNAME and GITHUB_PASSWORD or GITHUB_ACCESSTOKEN in environment variable. If doesn't you must define the -u ${username} and -p ${password} or --access ${accessToken}

- Add repository
```
$ python3 github-service.py repo add -r ${repositoryName}
```

- Add private repository
```
$ python3 github-service.py repo add -r ${repositoryName} --private
```

- Add repository with init
```
$ python3 github-service.py repo add -r ${repositoryName} --init
```

- Delete repository
```
$ python3 github-service.py repo del -r ${repositoryName}
```

## Collaborator
> This command if you already setup the GITHUB_USERNAME and GITHUB_PASSWORD or GITHUB_ACCESSTOKEN in environment variable. If doesn't you must define the -u ${username} and -p ${password} or --access ${accessToken}

- Add one collaborator into repository
```
$ python3 github-service.py collab add -r ${repositoryName} -c ${collaboratorName}
```

- Add multiple collaborator into repository
```
$ python3 github-service.py collab add -r ${repositoryName} -c ${collaboratorName} -c ${collaboratorName}
```

- Delete one collaborator from repository
```
$ python3 github-service.py collab del -r ${repositoryName} -c ${collaboratorName}
```

- Delete multiple collaborator from repository
```
$ python3 github-service.py collab del -r ${repositoryName} -c ${collaboratorName} -c ${collaboratorName}
```
