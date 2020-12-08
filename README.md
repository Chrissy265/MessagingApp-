
# Messaging App
## Features
- built in google authentication for login
- Modern programming teqniques: GoLang server (new language for all members)
	- server is API accessible
- Persistent storage using MySQL
- Ability to send and recieve messages in real time
- Automatic build and deployment with docker-compose
	- ability to set config file per environment
	- includes MySQL dependency running in a separate container (with a default DB creation script) or AWS RDS MySQL server can can connected to

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

![](Overview.PNG?raw=true)



## Component View

![](ComponentsDiagram.PNG?raw=true)

## MySQL server layout

![](sqlDiagram.PNG?raw=true)

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

![](SequenceDiagram.PNG?raw=true)  

Login Page: 
The user go on the web application Messaging app. The user have to login using their username and password or use Google OAuth sign in. If the user sign in with their Google crendentials email and password.  The Google OAuth login goes through the Go Lang server, if successful then return authorization code if success. The Go Lang server will request token to the Google API authenication then request user data.  

Create New User: 

Successful login will create a new user in the database, then it will return the user id from the table. 
Every user that is in the messaging app will be assigned a user id which will track each messages. 

Conversation Page: 

The chat page will render after a successful login and you are establish user with a user ID. The current user can add a contact for another user.  User #2 can access the web app and repeat the same process, login and get assign a user id.  User #1 now can add User #2 through the contact page. Both users now can send messages to each other. The messages are sent through the Go Lang server, saves the message then store in the MySQL database. User #2 will be able to render the new messages and respond to those messages through the web app. The messages will be retrieved from the MySQL database. 




 




