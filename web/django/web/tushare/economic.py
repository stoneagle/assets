'''
tushare宏观经济数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取存款利率
def getDepositRate(request):
    return HttpResponse()


# 获取贷款利率
def getLoanRate(request):
    return HttpResponse()


# 获取存款准备金率
def getRRRRate(request):
    return HttpResponse()


# 获取货币供应量
def getMoneySupply(request):
    return HttpResponse()


# 获取货币供应量(年底余额)
def getMoneySupplyBal(request):
    return HttpResponse()


# 国内生产总值(季度)
def getGDPQuarter(request):
    return HttpResponse()


# 三大需求对GDP贡献
def getGDPFor(request):
    return HttpResponse()


# 三大产业对GDP拉动
def getGDPPull(request):
    return HttpResponse()


# 三大产业贡献率
def getGDPContrib(request):
    return HttpResponse()


# 居民消费价格指数
def getCPI(request):
    return HttpResponse()


# 工业品出厂价格指数
def getPPI(request):
    return HttpResponse()
