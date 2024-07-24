# Checkit Enron Email Search App

## Introduction

Welcome to the Enron Email Search App repository! This project provides a simple and intuitive way to explore the Enron Email Dataset, offering users the ability to browse and search through the extensive collection of emails from the Enron Corporation.

## About the Project

This app was developed to allow users to delve into the Enron emails, gaining insights into corporate communication and operations. The Enron Email Dataset was released to the public during the legal investigation following the Enron scandal, making it a valuable resource for research and education.

## Project Structure

The project is divided into three main parts:

1. **Frontend:** Built with Vue.js and TailwindCSS, the frontend provides a clean and responsive user interface for browsing and searching emails.
2. **Backend:** Developed using Go and Chi, the backend handles API requests and manages communication between the frontend and the ZincSearch database.
3. **Indexer:** Also written in Go, the indexer processes and indexes the Enron email data into ZincSearch.

## Features

- **Simple User Interface:** An easy-to-use UI for browsing the emails.
- **Search Engine:** A powerful search engine to help users find specific concepts and topics of interest.

## Getting Started

### Prerequisites

To run this project, you will need:

- Go 1.18 or higher
- ZincSearch
- Other dependencies listed in `go.mod` and `package.json`
- Node.js and npm
- Vue
- Docker
- Chi

### Installation

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/davidsequera/checkit.git
    ```
2. Navigate to the project directory:
    ```bash
    cd checkit
    ```

#### Backend

3. Install the required Go modules:
    ```bash
    cd backend
    go mod tidy
    ```

4. Set up ZincSearch:
    - A simple docker container.
    - Start the ZincSearch service.

5. Index the Enron email data into ZincSearch:
    ```bash
    go run cmd/indexer/main.go
    ```

6. Run the backend server:
    ```bash
    go run cmd/server/main.go
    ```

#### Frontend

7. Navigate to the frontend directory:
    ```bash
    cd ../frontend
    ```

8. Install the required npm packages:
    ```bash
    npm install
    ```

9. Run the frontend development server:
    ```bash
    npm run serve
    ```

### Usage

Once the app is running, you can access it via your web browser at `http://127.0.0.1:8080` for the frontend. Use the UI to browse emails or utilize the search feature to find specific content.

## Contributing

We welcome contributions! If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

1. Fork the repository.
2. Create a new branch:
    ```bash
    git checkout -b feature/YourFeature
    ```
3. Make your changes and commit them:
    ```bash
    git commit -m 'Add some feature'
    ```
4. Push to the branch:
    ```bash
    git push origin feature/YourFeature
    ```
5. Create a new pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Acknowledgements

- Enron Email Dataset provided by [Carnegie Mellon University](https://www.cs.cmu.edu/~./enron/).
- Inspiration and guidance from the open-source community.

## Contact

For any questions or inquiries, please open an issue in the repository or contact us at your-email@example.com.

---

We hope you find this app useful for exploring the Enron emails. Happy searching!