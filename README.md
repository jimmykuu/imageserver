# imageserver
图片服务器

## 开发缘由

golangtc.com 的静态图片原本都放在七牛，由于 golangtc.com 没有备案，因此无法使用 https 获取资源。因此需要自己搭建服务器。

该图片服务器使用 https://github.com/pierrre/imageserver 实现了图片服务器的功能，并支持缩放。

## 访问方式

```
/large.jpg
/avatar/jimmykuu.jpg
```