# Key-Haven

Key-Haven is a secure and efficient password manager designed to store and manage your credentials safely. It provides encryption, easy access, and a user-friendly interface to help you manage your passwords securely.

## Requirements
- Go 1.24 or higher
- Docker and Docker Compose for running the dependent services

## Features
- **Secure Storage**: Uses strong encryption to store passwords safely.
- **User Authentication**: Ensures only authorized users can access their vault.
- **Cross-Platform**: Accessible on multiple devices.
- **Easy Retrieval**: Quick and efficient password retrieval.


## Installation
To install Key-Haven, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/RaposoG/Key-Haven.git
   ```
2. Navigate into the project directory:
   ```sh
   cd Key-Haven
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Start services required by the application:
   ```sh
   docker-compose up -d
   ```
5. Build the application:
   ```sh
   mkdir -p ./bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api ./cmd/server
   ```
6. Run the application:
   ```sh
   ./bin/api
   ```

## Usage
1. Start the application server
2. Access the API endpoints for password management
3. Use the encryption and decryption features for secure password storage
4. Authenticate with your master password for added security

## Security Measures
- **End-to-End Encryption**: All passwords are encrypted using AES-GCM
- **Master Password Protection**: Your vault is protected by a strong master password
- **Secure Storage**: Sensitive data is never stored in plain text
- **Hash Protection**: User authentication uses bcrypt for password hashing

## Roadmap
- [ ] Add This
- [ ] Add That
- [ ] Secure password sharing

## Contributing
We welcome contributions! Please submit a pull request or open an issue to discuss your ideas.

## License
This project is licensed under the MIT License. See the LICENSE file for details.