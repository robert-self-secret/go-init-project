# How to use this project template

---

- clone this repository:

  ```bash
  git clone https://github.com/robert-self-secret/go-init-project.git
  ```

- rename the folder from the git clone result, from `go-init-project` to `[your-project-folder-name]`
  <br>
- remove git remote for this repository:
  ```bash
  git remote rm origin
  ```
- link to your new github repository:

  ```bash
  git remote add origin [your github repository]
  ```

- Change the name of the old golang module to your module name in the `go.mod` file

  from:

  ```go
  module github.com/robert-self-secret/go-init-project.git

  go 1.21.1
  ```

  to

  ```go
  module [your golang project module name]

  go 1.21.1
  ```

- To check old import path, on Unix-based systems (Linux, macOS), you can use the grep command to find all references to the old module name:

  ```bash
  grep -r "github.com/robert-self-secret/go-init-project.git" .
  ```

- To automatically replace old import paths with the new module name, you can use the sed command:

  ```bash
  find . -type f -name "*.go" -exec sed -i 's/github.com\/robert-self-secret\/go-init-project.git/[your-project-module]/g' {} +
  ```

  replace [your-project-module] with your module name, example module name : `github.com/robert-self-secret/go-authentication.git` :

  ```bash
  find . -type f -name "*.go" -exec sed -i 's/github.com\/robert-self-secret\/go-init-project.git/github.com\/robert-self-secret\/go-authentication.git/g' {} +
  ```

- After updating all import paths, run the following command to tidy up the module dependencies:
  ```go
  go mod tidy
  ```

## Good luck
