from rest_framework import exceptions, status, viewsets
from rest_framework.decorators import action
from userservice.service.auth import AuthService
from rest_framework.response import Response


class AuthViewSet(viewsets.ViewSet):
    @action(detail=False, methods=["POST"], permission_classes=())
    def login(self, request):
        data = request.data.copy()
        username = data.get("username", "")
        password = data.get("password", "")
        res = AuthService.login(username=username, password=password)
        return Response(dict(success=res.success, message=res.message))

    @action(detail=False, methods=["POST"], permission_classes=())
    def register(self, request):
        data = request.data.copy()
        username = data.get("username", "")
        password = data.get("password", "")
        res = AuthService.register(username=username, password=password)

        return Response(dict(success=res.success, message=res.message))
