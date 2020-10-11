import request from '@/utils/request'

export default {
    
    // 获取blog列表
    getBlogList(page, size) {
        return request({
            url: `/api/v1/blog?p=${page}&limit=${size}`,
            method: 'get'
        })
    },
    // 修改记录
    updateBlogRow(content) {
        return request({
            url: `/api/v1/blog/${content.id}`,
            method: 'put',
            data: content
        })
    },
    // 删除记录
    deleteBlogRow(id) {
        return request({
            url: `/api/v1/blog/${id}`,
            method: 'delete'
        })
    },
    // 新增博客
    postBlogRow(content) {
        return request({
            url: `/api/v1/blog`,
            method: 'post',
            data: content,
        })
    },
    // 查看博客
    getBlogDetail(id) {
        return request({
            url: `/api/v1/blog/${id}`,
            method: 'get'
        })
    },
}