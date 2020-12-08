
# Messaging App
## Features
- ~~Built in google authentication for login~~ (could not get this working ☹)
	- note- you will see this in the diagrams below, but it is not currently functional. It has been left in to show the intention of how the login process should have worked
- Modern programming teqniques: GoLang server (new language for all members)
	- server is API accessible
- Persistent storage using MySQL
- Ability to send and recieve messages in real time
- Automatic build and deployment with docker-compose
	- ability to set config file per environment
	- includes MySQL dependency running in a separate container (with a default DB creation script) or AWS RDS MySQL server can can connected to

## Note
- As we could not get authentication working, we have created a log in short cut for you to use to test the app. When you follow the deploy steps below, you will be taken to a "login" that simpole asks for your userid. We have prepopulated some users in our user table that you can use to test our application

<INSERT TABLE IMAGE HERE>

- The initial architectural plan was to have React, websockets or socket.io, google api authentication, docker, and deploy to a free AWS tier. Unfortunately, we found that we tried to  do too many new things. All of these technologies were new to us. We got all of these features almost working at certain points, but had to move on from them in order to produce a working application.

##  Build and Deploy
### Backend: Golang server and MySQL setup
To build and deploy the Go server, we utilize docker. 

Installation Prerequisites:
- docker
- docker-compose

After cloning the branch, navigate to the backend folder: We have 2 options for building and deploying.
1. With a local MySQL docker container
- docker-compose build --build-arg CONFIG=local
- docker-compose up
	
2. Connect  to the AWS RDS MySQL server
- docker-compose build --build-arg CONFIG=aws
- docker-compose up
	
The golang server should now be running on port 8080

### FrontEnd

To deploy the bootstrab front end, we utilize http-serve (ideally, this would run on an S3 service in AWS and would not need docker)
From the top level, navigate to the frontend folder

The first time you run the project, make sure you run this command first:
- npm install http-server

After this is installed, the front end can be hosted with 
- http-server

take note of the port the server starts on and replace 8080 if necessary at the following URL.

http://localhost:8080/conversations.html

The application should now be hosted!

# Architecture
## Architecture Overview

![](Overview.PNG?raw=true)  
 
Users/Roles: 
Actor 1 can send and recieve messages from Actor 2 by using the messaging application
Actor 2 can send and recieve messages from Actor 1 by using the messaging application  

Messaging App Boostrap Application: 
The messaging app application consist of three pages, we use boostrap for the design of the UI.
The application consist of HTML and JavaScript.  

Server: Go Lang 
The server handles the logic of the application. The backend is constructed using go lang. 

Database: 
The user id will be stored to keep track messages. When the user calls for the messages it will retrieve from MySQL table called message0. 
The messages will be stored at the message0 table.  

## Component View

![](ComponentsDiagram.PNG?raw=true)

 The user will access the Messaging UI component of the application. The frontend component of the application consist of three page which is the login, conversation, and contacts. The UI layer is created by using Bootstrap.  The Web API requests will contact the go lang server. The intention is to authenicate users using Google API authentication. The implementation is present but not functional at this time. The alternative route we use is on the login page the user insert their name then hit submit to be assign a user id. The data will be retrieve and save at MySQL database.  
 
 Front End of the applciation consist of the following: 
 
- Login page 
- Conversations page
- Contacts page
 
 The backend of the application consist of the following: 
 
  - API Server Layer: 
  - If the user logins using the Google OAuthenication (not currently functional at this time), what is suppose to happen is the UserAuthenication module will call for the Google API Authenication to authorize the user. 
  - The web api request will access the module pkg and repository that contains, mysqlContactRepository, mysqlUserRepository, and mysqlMessageAndChatRepository. The data will format from the data models module.  
 
 Database Connection Layer: 
  - The repository will connect to the database connection layer connects to the MySQL database to handle retrieve and store data. 

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

***SIMULATE TWO SEPERATE USERS SENDING MESSAGES TO EACH OTHER***** 
Need to open two seperate browsers windows open to simulate two seperate users. 

Create New User: 

Successful login will create a new user in the database, then it will return the user id from the table. 
Every user that is in the messaging app will be assigned a user id which will track each messages. 

Conversation Page: 

The chat page will render after a successful login and you are establish user with a user ID. The current user can add a contact for another user.  User #2 can access the web app and repeat the same process, login and get assign a user id.  User #1 now can add User #2 through the contact page. Both users now can send messages to each other. The messages are sent through the Go Lang server, saves the message then store in the MySQL database. User #2 will be able to render the new messages and respond to those messages through the web app. The messages will be retrieved from the MySQL database. 
