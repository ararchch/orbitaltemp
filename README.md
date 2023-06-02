# README
## Tech Design:
https://www.notion.so/Tech-Design-for-API-Gateway-70b9ec603f56424d80603b394bd3c3ec 
---------------------------------------------------------------------------------------
## Current Proof of Concept:

### HERTZ
**Before you run:** Ensure you have installed and setup hertz and golang. <br>
**Step 1:** In your terminal cd into 'hertz' directory<br>
**Step 2:** enter: go run main.go<br>
**Step 3:** Open a new terminal and enter: curl http://127.0.0.1:8888/ping. <br>
**Response:** You should recieve a 200 (success) response and the message 'pong'. <br>

### KITEX
**Before you run:** Ensure you have installed and setup kitex, thriftgo, and golang. <br>
**Step 1:** In your terminal cd into 'example' directory<br>
**Step 2:** enter: go run main.go handler.go<br>
**Step 3:** The server should be up and running. Then cd into the 'client' directory.<br> 
**Step 4:** enter: go run main.go <br>
**Response:** You should recieve a response with the message 'hello'. You can edit the message in the main.go file (inside client dir) and resend. The response will echo your message. <br>
