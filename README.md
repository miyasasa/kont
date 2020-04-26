
# kont


## Overview

kont is the application developed by Go as open-source. It provides you ability that assigns as many reviewers as given
to pull/merge request to the protected branch by based some algorithms. It aims to improve code quality in code-review process.

## Getting Started

To run this application, you'll need [Docker](https://docs.docker.com/get-docker/) installed in your computer. From your command line:

```
docker run -p 1903:1903 vyasinw/kont
```
Via above command, docker will pull(if it is not exist on your locale) and run the latest image of the application on 1903 port(default-port). 

In addition, 
You can change the port by set **SERVER_PORT** environment variable and add volume for */var/lib/kont* path for ensuring the persistence of the data.


```
docker run -p 9090:9090 -e SERVER_PORT=9090  -v /var/lib/kont:/var/lib/kont vyasinw/kont
```
## Usage
For now, only available with The API which is generally RESTFUL and returns results in JSON.

### Resource components and identifiers
Resource components can be used in conjunction with identifiers to do CRUD operations for that identifier on repositories.

| resource          | method        | description  |
| -------------     |:-------------:|:-------------|
| /repositories       | POST          | create/update the given repository, by name |
| /repositories       | GET           | returns a list of all repositories |
| /repositories/:name | GET           | returns the repository named |
| /repositories/:name | DELETE        | deletes the repository named |

#### Sample Repository in kont
kont needs some information about the remote repository for starting to observe pull/merge requests. So, a repository record
must consist of below parameters in JSON format.

```json5

{
  "host": "http://10.120.0.145:7990",
  "token": "Bearer Ab...",
  "projectName": "BESG",
  "name": "core-network",
  "developmentBranch": "develop",
  "provider": "BITBUCKET",
  "defaultComment": "Merhaba @{{name}} \n ## **Reviewer koda bakmadan evvel, kendin bir kez daha review etmeye ne dersin? Eminim ELF gözlerin bişeyler görecektir.** \n Bunun için aşağıdaki maddeleri kontrol edebilirsin \n * Reformat \n * SonarLint \n * Analyze -> Inspect code \n * Mimari olarak düzgün mü? (heryer heryerde olmasın lütfen) \n * Test isimlerini daha anlaşılır yapabilirsin(given-when-then) \n * Fazla test, göz çıkarmaz \n\n Kolay Gelsin Hacım :) ",
  "stages": [
    {
      "name": "Stage1",
      "policy": "RANDOMINAVAILABLE",
      "reviewers": [
        {
          "priority": 3,
          "user": {
            "name": "atiba",
            "displayName": "Atiba Hutchinson"
          }
        },
        {
          "priority": 2,
          "user": {
            "name": "vida",
            "displayName": "Domagoj Vida"
          }
        },
        {
          "priority": 1,
          "user": {
            "name": "gonul",
            "displayName": "Gokhan Gonul"
          }
        }
      ]
    },
    {
      "name": "Stage2",
      "policy": "RANDOMINAVAILABLE",
      "reviewers": [
        {
          "priority": 1,
          "user": {
            "name": "pepe",
            "displayName": "Pepe"
          }
        },
        {
          "priority": 2,
          "user": {
            "name": "fabri",
            "displayName": "Fabricio Agosto"
          }
        }
      ]
    },
    {
      "name": "Stage3",
      "policy": "RANDOMINAVAILABLE",
      "reviewers": [
        {
          "priority": 1,
          "user": {
            "name": "sergen",
            "displayName": "Sergen Yalcin"
          }
        }
      ]
    }
  ]
}

```
**Keys Description**

**host**: Consists of scheme and host of the provider for your Git repository.

**token**: Personal access bearer token which has necessary permissions(repository write/admin)

**projectName**: Key of the project which contains the repository

**name**: Name of the remote repository

**developmentBranch**: Development branch of the remote repository for observing pull/merge requests to this branch.

**provider**: String Upper-case provider name, any of ("GITHUB", "BITBUCKET", "GITLAB")

**defaultComment**(Optional): Default comment text for a new pull/merge request(Markdown syntax in string), 
```{{name}}``` statement will replaced by the Author name of the pull request 

**stage**: Each stage consists of name, policy and list of reviewer keys. 
kont selects a *available reviewer* in the given reviewer list based *policy type*. 
So, reviewers will be assigned as many as the number of stages.

* ```availability``` is that, a reviewer has not been assigned to any pull/merge request or 
a reviewer has approved all pull/merge requests which he/she had been assigned.

* ```policy``` is the strategy to select a reviewer, it can be any of ("RANDOMINAVAILABLE","BYPRIORITYINAVAILABLE")
    * ```RANDOMINAVAILABLE```: Selects a random reviewer in available reviewers in the stage.
    * ```BYPRIORITYINAVAILABLE```: Sorts available reviewers and select first which has high priority value.

    if there is no available reviewer in stage, kont will ignore policy-type and assign a reviewer randomly

Notes:
* An author of the pull request can not be assigned
* If a stage contains only one reviewer who is author of a pull request at the same time, kont select one more reviewer from next stage
(All stages compose a circle)

## Notes
1. kont just has integration with **Bitbucket-Server(based on Rest Api v1)** currently

## How to Contribute
1. Clone repo and create a new branch: ```$ git checkout -b name_for_new_branch```.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes