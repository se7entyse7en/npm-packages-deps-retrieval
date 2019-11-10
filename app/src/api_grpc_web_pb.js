/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */


const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.api = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.DependenciesServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.DependenciesServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.DependenciesRequest,
 *   !proto.api.DependenciesResponse>}
 */
const methodDescriptor_DependenciesService_GetDependencies = new grpc.web.MethodDescriptor(
  '/api.DependenciesService/GetDependencies',
  grpc.web.MethodType.UNARY,
  proto.api.DependenciesRequest,
  proto.api.DependenciesResponse,
  /**
   * @param {!proto.api.DependenciesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.DependenciesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.DependenciesRequest,
 *   !proto.api.DependenciesResponse>}
 */
const methodInfo_DependenciesService_GetDependencies = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.DependenciesResponse,
  /**
   * @param {!proto.api.DependenciesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.DependenciesResponse.deserializeBinary
);


/**
 * @param {!proto.api.DependenciesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.DependenciesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.DependenciesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.DependenciesServiceClient.prototype.getDependencies =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.DependenciesService/GetDependencies',
      request,
      metadata || {},
      methodDescriptor_DependenciesService_GetDependencies,
      callback);
};


/**
 * @param {!proto.api.DependenciesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.DependenciesResponse>}
 *     A native promise that resolves to the response
 */
proto.api.DependenciesServicePromiseClient.prototype.getDependencies =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.DependenciesService/GetDependencies',
      request,
      metadata || {},
      methodDescriptor_DependenciesService_GetDependencies);
};


module.exports = proto.api;

