let grpc = require('grpc');
let protoLoader = require('@grpc/proto-loader');
let discount = require('../discount/discount');

const protoPath = __dirname + '/protos/discount.proto';
const host = process.env.GRPC_HOST || '0.0.0.0';
const port = process.env.GRPC_PORT || ':50051';
const address = host + port

let packageDefinition = protoLoader.loadSync(
    protoPath,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
let discountProto = grpc.loadPackageDefinition(packageDefinition).protos;
let server = new grpc.Server();

server.addService(discountProto.DiscountService.service, {
    discount: async (call, callback) => {
        let requestID = call.request.requestID;
        let productID = call.request.productID;
        let userID = call.request.userID;
        try {
            let result = await discount.getPercentage(requestID, productID, userID);
            callback(null, {
                percentage: result,
            });
        } catch(err){
            console.log("requestID: " + requestID + " |  failed to get percentage" + err);
        };
    },
});

module.exports.start = function(){
    server.bindAsync(address, grpc.ServerCredentials.createInsecure(), () => {
        server.start();
        console.log("Server running at " + port);
    });
};