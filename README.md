# rest-api-quiz
A small project creating a quiz game to learn how to create a REST API.
The client will be a CLI that uses Cobra. 

## Usage
The RESTful API can be run with the usual go run command, in the server directory.

```bash
go run server.go
```

The client CLI can be built by running go build, in the client directory.

```bash
go build .
```

## Commands

### play
Starts a new round of the quiz. 

Flags:
* -u, --username string		Set your username to automatically submit your answers after finishing the quiz

```bash
quiz play -u name
```

### new
Adds a new question to the quiz, requires all of the flags to be filled out. 

Flags:
* -q, --question string		The question
* -1, --answer1 string		Answer alternative 1
* -X, --answerX string		Answer alternative X
* -2, --answer2 string		Answer alternative 2
* -c, --correct string		The correct answer for the question (1, X, or 2)

```bash
quiz new -q "Question?" -1 "Yes" -X "Maybe" -2 "No" -c 1
```

### del
Deletes a quiz question by ID.

```bash
quiz del 3
```
This will delete question with ID 3.