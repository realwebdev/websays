# websays
RESTful JSON API

In the assessment task:

This is a time tracking app.
Multiple users can create project with unique id's can spend their time on it and sign out. In the end the app gives you the summary of time that you spend on each projects.
It uses:
1. gorm as relational mapper
2. postgres
3. jwt token for authentication
4. gin as framework for api development

As there were multiple requirements like

Creating CRUD for entity Category using memory persistance.
Create CRUD for entity Article using text file persistance.

Due to time constraint There are lot of functionalities to cover but following areas covered somewhat:

1. go unit test: 1 example not covered the whole app.
2. make file commands
3. github workflows
4. secret management using .env file
5. tried to show isolation in apps. so there is less and less dependencies
