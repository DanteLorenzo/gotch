<p align="center">
  <img src="media/logo.png" width="200" />
</p>

# Gotch - Binary File Patching Tool

## Overview

This Go application is designed to automate the process of patching binary files using predefined patterns. The tool reads patterns from a JSON file and applies them to modify an original executable file, saving the modified version and preserving the original.

## Features

- **Automatic File Handling**: The original file is renamed and preserved, while the modified file is created with the specified changes.
- **Pattern Matching and Replacement**: Reads binary data, finds specified patterns, and replaces them with new ones.
- **Detailed Logging**: Color-coded logging for easy identification of information, success, errors, and timing details.
- **Performance Tracking**: Logs the start and end times of the patching process, as well as the total duration.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/DanteLorenzo/gotch.git
   ```

2. Navigate to the project directory:
   ```bash
   cd gotch
   ```

3. Build the application:
   ```bash
   go build .
   ```

## Usage

1. Ensure that you have a `patterns.json` file in the project directory. This file should contain the patterns you want to find and replace.

2. Run the application with the path to the executable you want to patch:
   ```bash
   ./binary-patcher /path/to/your/executable
   ```

   Example:
   ```bash
   ./binary-patcher myprogram.exe
   ```

3. The application will create a backup of the original file (e.g., `myprogram.exe.origin`) and produce the modified file with the same original name.

## Configuration

- **patterns.json**: This file contains the patterns to search for and replace in the binary file. The format should be a list of pattern objects with `OldPattern` and `NewPattern` fields.

  Example `patterns.json`:
  ```json
  [
      {
          "OldPattern": "AB CD EF 12 34",
          "NewPattern": "12 34 56 78 90"
      },
      {
          "OldPattern": "DE AD BE EF",
          "NewPattern": "00 11 22 33"
      }
  ]
  ```

## Logging

- The application uses color-coded logging for better readability:
  - **Information**: Cyan
  - **Success**: Green
  - **Errors**: Red
  - **Start/End**: Green
  - **Duration**: Green

Logs include detailed information about the patterns found in the original and patched files, any errors encountered, and the duration of the process.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---