/* eslint-disable no-console */
import { HelloRequest, ProductRequest } from './product_pb.js';
import { ProductServiceClient } from './product_grpc_web_pb.js';

const client = new ProductServiceClient('http://localhost:8080', null, null);

document.getElementById('button').addEventListener('click', () => {
  const req = new HelloRequest();
  const name = document.getElementById('input').value || 'world';
  req.setName(name);

  const stream = client.sayHello(req, {});
  stream.on('data', res => {
    document.getElementById('count_result').textContent = res.getMessage();
  });
  stream.on('status', status => {
    console.log(status);
    console.log(status.details);
    console.log(status.metadata);
  });
  stream.on('error', err => {
    console.log(err);
  });
  stream.on('end', () => {
    console.log('end');
  });

});

document.getElementById('count').addEventListener('click', () => {
  const req = new ProductRequest();
  req.setN(10);
  const stream = client.countStream(req, {});
  stream.on('data', res => {
    document.getElementById('count_result').textContent = res.getCount();
  });
  stream.on('status', status => {
    console.log(status);
  });
  stream.on('error', err => {
    console.log(err);
  });
  stream.on('end', () => {
    console.log('end');
  });
});
