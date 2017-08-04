'''
tushare龙虎榜数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 按日期获取历史当日上榜的个股数据，如果一个股票有多个上榜原因，则会出现该股票多条数据
def getTopList(request):
    return HttpResponse()


# 获取近5、10、30、60日个股上榜统计数据,包括上榜次数、累积购买额、累积卖出额、净额、买入席位数和卖出席位数
def getCapTop(request):
    return HttpResponse()


# 获取营业部近5、10、30、60日上榜次数、累积买卖等情况
def getBrokerTops(request):
    return HttpResponse()


# 获取机构近5、10、30、60日累积买卖次数和金额等情况
def getInstTop(request):
    return HttpResponse()


# 获取最近一个交易日机构席位成交明细统计数据
def getInstDetail(request):
    return HttpResponse()
