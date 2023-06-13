# README: CloudWeGo API Gateway
## Tech Design:
#### https://drive.google.com/file/d/12YdO1ZMxGWcnbM7dVHZ8ZAMn8eJMTPbL/view?usp=drive_link 
<hr>
## Current Proof of Concept Description: <br>
### Basic echo functionality which accepts a JSON request containing a name and email address, and returns "Hello, [Name]" <br>
**How it works** <br>
1. Client sends HTTP request with a JSON body with a "name" and "email" attribute. <br>
2. Hertz server processes reqeuest and logs the "name" and "email" values. <br>
3. In the same file, a kitex client is initiated which calls the echo function on the Kitex server with the "name" value. <br>
4. Kitex server accepts request and generates response: 'Hello, [Name]', returning it to the client. <br>
5. Using Hertz, the kitex response is encoded in a JSON body and returned to the original client. <br>
<hr>
## How to test the Proof of Concept:

### Before you run:
1. Ensure you have installed and setup Hertz, Kitex, ThriftGo and Go. <br>
2. Install and setup Postman (or any similar service)

### Running the POC:

#### 1. Starting Hertz and Kitex servers
**Step 1.1:** In your terminal cd into 'example' directory<br>
**Step 1.2:** enter and execute: go run main.go handler.go<br>
**Step 1.3:** Open a new terminal and cd into the 'hz-server' directory<br>
**Step 1.4:** enter and execute: go run hz.go _(or go run .)_ <br>

#### 2. Sending HTTP requests and receiving responses on Postman
**Step 2.1:** Open postman and create a new 'POST' request<br>
              Request Address: http://localhost:8080/endpoint <br>
              Request Body: ```{"name": "ArRay", "email": "ArRay@example.com"}``` _(you can change the value of name and email as you please)_ <br>
**Step 2.2:** Send the request. you should receive the following response:<br>
              ```
                {
                    "kitexResponse": {
                        "message": "Hello, ArRay"
                    }
                }
              ```
