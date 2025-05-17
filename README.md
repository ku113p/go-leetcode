# LeetCode Solutions in Go

This repository contains solutions to various LeetCode problems implemented in Go. Each problem is organized into its own directory under the `exercises/` folder, with separate files for the solution and its corresponding tests.

## How to Run

1. Clone the repository:
   ```sh
   git clone https://github.com/august346/leetcode-go.git
   cd leetcode-go
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run a specific test:
   ```sh
   go test ./exercises/ex926
   ```

4. Run all tests:
   ```sh
   go test ./...
   ```

## Utilities

The `utils/` folder contains helper functions:
- `InTime`: Ensures a test completes within a specified duration.
- `ReadStringFromFile`: Reads the content of a file as a string.

## Example Problems

### Problem 926: Flip String to Monotone Increasing
- **Solution:** [exercises/ex926/solution.go](exercises/ex926/solution.go)

### Problem 904: Fruit Into Baskets
- **Solution:** [exercises/ex904/solution.go](exercises/ex904/solution.go)

### Problem 838: Push Dominoes
- **Solution:** [exercises/ex838/solution.go](exercises/ex838/solution.go)

### Problem 153: Find Minimum in Rotated Sorted Array
- **Solution:** [exercises/ex153/solution.go](exercises/ex153/solution.go)

### Problem 1560: Most Visited Sector in a Circular Track
- **Solution:** [exercises/ex1560/solution.go](exercises/ex1560/solution.go)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.