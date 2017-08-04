'''
tushare投资参考数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 返回股票的送转和分红预案情况
def getProfitData(request):
    return HttpResponse()


# 按年度、季度获取业绩预告数据
def getForecastData(request):
    return HttpResponse()


# 以月的形式返回限售股解禁情况，通过了解解禁股本的大小，判断股票上行的压力
def getXSGData(request):
    return HttpResponse()


# 获取每个季度基金持有上市公司股票的数据
def getFundHoldings(request):
    return HttpResponse()


# 获取IPO发行和上市的时间列表，包括发行数量、网上发行数量、发行价格已经中签率信息等
def getNewStocks(request):
    return HttpResponse()


# 沪市的融资融券数据从上海证券交易所网站直接获取，提供了有记录以来的全部汇总和明细数据
def getSHMargins(request):
    return HttpResponse()


# 沪市融资融券明细数据
def getSHMarginsDetail(request):
    return HttpResponse()


# 深市的融资融券数据从深圳证券交易所网站直接获取，提供了有记录以来的全部汇总和明细数据
def getSZMargins(request):
    return HttpResponse()


# 深市融资融券明细数据
def getSZMarginsDetails(request):
    return HttpResponse()
