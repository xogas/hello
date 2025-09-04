# -*- coding: utf-8 -*-
from django.http import HttpResponse


def index(request):
    return HttpResponse("Hello, world. Your blueking app is running successfully now.")
