let server = require('../grpc')
require('log-timestamp')(function() { return new Date().toISOString() + ' | %s' });;

module.exports.app = function() { 
    server.start()
}