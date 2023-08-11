const path = require("path")
const grpc = require("@grpc/grpc-js")
const protoLoader = require("@grpc/proto-loader")

const proto = protoLoader.loadSync(path.join(__dirname, "..", "posts_service.proto"))
const definition = grpc.loadPackageDefinition(proto)

const postList = [
    {id: 1, title: "Title 1", text: "Text 1"},
    {id: 2, title: "Title 2", text: "Text 2"}
]

const getPosts = (call, callback) => {
    callback(null, { posts : postList })
}

const serverURL = "localhost:10000"
const server = new grpc.Server()

server.addService(definition.PostService.service, { getPosts })
server.bindAsync(serverURL, grpc.ServerCredentials.createInsecure(), port => {
    console.log(`Server running on ${serverURL}`)
    server.start()
})