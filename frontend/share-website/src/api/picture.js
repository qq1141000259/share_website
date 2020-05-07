import request from '@/utils/request'

export default {
    // 获取会员列表数据
    getPicList(page, size) {
        return request({
            url: `/api/v1/picture?p=${page}&limit=${size}`,
            method: 'get'
        })
    },
    // 获取视频列表
    getVideoList(page, size) {
        return request({
            url: `/api/v1/video?p=${page}&limit=${size}`,
            method: 'get'
        })
    },
    // 获取留言版列表
    getMsgList(page, size) {
        return request({
            url: `/api/v1/message?p=${page}&limit=${size}`,
            method: 'get'
        })
    },
    // 删除记录
    deleteMsgRow(id) {
        return request({
            url: `/api/v1/message/${id}`,
            method: 'delete'
        })
    },
    // 新增留言
    postMsgRow(content) {
        return request({
            url: `/api/v1/message`,
            method: 'post',
            data: {content: `${content}`}
        })
    },
    // 获取背景图片list
    getBackPicList() {
        return request({
            url: `/api/v1/back_img_url`,
            method: 'get'
        })
    }
    
}