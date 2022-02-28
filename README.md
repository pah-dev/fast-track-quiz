# FAST TRACK QUIZ

### Installation Instructions:

1. go install in project cli to install quiz cli.
2. go run . in project api to start the API.

### As suggested, the databases are in memory and communication with the API is through the CLI.

### Game Instructions:

- To start the quiz: quiz init
- To get a question: quiz qn
- To answer a question: quiz an -i X -a Y where X is the question ID and Y is the answer number.
- To finish the game and see results: quiz end

### Things to improve:

- Better logic when saving the pending question, it could be done in the configuration file instead of having to make an API request.
- Control of the existence of the Player ID when the API server is down.
- Make the question options more flexible, not set at four.
