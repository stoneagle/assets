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
    result = ts.get_stock_basics()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取业绩报表数据
def getReportData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_report_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取盈利能力数据
def getProfitData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_profit_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取营运能力数据
def getOperationData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_operation_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取成长能力数据
def getGrowthData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_growth_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取偿债能力数据
def getDebtpayingData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_debtpaying_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 按年度、季度获取现金流量数据
def getCashflowData(request):
    params = {}
    if request.POST:
        params['year'] = request.POST['year']
        params['season'] = request.POST['season']

    if (not params or not fw.checkParamsEmpty(params)):
        ret = fw.formatResponse(vs.ERRNO_PARAMS, {})
    else:
        result = ts.get_cashflow_data(int(params['year']), int(params['season']))
        if result is None:
            ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
        else:
            result = result.to_json(orient='records')
            ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
