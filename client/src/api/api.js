import axios from 'axios';

var instance = axios.create({
  baseURL: '/api',
  timeout: 60000,
  headers: {
    'Accept': 'application/json',
  },

});

instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    var token = localStorage.getItem('token');
    if(token) {
      
      config.headers.authorization = " Bearer " + token;
    }

    console.log(config);
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  });

export const requestRegister = params => { return instance.post(`/register`, params).then(res => res.data); };

export const requestLogin = params => { return instance.post(`/login`, params).then(res => res.data); };

export const createChannel = params => { return instance.post(`/channels`, params).then(res => res.data); };

export const joinChannel = params => { return instance.post(`/channels/${params.channelName}/peers`, params).then(res => res.data); };

export const queryChannel = params => { return instance.get(`/channels`, {params: params}).then(res => res.data); };

export const installChaincode = params => { return instance.post(`/chaincodes`, params).then(res => res.data); };

export const queryChaincode = params => { return instance.get(`/chaincodes`, {params: params}).then(res => res.data); };

export const initChaincode = params => { return instance.post(`/channels/${params.channelName}/chaincodes`, params).then(res => res.data); };

export const query = params => { return instance.get(`/channels/${params.channelName}/chaincodes/${params.chaincodeName}`, {params: params}).then(res => res.data); };

export const invokeChaincode = params => { return instance.post(`/channels/${params.channelName}/chaincodes/${params.chaincodeName}`, params).then(res => res.data); };

export const pay = params => { return instance.post(`/channels/${params.channelName}/chaincodes/${params.chaincodeName}`, params).then(res => res.data); };
