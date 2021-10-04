/**
 * @fileoverview gRPC-Web generated client stub for product
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.product = require('./product_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.product.ProductServiceClient =
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
proto.product.ProductServicePromiseClient =
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
 *   !proto.product.ProductRequest,
 *   !proto.product.ProductResponse>}
 */
const methodDescriptor_ProductService_Query = new grpc.web.MethodDescriptor(
  '/product.ProductService/Query',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.product.ProductRequest,
  proto.product.ProductResponse,
  /**
   * @param {!proto.product.ProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.ProductResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.product.ProductRequest,
 *   !proto.product.ProductResponse>}
 */
const methodInfo_ProductService_Query = new grpc.web.AbstractClientBase.MethodInfo(
  proto.product.ProductResponse,
  /**
   * @param {!proto.product.ProductRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.ProductResponse.deserializeBinary
);


/**
 * @param {!proto.product.ProductRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.product.ProductResponse>}
 *     The XHR Node Readable Stream
 */
proto.product.ProductServiceClient.prototype.query =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/product.ProductService/Query',
      request,
      metadata || {},
      methodDescriptor_ProductService_Query);
};


/**
 * @param {!proto.product.ProductRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.product.ProductResponse>}
 *     The XHR Node Readable Stream
 */
proto.product.ProductServicePromiseClient.prototype.query =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/product.ProductService/Query',
      request,
      metadata || {},
      methodDescriptor_ProductService_Query);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.product.HelloRequest,
 *   !proto.product.HelloReply>}
 */
const methodDescriptor_ProductService_SayHello = new grpc.web.MethodDescriptor(
  '/product.ProductService/SayHello',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.product.HelloRequest,
  proto.product.HelloReply,
  /**
   * @param {!proto.product.HelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.HelloReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.product.HelloRequest,
 *   !proto.product.HelloReply>}
 */
const methodInfo_ProductService_SayHello = new grpc.web.AbstractClientBase.MethodInfo(
  proto.product.HelloReply,
  /**
   * @param {!proto.product.HelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.product.HelloReply.deserializeBinary
);


/**
 * @param {!proto.product.HelloRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.product.HelloReply>}
 *     The XHR Node Readable Stream
 */
proto.product.ProductServiceClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/product.ProductService/SayHello',
      request,
      metadata || {},
      methodDescriptor_ProductService_SayHello);
};


/**
 * @param {!proto.product.HelloRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.product.HelloReply>}
 *     The XHR Node Readable Stream
 */
proto.product.ProductServicePromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/product.ProductService/SayHello',
      request,
      metadata || {},
      methodDescriptor_ProductService_SayHello);
};


module.exports = proto.product;

