# Rx kevin

## Reactive Programming

```
users = ['aria', 'latte', 'choco-latte']

for user in users:
  getFollowers(user, { resultJson ->
    followers = resultJson.toFollowers()

    for follower in followers:
      getRepos(follower, { resultJson ->
        repos = resultJson.toRepos()
        println(repos)
      })
    })
```

## ReactiveX

```
users = ['aria', 'latte', 'choco-latte']

Observable.fromArray(users)
  .flatMap { user -> getFollowers(user) }
  .map { resultJson ->
    resultJson.toFollowers()
  }
  .flatten()
  .flatMap { follower -> getRepos(follower) }
  .map { resultJson -> resultJson.toRepos() }
  .subscribe { repos -> println(repos) }
```

## CSP1

```
users = ['aria', 'latte', 'choco-latte']

fnnc getFollowers(users)
  followers = makeChannel()
  go
    user <- users
    repos <- getFollowers*(user).toRepos()
  repos

func getRepos(followers)
  repos = makeChannel()
  go
    follower <- followers
    repos <- getRepos*(follower).toRepos()
  repos

func printRepo(repos)
  go
    reop <- repos
    println(repo)

users = makeChannelFromArray(['aria', 'latte', 'choco-latte'])

followers = getFollowers(users)
repos = getRepos(followers)
printRepo(repos)
```

## CSP2

```
go func() {
  followers := getFollowers(users)
  for follower := range followers {
    repos := getRepos(follower)
    for repo := range repos {
      fmt.Println(<- repo)
    }
  }
}()
```

## Sync

```
users = ['aria', 'latte', 'choco-latte']

for user in users:
  followers = getFollowers(user).toFollowers()
  for follower in followers:
    repos = getRepos(follower).toRepos()
    println(repo)
```
