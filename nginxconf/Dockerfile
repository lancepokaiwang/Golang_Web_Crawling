FROM nginx:1.21.3


RUN ln -sf /dev/stdout /var/log/nginx/access.log \
 && ln -sf /dev/stderr /var/log/nginx/error.log
 

COPY ./grpc_proxy.conf.template /etc/nginx/conf.d/grpc_proxy.conf.template
COPY ./nginx.conf /etc/nginx/nginx.conf
RUN rm /etc/nginx/conf.d/default.conf
