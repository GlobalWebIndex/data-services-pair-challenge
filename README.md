# Interview Code Challenge

Welcome to the interview code challenge for potential candidates! In this challenge, you will work on a simple CRUD application that manages audiences and their expressions. The challenge is designed to be run as a pair-programming session where we will be looking over your shoulder to see how you work and collaborate.

## Task Options

### Option 1: Audience and Expression JSON Marshaling

**Issue:** The current implementation of JSON unmarshaling for the `Expression` struct is causing the tests to fail. Your task is to fix the unmarshaling implementation so that the tests pass.

**Instructions:** We will work together to update the `domain/audience.go` file in the `option_1` branch to modify the JSON unmarshaler implementation of `Expression` structs, ensuring correct decoding of the JSON data. We will use the `json.Unmarshal` function from the Go standard library to handle the decoding.

#### Steps (Option 1):

1. Clone this repository to your local development environment.
2. Checkout the `option_1` branch.
3. Run the unit tests.

### Option 2: Correctly Returning ID after Audience Creation

**Issue:** When creating an audience, the code in the `Create` method of the `audienceRepository` does not return the inserted ID in the response. Your task is to modify the `Create` method to correctly return the inserted ID after an audience is created.

**Instructions:** We will work together to update the `db/audience_repository.go` file in the `option_2` branch to modify the `Create` method implementation. After inserting an audience record into the database, we will fetch the inserted ID and return it in the response. We will ensure that the `AudienceID` type is correctly used for the ID field.

#### Steps (Option 2):

1. Clone this repository to your local development environment.
2. Checkout the `option_2` branch.
3. Run the application with Docker Compose.
4. Send a POST request using curl to create an audience with expression data:
5. `curl -X POST http://localhost:8080/audience --data '{"name":"my audience", "expression": {"question_code":"q1"}}'`

The response should be: `{"id":"1","name":"my audience","expression":{"question_code":"q1"}}`.

## Note

- The code provided in this challenge is intentionally not perfect. It may contain bugs, inefficiencies, or missing functionality. During our pair-programming session, feel free to talk about any issues you encounter and provide suggestions for better solutions.

We look forward to working with you in the pair-programming session and learning more about your coding skills and collaboration approach!
