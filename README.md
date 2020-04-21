
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

## How to Contribute
1. Clone repo and create a new branch: ```$ git checkout -b name_for_new_branch```.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes