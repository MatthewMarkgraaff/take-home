# Programming take-home project
The goal of this project is for you, the candidate, to demonstrate your programming ability - while being comfortable working in your own environment.

### The Assignment
Spend no more than two hours building the following:  
A simple Go webserver that serves a single HTML page at it's root url.  
The homepage should have some simple markup explaining the project, a single text input field & submit button.  
The user should be prompted to enter a github username into the text field and submit the form.  
The server should receive this request and respond by rendering a list / table containing all of the user's public Github repositories. Feel free to show information IE stars, repository name, links to each repository.  
You do not need to persist any data for this project.

To acheive this, you will need to integrate with the GitHub REST API (https://docs.github.com/en/rest/reference/repos#list-repositories-for-a-user).  
Use whatever packages, frameworks & tools you require to achieve the above. Aim for simplicity, conciceness and of course, working code.

### Project Requirements
The following must be included in each submission:

1. An amendment to this README containing instructions on how to test and run your submission
2. You should work on your own git branch and submit a pull request to the `main` branch when you're ready for your work to be reviewed

P.S Writing unit tests are highly encouraged
