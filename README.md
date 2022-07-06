# Fundamentals of Networking for Effective Backend Design

Exercises and code from Hussein Nasser's Course [Fundamentals of Networking for Effective Backend Design](https://www.udemy.com/course/fundamentals-of-networking-for-effective-backend-design).

## How to Run

Install the [latest Go version](https://go.dev/) on your computer if it is not already installed.
Also, if running the Nodejs servers, install the [LTS Version](https://nodejs.org/en/).

Then, go to the `udpserver` or `tcpserver` respective directories and execute the following command:

Go:

```bash
go run main.go
```

Nodejs:

```bash
node index.mjs
```

In your terminal, use the Netcat utility to make the requests.

```bash
# TCP
nc -t 127.0.0.1 5500

# UDP
nc -u 127.0.0.1 5500
```

## License

[MIT](LICENSE) © André Brandão
