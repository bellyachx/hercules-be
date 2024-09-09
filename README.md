
# Hercules Exercise Service

Hercules Exercise Service is a gRPC-based service that allows you to create and manage workout exercises. This service is part of the Hercules system and handles exercise-related operations such as creation and retrieval of exercises.

## Features

- **Create Exercise**: Allows users to create a new exercise by specifying its name.
- **Get Exercise** (future implementation): Retrieve an exercise by its ID.

## gRPC Endpoints

### 1. CreateExercise

Creates a new exercise and returns a confirmation message.

- **RPC**: `CreateExercise(CreateExerciseRequest) returns (ExerciseCreatedResponse)`
- **Request**:
  ```proto
  message CreateExerciseRequest {
    string name = 1; // Name of the exercise
  }
  ```
- **Response**:
  ```proto
  message ExerciseCreatedResponse {
    string message = 1; // Confirmation message
  }
  ```

### 2. GetExercise (Coming Soon)

Retrieve details of an exercise by its ID.

## Project Structure

- **/cmd**: Contains the main entry point for running the service.
- **/internal**: Core business logic, including exercise service and gRPC server implementation.
- **/api**: Protocol Buffers definition files for gRPC.

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/yourrepo/hercules-exercise-service.git
   ```

2. Compile protobufs:
   ```bash
   make generate
   ```

3. Run the service:
   ```bash
   make run
   ```

The service will start on port `8080` by default.

## Dependencies

- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

## License

This project is licensed under the MIT License.
