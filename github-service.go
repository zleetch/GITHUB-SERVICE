package main

import (
  "fmt"
  "flag"

  "github.com/google/go-github/github"
  "context"
  "golang.org/x/oauth2"
  // "reflect"
)

type Operation struct {
  Task        string
  Action      string
}

type Permission struct {
  Pull        bool
  Push        bool
  Admin       bool
}

type Collaborator struct {
  Name        string
  Permission
}

type Repository struct {
  Name        string
  Privacy     bool
}

//SEARCHREPO
func (r Repository) SearchRepo(ctx context.Context, client *github.Client) bool {
  repos, _, _ := client.Repositories.List(ctx, "", nil)
  for _, repo := range repos {
    if *repo.Name == r.Name {
      fmt.Println(r.Name, "REPOSITORY FOUND")
      return true
    }
  }
  fmt.Println(r.Name, "REPOSITORY NOT FOUND")
  return false
}

//CREATEREPO
func (r Repository) CreateRepo(ctx context.Context, client *github.Client) {
  repo := &github.Repository{
    Name: github.String(r.Name),
    Private: github.Bool(r.Privacy),
  }
  if match := r.SearchRepo(ctx, client); match {
    fmt.Println(r.Name, "REPO ALREADY EXIST")
  } else {
    fmt.Println("CREATING ", r.Name)
    client.Repositories.Create(ctx, "", repo)
  }
}

func (o Operation) Todo(r Repository, c Collaborator) {
}

func main() {
  //ARGUMENT
  task := flag.String("task", "none", "Operation (REPO, COLLAB)")
  action := flag.String("action", "none", "Operation (CREATE, DELETE, EDIT)")
  repoName := flag.String("repo", "reponame", "Repository Name")
  repoPrivacy := flag.Bool("privacy", false, "Repository Privacy")
  permissionPull := flag.Bool("pull", false, "Collab Pull Privacy")
  permissionPush := flag.Bool("push", false, "Collab Push Privacy")
  permissionAdmin := flag.Bool("admin", false, "Collab Admin Privacy")
  collabName := flag.String("collab", "collabname", "Collaborator Name")
  flag.Parse()

  //SETTER
  var operation = Operation{Task : *task, Action : *action}
  var repo = Repository{Name : *repoName, Privacy : *repoPrivacy}
  var permission = Permission{Pull : *permissionPull, Push : *permissionPush, Admin : *permissionAdmin}
  var collab = Collaborator{Name : *collabName, Permission : permission}

  //AUTHENTICATION
  //LOGIN
  ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "ACCESSTOKEN"},
    //
  )
  tc := oauth2.NewClient(ctx, ts)
  client := github.NewClient(tc)
  fmt.Println(operation, repo, permission, collab)
  operation.Todo(repo, collab)
  repo.CreateRepo(ctx, client)
}
