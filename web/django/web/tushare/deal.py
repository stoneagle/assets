'''
tushare交易数据
'''
from django.http import HttpResponse
import tushare as ts
from web.library import vars as vs
from web.library import framework as fw
import json


# 获取K线数据，融合get_hist_data与get_h_data两个接口功能
def getKData(request):
    return HttpResponse()


# 获取历史数据
def getHistData(request):
    params = {}
    if request.POST:
        params['code'] = request.POST['code']
        params['ktype'] = request.POST['ktype']
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']

    if not fw.checkParamsEmpty(params):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_hist_data(code=params['code'], start=params['start'], end=params['end'], ktype=params['ktype'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result['date'] = result.index
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))

    return HttpResponse(ret)


# 获取复权数据
def getStockBasics(request):
    return HttpResponse()


# 获取当前交易所有股票的行情数据（如果是节假日，即为上一交易日
def getTodayAll(request):
    return HttpResponse()


# 获取个股以往交易历史的分笔数据明细
def getTickData(request):
    return HttpResponse()


# 获取实时分笔数据，可以实时取得股票当前报价和成交信息
def getRealtimeQuotes(request):
    return HttpResponse()


# 获取当前交易日（交易进行中使用）已经产生的分笔明细
def getTodayTicks(request):
    return HttpResponse()


# 获取大盘指数行情列表
def getIndex(request):
    return HttpResponse()


# 获取大单交易数据，默认为大于等于400手
def getSinaDD(request):
    return HttpResponse()
