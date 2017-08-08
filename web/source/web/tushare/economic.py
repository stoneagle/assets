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
    result = ts.get_deposit_rate()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取贷款利率
def getLoanRate(request):
    result = ts.get_loan_rate()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取存款准备金率
def getRRRRate(request):
    result = ts.get_rrr()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取货币供应量
def getMoneySupply(request):
    result = ts.get_money_supply()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 获取货币供应量(年底余额)
def getMoneySupplyBal(request):
    result = ts.get_money_supply_bal()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 国内生产总值(季度)
def getGDPYear(request):
    result = ts.get_gdp_year()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 国内生产总值(季度)
def getGDPQuarter(request):
    result = ts.get_gdp_quarter()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 三大需求对GDP贡献
def getGDPFor(request):
    result = ts.get_gdp_for()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 三大产业对GDP拉动
def getGDPPull(request):
    result = ts.get_gdp_pull()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 三大产业贡献率
def getGDPContrib(request):
    result = ts.get_gdp_contrib()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 居民消费价格指数
def getCPI(request):
    result = ts.get_cpi()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)


# 工业品出厂价格指数
def getPPI(request):
    result = ts.get_ppi()
    if result is None:
        ret = fw.formatResponse(vs.ERRNO_TUSHARE, {})
    else:
        result = result.to_json(orient='records')
        ret = fw.formatResponse(vs.ERRNO_OK, json.loads(result))
    return HttpResponse(ret)
