# Klaras Book Vault

📚 **Klaras Book Vault** is a minimalist, desktop-based library management application designed to help users organize and track their personal book collections. Built with Go, it utilizes a local binary storage system and a simple GUI framework for an efficient and user-friendly experience.

## Features

-   **Local Binary Storage**: Stores book data in a local binary file (`books.klara`) for quick access and offline use.
-   **Simple GUI**: Employs a straightforward graphical interface for ease of use.
-   **Book Management**: Add, edit, and view book entries with essential details.
-   **ISBN Handling**: Includes functionality to manage and validate ISBN numbers.([GitHub][1], [GitHub][2])

## Building

### Prerequisites

-   Go installed and configure on your machine

### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/davidstyrbjorn/klaras-book-vault.git
    ```

2\. Navigate to the project directory:

```bash
cd klaras-book-vault
```

3\. Build the application:

```bash
go build -o klaras-book-vault
```

4\. Run the application:([GitHub][1])

```bash
./klaras-book-vault
```

## Usage

Upon launching, the application presents a user-friendly interface to manage your book collection.([GitHub][1])

## File Structure

-   `main.go`: Entry point of the application.
-   `addBookView.go`, `editBookView.go`, `bookshelfView.go`, `homeView.go`: GUI components for different views.
-   `binary.go`: Handles binary file operations for storing book data.
-   `isbn.go`: Functions related to ISBN processing.
-   `state.go`: Manages application state.
-   `constants.go`: Defines constant values used across the application.
-   `books.klara`: Binary file where book data is stored.([GitHub][4], [GitHub][2])

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/davidstyrbjorn/klaras-book-vault/blob/main/LICENSE) file for details.
