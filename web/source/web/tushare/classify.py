'''
tushare股票分类数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 按照sina财经对沪深股票进行的行业分类，返回所有股票所属行业的信息
def getIndustryClassified(request):
    return HttpResponse()


# 返回股票概念的分类数据
def getConceptClassified(request):
    return HttpResponse()


# 按地域对股票进行分类
def getAreaClassified(request):
    return HttpResponse()


# 获取中小板股票数据，即查找所有002开头的股票
def getSMEClassified(request):
    return HttpResponse()


# 获取创业板股票数据，即查找所有300开头的股票
def getGEMClassified(request):
    return HttpResponse()


# 获取风险警示板股票数据，即查找所有st股票
def getSTClassified(request):
    return HttpResponse()


# 获取沪深300当前成份股及所占权重
def getHS300s(request):
    return HttpResponse()


# 获取上证50成份股
def getSZ50s(request):
    return HttpResponse()


# 获取中证500成份股
def getZZ500s(request):
    return HttpResponse()


# 获取已经被终止上市的股票列表，数据从上交所获取，目前只有在上海证券交易所交易被终止的股票
def getTerminated(request):
    return HttpResponse()


# 获取被暂停上市的股票列表，数据从上交所获取，目前只有在上海证券交易所交易被终止的股票
def getSuspended(request):
    return HttpResponse()
