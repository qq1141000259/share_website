import request from '@/utils/request'

export default {
    
    // 获取留言版列表
    getWordList(page, size, word, annotation, remark) {
        console.log(word)
        return request({
            url: `/api/v1/word_book?p=${page}&limit=${size}&word=${word}&annotation=${annotation}&remark=${remark}`,
            method: 'get'
        })
    },
    // 修改记录
    updateWordRow(content) {
        return request({
            url: `/api/v1/word_book/${content.id}`,
            method: 'put',
            data: content
        })
    },
    // 删除记录
    deleteWordRow(id) {
        return request({
            url: `/api/v1/word_book/${id}`,
            method: 'delete'
        })
    },
    // 新增单词
    postWordRow(content) {
        return request({
            url: `/api/v1/word_book`,
            method: 'post',
            data: content,
        })
    },
    
}