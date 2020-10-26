## Git Repositories 
This is a React and Golang project that was developed to display a list of Github public repositories and commits. The Golang api is responsible to feetch data from Github api parse and serve to out React dashboard.

##### Commands to run the React dashboard  (http://localhost:8080/)
  ```
    cd github-dashboard
    npm install
    npm start
```
##### Commands to run the Go project (http://localhost:8090/)
  ```
  cd github-api
  dep ensure
  go run cmd/main.go
```

##### Run Go tests
  ```
  cd clients
  go test
```

##### Features implemented:
- Task 0 - Create a local API server - Golang
- Task 1 - Connect to the Github API
- Task 2 - Load last commits
