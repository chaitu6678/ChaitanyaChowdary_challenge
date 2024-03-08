# Validating Credit Card Numbers

This Go program validates credit card numbers based on certain criteria. It checks whether the given credit card numbers are valid according to the following rules:

1. It must start with a 4, 5, or 6.
2. It must have exactly 16 digits.
3. It must only consist of digits (0-9).
4. It may have digits in groups of 4, separated by one hyphen "-".

## How to Run

To run the program, make sure you have Go installed on your system. Then, follow these steps:

1. Clone or download the repository to your local machine.
2. Navigate to the directory containing the `main.go` file.
3. Open a terminal window and run the following command:

   ```bash
   go run main.go

Give Input
6
4123456789123456
5123-4567-8912-3456
61234-567-8912-3456
4123356789123456
5133-3367-8912-3456
5123 - 3567 - 8912 - 3456

Output:
Valid
Valid
Invalid
Valid
Invalid
Invalid
