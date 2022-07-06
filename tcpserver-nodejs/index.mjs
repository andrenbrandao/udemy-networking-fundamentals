import net from 'net';

const server = net.createServer(socket => {
  console.log(`TCP Handshake successful with ${socket.remoteAddress}:${socket.remotePort}`)

  socket.write('Hello client!\n')
  socket.on('data', data => {
    console.log(`Received data: ${data.toString()}`)
  })
})

server.listen(5500, '127.0.0.1')