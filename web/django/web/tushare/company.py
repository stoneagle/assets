'''
tushare基本面数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取沪深上市公司基本情况
def getStockBasics(request):
    return HttpResponse()


# 按年度、季度获取业绩报表数据
def getReportData(request):
    return HttpResponse()


# 按年度、季度获取盈利能力数据
def getProfitData(request):
    return HttpResponse()


# 按年度、季度获取营运能力数据
def getOperationData(request):
    return HttpResponse()


# 按年度、季度获取成长能力数据
def getGrowthData(request):
    return HttpResponse()


# 按年度、季度获取偿债能力数据
def getDebtpayingData(request):
    return HttpResponse()


# 按年度、季度获取现金流量数据
def getCashflowData(request):
    return HttpResponse()
