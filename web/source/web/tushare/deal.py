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
    params = {}
    if request.POST:
        params['code'] = request.POST['code']
        params['ktype'] = request.POST['ktype']
        params['autype'] = request.POST['autype']
        params['index'] = request.POST.get('index', False)
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_k_data(
            code=params['code'],
            start=params['start'],
            end=params['end'],
            ktype=params['ktype'],
            autype=params['autype'],
            index=params['index']
        )
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            # result['date'] = result.index
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取历史数据
def getHistData(request):
    params = {}
    if request.POST:
        params['code'] = request.POST['code']
        params['ktype'] = request.POST['ktype']
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']

    if (not params or not fw.checkParamsEmpty(params)):
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
    result = ts.get_today_all()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取个股以往交易历史的分笔数据明细
def getTickData(request):
    params = {}
    if request.POST:
        params['code'] = request.POST['code']
        params['date'] = request.POST['date']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_tick_data(code=params['code'], date=params['date'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取实时分笔数据，可以实时取得股票当前报价和成交信息
def getRealtimeQuotes(request):
    params = {}
    if request.POST:
        params['symbols'] = request.POST.getlist('symbols[]')

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_realtime_quotes(symbols=params['symbols'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取当前交易日（交易进行中使用）已经产生的分笔明细
def getTodayTicks(request):
    params = {}
    if request.POST:
        params['code'] = request.POST.get('code')

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_today_ticks(code=params['code'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取大盘指数行情列表
def getIndex(request):
    result = ts.get_index()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取大单交易数据，默认为大于等于400手
def getSinaDD(request):
    params = {}
    if request.POST:
        params['code'] = request.POST.get('code')
        params['date'] = request.POST.get('date')
        params['vol'] = request.POST.get('vol', 400)

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_sina_dd(code=params['code'], date=params['date'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
