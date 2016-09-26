# go-chat-server
Simple chat server project

The goal of this project is to familiarize myself with Go. I was inspired by this [quora question](https://www.quora.com/What-is-a-good-golang-project-to-work-on-for-beginners-that-can-become-a-decent-project-as-they-get-better-at-golang)

I've decided to follow the approach laid out by this [class](http://pirate.shu.edu/~wachsmut/Teaching/CSAS2214/Virtual/Lectures/chat-client-server.html)

Basically, step-by-step process:

1. A simple server that will accept a single client connection and display everything the client says on the screen. If the client user types ".bye", the client and the server will both quit.
2. A server as before, but this time it will remain 'open' for additional connection once a client has quit. The server can handle at most one connection at a time.
3. A server as before, but this time it can handle multiple clients simultaneously. The output from all connected clients will appear on the server's screen.
4. A server as before, but this time it sends all text received from any of the connected clients to all clients. This means that the server has to receive and send, and the client has to send as well as receive
5. Wrapping the client from step 4 into a very simple GUI interface but not changing the functionality of either server or client. The client is implemented as an Applet, but a Frame would have worked just as well (for a stand-alone program).
