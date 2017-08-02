from django.http import HttpResponse
import tushare as ts


def hello(request):
    return HttpResponse("Hello world!" + ts.__version__)
