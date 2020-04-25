
# kont


## Overview

kont is the application developed by Go as open-source. It provides you ability that assigns as many reviewers as given
to pull/merge request to the protected branch by based some algorithms. It aims to improve code quality in code-review process of 5G project at Havelsan.

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

#### Sample Bitbucket Repository in kont
kont needs some information about the remote repository for starting to observe pull/merge requests. So, a repository record
must consist of below parameters:

```json

{
	"host":"http://10.120.0.145:7990",
	"token": "Bearer AB...",
	"projectName": "BESG",
	"name": "core-network",        
	"developmentBranch": "develop",
    "provider": "BITBUCKET",
    "defaultComment": "Merhaba @{{name}} \n ## **Reviewer koda bakmadan evvel, kendin bir kez daha review etmeye ne dersin? Eminim ELF gözlerin bişeyler görecektir.** \n Bunun için aşağıdaki maddeleri kontrol edebilirsin \n * Reformat \n * SonarLint \n * Analyze -> Inspect code \n * Mimari olarak düzgün mü? (heryer heryerde olmasın lütfen) \n * Test isimlerini daha anlaşılır yapabilirsin(given-when-then) \n * Fazla test, göz çıkarmaz \n\n Kolay Gelsin Hacım :) ",
    "stages": [
		{
		"name":"Stage1",
		"policy":"RANDOMINAVAILABLE",
		"reviewers":[
			{
				"priority": 7,
				"user":{
			             "name":"ataday",
                         "displayName": "Alper TADAY"
				}
			},
	
	    	{
	    	   "priority": 2,
			   "user":{
						"name":"baydogdu",
						"displayName": "Büşra Cennet AYDOĞDU"
				}
		
			},
			{
				"priority": 6,
				"user": {
						"name":"mcil",
				        "displayName": "Mevlüt Mert ÇİL"
				}
			},
			{
				"priority": 5,
				"user": {
						"name":"nkoc",
				        "displayName": "Nilgün KOÇ"
				}
			},
			{
				"priority": 4,
				"user": {
						"name":"tunsal",
				        "displayName": "Tuğba ÜNSAL"
				}
			}
		
		]
		
	  },
	  	{
		"name":"Stage2",
		"policy":"RANDOMINAVAILABLE",
		"reviewers":[
			{
				"priority": 1,
				"user":{
            			"name": "huseyiny",
        				 "displayName": "Hüseyin YILDIRIM"
				}
			},
	    	{
	    	   "priority": 2,
			   "user":{
        				 "name": "eunal",
            			"displayName": "Eylül ÜNAL"
				}
		
			},
	    	{
	    	   "priority": 3,
			   "user":{
        				 "name": "hsumer",
            			"displayName": "Halil İbrahim SÜMER"
				}
		
			}
		]
		
	  },
	  {
		"name":"Stage3",
		"policy":"RANDOMINAVAILABLE",
		"reviewers":[
			{
				"priority": 1,
				"user":{
            			"name": "edincer",
            			"displayName": "Emre DİNÇER"
				}
			}
		]
		
	  }
	]
}
```

**host**: Scheme and host of the provider for your Git repository.

**token**: Personal access bearer token which has necessary permissions(repository write/admin)

**projectName**: Key of the project which contains the repository

**name**: Name of the remote repository

**developmentBranch**: Development branch of the remote repository for observing pull/merge requests to this branch.

**provider**: String Upper-case provider name, any of ("GITHUB", "BITBUCKET", "GITLAB")

**defaultComment**(Optional): Default comment text for each new pull/merge request(Markdown syntax in string)

**stage**: Each stage consists of name, policy and list of reviewer parameters. 
kont selects a *available reviewer* in the given reviewer list based *policy type*.

* ```availability``` is that, The reviewer has not been assigned to any pull/merge request or 
the reviewer has approved all pull/merge requests which he/she had been assigned.

* ```policy``` type can be any of ("RANDOMINAVAILABLE","BYPRIORITYINAVAILABLE"): 
    * ```RANDOMINAVAILABLE```: Selects a random reviewer in available reviewers in the stage.
    * ```BYPRIORITYINAVAILABLE```: Sorts available reviewers and select first which has high priority value.



## Notes
1. kont just has integration with **Bitbucket-Server(based on Rest Api v1)** currently

## How to Contribute
1. Clone repo and create a new branch: ```$ git checkout -b name_for_new_branch```.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes