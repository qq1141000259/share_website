import request from '@/utils/request'
import defaultReq from '@/utils/request'
import uuidv1 from 'uuid/v1' 


export default {
    // 获取会员列表数据
    getPolicy() {
        return request({
            url: `/api/v1/picture?p=${page}&limit=${size}`,
            method: 'get'
        })
    },
    upload(data, filetype, title) {
        var COS = require('cos-js-sdk-v5');
        var cos = new COS({
            // 必选参数
            getAuthorization: function (options, callback) {
                // 服务端 JS 和 PHP 例子：https://github.com/tencentyun/cos-js-sdk-v5/blob/master/server/
                // 服务端其他语言参考 COS STS SDK ：https://github.com/tencentyun/qcloud-cos-sts-sdk
                // STS 详细文档指引看：https://cloud.tencent.com/document/product/436/14048
                request({
                    url: `/api/v1/upload/credentials`,
                    method: 'get'
                }).then(response => {
                    const resp = response.data
                    if(!resp.code){
                        console.log(resp.data)
                        console.log(resp.code)
                        var credentials = resp.data;
                        if (!credentials) return console.error('credentials invalid');
                        callback({
                            TmpSecretId: credentials.TmpSecretId,
                            TmpSecretKey: credentials.TmpSecretKey,
                            XCosSecurityToken: credentials.Token,
                            // 建议返回服务器时间作为签名的开始时间，避免用户浏览器本地时间偏差过大导致签名错误
                            StartTime: credentials.StartTime, // 时间戳，单位秒，如：1580000000
                            ExpiredTime: credentials.ExpiredTime, // 时间戳，单位秒，如：1580000900
                        });
                    }
                });
            }
        });
        const suffix = data.name.split(".")[1]
        console.log(suffix)
        const key = uuidv1().replace(new RegExp("-","gm"),"")
        cos.putObject({
            Bucket: 'zghtly-bucket-1301493247', /* 必须 */
            Region: 'ap-shanghai',     /* 存储桶所在地域，必须字段 */
            Key: key + "." + suffix,              /* 必须 */
            StorageClass: 'STANDARD',
            Body: data, // 上传文件对象
            onProgress: function(progressData) {
                console.log(JSON.stringify(progressData));
            }
        }, function(err, data) {
            console.log(err || data);
            if(data){
                const addDat = {
                    key: key,
                    suffix: suffix,
                    type: filetype,
                    title: title,
                }
                request({
                    url: `/api/v1/upload`,
                    method: 'post',
                    data: addDat,
                }).then(response => {
                    if(response.data.code){
                        alert("上传失败")
                    }
                })
            }
        });
    }
}
