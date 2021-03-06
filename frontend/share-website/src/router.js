import Vue from "vue";
import Router from "vue-router";
// import Login from './views/login/index.vue'
//  下面情况，默认的会导入 ./views/login 目录下的index.vue组件
import Login from './views/login'
import Layout from '@/components/Layout.vue'
import Home from './views/home'
import Blog from './views/blog'
import BlogEdit from './views/blog/blog_edit'
import BlogQuery from './views/blog/blog_query'
import Word from './views/cs_word'
import Picture from './views/picture_share'
import Video from './views/video_share'
import Message from './views/message'
import Upload from './views/upload'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login', //路由名称
      component: Login //组件对象
    },
    {
      path: '/',
      name: 'layout', //路由名称
      component: Layout, //组件对象
      redirect: '/home',
      children: [
        {
          path: '/home',
          component: Home,
          meta: {title: '首页'}
        }
      ]
    },
    {
      path: '/blog',
      component: Layout, //组件对象
      children: [
        {
          path: '/blog',
          component: Blog,
          meta: {title: '博客'}
        },
        {
          path: '/blog/edit',
          component: BlogEdit,
          meta: {title: '编辑'}
        },
        {
          path: '/blog/:id',
          component: BlogQuery,
          meta: {title: '查看'}
        },
      ]
    },
    {
      path: '/word',
      component: Layout, //组件对象
      children: [
        {
          path: '/',
          component: Word,
          meta: {title: '单词薄'}
        },
      ]
    },
    {
      path: '/picture',
      component: Layout, //组件对象
      children: [
        {
          path: '/',
          component: Picture,
          meta: {title: '图片'}
        },
      ]
    },
    {
      path: '/video',
      component: Layout, //组件对象
      children: [
        {
          path: '/',
          component: Video,
          meta: {title: '视频'}
        },
      ]
    },
    {
      path: '/message',
      component: Layout, //组件对象
      children: [
        {
          path: '/',
          component: Message,
          meta: {title: '留言薄'}
        },
      ]
    },
    {
      path: '/upload',
      component: Layout,
      children: [
        {
          path: '/', // /member/
          component: Upload,
          meta: {title: '上传'}
        }
      ]
    }
  ]
})
