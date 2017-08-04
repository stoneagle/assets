'''
tushare银行间同业拆放利率
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取银行间同业拆放利率数据，目前只提供2006年以来的数据
def getShiborData(request):
    return HttpResponse()


# 获取银行间报价数据，目前只提供2006年以来的数据
def getShiborQuoteData(request):
    return HttpResponse()


# 获取Shibor均值数据，目前只提供2006年以来的数据
def getShiborMaData(request):
    return HttpResponse()


# 获取贷款基础利率（LPR）数据，目前只提供2013年以来的数据
def getLprData(request):
    return HttpResponse()


# 获取贷款基础利率均值数据，目前只提供2013年以来的数据
def getLprMaData(request):
    return HttpResponse()
