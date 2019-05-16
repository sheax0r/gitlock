# Gitlock

A small go program to test that branches in a github repository can be used as a simple locking mechanis  a simple locking mechanism.

## Building it

```bash
go install ./...
```

## Running it

```bash
export GITHUB_USER=...
export GITHUB_TOKEN=...
export GITHUB_REPO=...

gitlock lock
gitlock unlock
```

## Sample output

```
❯ gitlock lock; echo $?
0

❯ gitlock lock; echo $?
2019/05/15 22:02:24 {"message":"Reference already exists","documentation_url":"https://developer.github.com/v3/git/refs/#create-a-reference"}
1

❯ gitlock unlock; echo $?
0

❯ gitlock unlock; echo $?
2019/05/15 22:02:39 {"message":"Reference does not exist","documentation_url":"https://developer.github.com/v3/git/refs/#delete-a-reference"}
1
```
