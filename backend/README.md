
# Messaging App
## Features
- built in google authentication for login
- Modern programming teqniques: GoLang server (new language for all members)
	- server is API accessible
- Persistent storage using MySQL
- Ability to send and recieve messages in real time
- Automatic build and deployment with docker-compose
	-ability to set config file per environment
	includes MySQL and redis dependencies each running in a separate container

##  Build and Deploy
### Backend: Golang server and MySQL setup
To build and deploy the Go server, we utilize docker. 

Prerequisites:
- docker
- docker-compos

After cloning the "backend" branch. We have 2 options for building and deploying.
1. With a local MySQL docker container
- docker-compose build --build-arg CONFIG=local
- docker-compose up
	
2. Connect  to the AWS RDS MySQL server
- docker-compose build --build-arg CONFIG=aws
- docker-compose up -d backend
	
The golang server should now be running on port 8080
### FrontEnd
????????????

# Architecture
## Architecture Overview

## Component View
----insert image here

## MySQL server layout

----insert image here

red nodes are foreign keys
yellow keys are primary keys

- user
	- stores all user data obtained from the google auth API
- userContacts
	- stores users and contacts in pairs. Users can have as many contacts as they want
- userpreferences
	- allows for a future where users would be able to customize the app. 
- userchatpreferences
	- allows for further customization of chats for a specific user. For now, this table functions as a look up table for what chats a user is part of. A chat has no limit on the number of users it can be created with
- chat 
	- includes basic chat information as well as a column called messageTable. 
	- Storing individual messages in 1 SQL table can create a very large table. A way around this is to have multiple messages tables and a hash function that would  assign each new chat to a message table depending on criteria such as the current sized of each messaging table. All Messages for a chat would then be placed in the table that corresponds to the messageTable (0 would correspond to messages0). 
	- Current implementation has all messages being stored in messages0
- messages0- stores all user messages

## Sequence Diagram
----insert image here
