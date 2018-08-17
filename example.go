
package main

import "fmt"

func main() {
  users := []string{"aria", "latte", "choco-latte"}
  followers := make(chan string)
  repos := make(chan string)

  go func() {
    for _, user := range users {
      followers <- user
    }
    close(followers)
  }()

  go func() {
    for follower := range followers {
      repos <- follower
    }
    close(repos)
  }()

  for repo := range repos {
    fmt.Println(repo)
  }
}
