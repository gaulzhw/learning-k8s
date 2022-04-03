# git forking workflow



## Flows

- Fork a repository

- Clone your fork

  ```shell
  git clone https://user@bitbucket.org/user/repo.git
  ```

- Add a remote

  ```shell
  git remote add upstream https://bitbucket.org/maintainer/repo
  git remote add upstream https://user@bitbucket.org/maintainer/repo.git
  ```

- Working in a branch: making & pushing changes

  ```shell
  git checkout -b some-feature
  
  # Edit some code
  git commit -a -m "Add first draft of some feature"
  
  git pull upstream main
  ```

- Making a Pull Request

  ```shell
  git push origin feature-branch
  ```



## References

https://www.atlassian.com/git/tutorials/comparing-workflows/forking-workflow