import service from '@/utils/request'

export const updateLukUserAddress = (data) => {
  return service({
    url: '/lukUserAddress/updateLukUserAddress',
    method: 'post',
    data
  })
}
export const getLukUserAddressList = (params) => {
  return service({
    url: '/lukUserAddress/getLukUserAddressList',
    method: 'get',
    params
  })
}
