'''
tushare新闻事件数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取即时财经新闻，类型包括国内财经、证券、外汇、期货、港股和美股等新闻信息
def getLatestNews(request):
    return HttpResponse()


# 获取个股信息地雷数据
def getNotices(request):
    return HttpResponse()


# 获取sina财经股吧首页的重点消息
def getGubaSina(request):
    return HttpResponse()
