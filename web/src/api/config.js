import service from '@/utils/request'

export const getConfig = () => {
  return service({
    url: '/lukConfig/getConfig',
    method: 'get'
  })
}

export const setConfig = (data) => {
  return service({
    url: '/lukConfig/setConfig',
    method: 'post',
    data
  })
}
