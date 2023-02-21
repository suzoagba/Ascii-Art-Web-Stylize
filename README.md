# Ascii-Art-Web-Stylize

## Usage

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface). Ascii-art-web-stylize adds designed interface.
The webpage allows turning user input into ascii-art and offers the use of three different font styles:

* shadow
* standard
* thinkertoy

## How To Run

1. run ```go run .```
Server running on localhost:8080

2. go to ```localhost:8080```

## Implementation

GET /: Sends HTML response, the main page.

POST /: that sends data to Go server (text and a banner) and ascii-art

Displaying data from the server using go templates. Server is built using Go HTTP handlers.

## Authors

Samuel Uzoagba.
Jeremiah Bakare.
Anton Urban.
