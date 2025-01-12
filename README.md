# Discord Invite Checker

A simple Go program to check Discord invite codes. It reads a list of invite codes from a file, checks if they are claimed or unclaimed, and writes unclaimed invites to an output file.

**Note:** Proxies are necessary for this tool to function properly. You must configure a valid proxy before running the program.

## Legend

- `[-]` = taken/blacklisted invite code
- `[+]` = not taken invite code (unclaimed)

## Prerequisites

- [Go](https://golang.org/) installed on your machine.
- The [Fatih Color](https://github.com/fatih/color) package for colored output. Install it using:
  ```bash
  go get github.com/fatih/color
  ```

## Setup

1. **Clone or download the repository** containing the source code.

2. **Prepare Input File:**
   - Create a file named `in.txt` in the same directory as the Go program.
   - Add Discord invite codes (one per line) to the `in.txt` file.

3. **Configure Proxy:**
   - Open the source code file.
   - Locate the line with the `Proxy` constant:
     ```go
     const Proxy = ""
     ```
   - Replace the empty string with your proxy URL in the correct format.  
   **Proxy Format:**  
     ```http
     http://NAME:PASS@IP:PORT
     ```
     For example:
     ```go
     const Proxy = "http://username:password@127.0.0.1:8080"
     ```
   - Save the changes.

## Usage

1. **Build the program:**
   ```bash
   go build -o discord-invite-checker
   ```

2. **Run the program:**
   ```bash
   ./discord-invite-checker
   ```

3. **Check Results:**
   - The program will read invite codes from `in.txt`, check their status via the Discord API using the configured proxy, and log output to the console.
   - Unclaimed invite codes will be appended to `out.txt`.
   - Outputs:
     - `[+] invite_code > Unclaimed!` indicates the invite code is not taken.
     - `[-] invite_code > Claimed!` indicates the invite code is taken or blacklisted.

## Notes

- The program creates/overwrites `out.txt` to store unclaimed invite codes.
- Ensure you have correct proxy settings and proper internet connectivity as the program queries Discord's API through the proxy.
- Proxies are mandatory for this program to work, so make sure the `Proxy` constant is set to a valid proxy URL.
