import grpc


host = 'localhost'
server_port = 9009

grpc_channel = grpc.insecure_channel(f'{host}:{server_port}')
