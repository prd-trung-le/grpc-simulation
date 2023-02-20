import grpc
from userservice.grpc_generated import auth_pb2 as pb2
from userservice.grpc_generated import auth_pb2_grpc as pb2_grpc
from config import grpc_channel


class AuthService:
    stub = pb2_grpc.AuthServiceStub(grpc_channel)

    @classmethod
    def register(cls, username, password):
        register_request = pb2.RegisterRequest(username=username,
                                               password=password)
        return cls.stub.Register(register_request)

    @classmethod
    def login(cls, username, password):
        login_request = pb2.LoginRequest(username=username, password=password)
        return cls.stub.Login(login_request)

    @classmethod
    def validate_token(cls, token):
        validate_token_request = pb2.ValidateTokenRequest(token=token)
        return cls.stub.ValidateToken(validate_token_request)
