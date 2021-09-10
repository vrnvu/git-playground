# flow

hello, fpero 

today we are going to learn how to use git and github

- hello, sounds cool but i already kno... 

shh dont worry buddy lets start

## fork

go to any github repo (we assume in this post everything is there, but this content works with anything but gerrit and such) and fork it

- how i do that?

[im not rewritting the getting started for you or translating it to spanish to sell it to my viewers for 9.99$](https://docs.github.com/es/get-started/quickstart)

just google before asking

- ok sorry 
  
## clone fork to local 

clone your fork

configure local envs if needed, some projects require this to build and test

for example, to contribute to k8s I have to set up my working_dir
```bash
export working_dir="$(go env GOPATH)/src/k8s.io"
```

- what is a local env? what is the export? you are going to fast

...

add remote upstream

```bash
git remote add upstream https://github.com/org/project.git
```

good practice, never push to upstream master

```bash
git remote set-url --push upstream no_push
```

ok you may think that you are set up, you are not run the verification tests and check that everything actually works, there is probably some script in a /hack or /make folder

## sync with upstream and origin

always rebase, either set your .gitconfig to flag pull --rebase true or do it manually

in most projects, merging is not allowed and the CI bot will complain and stop the pipelines from running

i'll be nice and add some references for my fellow fperos:

- [pull --rebase vs --merge](https://sdqweb.ipd.kit.edu/wiki/Git_pull_--rebase_vs._--merge)
- [merge vs rebase](https://mislav.net/2013/02/merge-vs-rebase/)

how does it look like in action?

```bash
git fetch XXXX
git checkout master
git rebase XXXX/master
```

where XXXX can be upstream or origin

question for the reader, where does the pull command by default pull from? 

```bash
git pull --rebase
```

- i don't know

huh

- i did not understand the XXXX, too difficult  

ok, let me give you a detailed example

i'm contributing to kubernetes and want to fetch and rebase 
```bash
git fetch upstream
git checkout master
git rebase upstream/master
```

i'm contributing to kubernetes in my fork with a collegue, author note do not do this pls, and I want to sync his changes in my branch
```bash
git fetch origin
git checkout master
git rebase origin/master
```

- ok I understand now so easy

## branches

sync and branch from master 99% of the time 

```bash
git checkout -b foo
```

and push it

```bash
git push --set-upstream origin foo
```

go to github and create a pull request / merge request

do bot stuff if needed if contributing to some open source

before submitting a pull request run local tests and checks

## work

you need to know commiting, amending, rebasing, bisecting, cherry picking and reverting 

you will be happy to know how to use worktrees (bare git) if it makes sense for your workflow honestly i see people constantly messing things up because they do not know how git works lol