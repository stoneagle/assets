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
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['top'] = request.POST['top']
    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.profit_data(year=str(params['year']), top=int(params['top']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取业绩预告数据
def getForecastData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.forecast_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 以月的形式返回限售股解禁情况，通过了解解禁股本的大小，判断股票上行的压力
def getXSGData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['month'] = request.POST['month']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.xsg_data(int(params['year']), int(params['month']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取每个季度基金持有上市公司股票的数据
def getFundHoldings(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.fund_holdings(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取IPO发行和上市的时间列表，包括发行数量、网上发行数量、发行价格已经中签率信息等
def getNewStocks(request):
    result = ts.new_stocks()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 沪市的融资融券数据从上海证券交易所网站直接获取，提供了有记录以来的全部汇总和明细数据
def getSHMargins(request):
    params = {}
    if request.POST:
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.sh_margins(start=params['start'], end=params['end'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 沪市融资融券明细数据
def getSHMarginsDetail(request):
    params = {}
    if request.POST:
        params['date'] = request.POST['date']
        params['symbol'] = request.POST['symbol']
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.sh_margin_details(symbol=params['symbol'], date=params['date'], start=params['start'], end=params['end'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 深市的融资融券数据从深圳证券交易所网站直接获取，提供了有记录以来的全部汇总和明细数据
def getSZMargins(request):
    params = {}
    if request.POST:
        params['start'] = request.POST['start']
        params['end'] = request.POST['end']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.sz_margins(start=params['start'], end=params['end'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 深市融资融券明细数据
def getSZMarginsDetails(request):
    params = {}
    if request.POST:
        params['date'] = request.POST['date']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.sz_margin_details(date=params['date'])
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
