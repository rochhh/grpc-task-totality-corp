Getting Started

To run this application, you can choose one of the following methods

Running Locally

1. Make sure you have Go installed on your system.
2. Clone this repository

git clone https://github.com/rochhh/grpc-task-totality-corp.git
cd ./


3. Build and run the application 

go build ./usermgmt-server 
go build ./usermgmt-client

./usermgmt-server
./usermgmt-client 


Running with Docker

1. docker build -t grpc-usermgmt .
2. docker run -p 50051:50051 grpc-usermgmt

