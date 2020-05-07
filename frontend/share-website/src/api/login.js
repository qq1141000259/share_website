import request from '@/utils/request'


export function login(username, password) {
    return request({
        url: 'api/user/login',
        method: 'post',
        data: {
            username, // username: username
            password
        }
    })
}

export function getUserInfo(token) {
    return request({
        url: `/user/info/${token}`,
        method: 'get'
    })
}

export function logout(token) {
    return request({
        url: `api/user/logout`,
        method: 'post'
    })
}

export function register(data) {
    return request({
        url: `api/user/register`,
        method: 'post',
        data: data,
    })
}

