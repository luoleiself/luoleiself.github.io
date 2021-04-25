# Hexo blog

```bash
  npm start # 启动服务器,开发环境
  npm run build # 生成静态文件
  npm run clean # 清除缓存文件(db.json)和静态文件(public)
  npm run deploy # 生成静态文件并部署
```

## 部署

hexo-deployer-git

\_config.yml

```yml
# Deployment
## Docs: https://hexo.io/docs/one-command-deployment
deploy:
  type: git
  repo: https://github.com/luoleiself/luoleiself.github.io
  # repo: https://bitbucket.org/JohnSmith/johnsmith.bitbucket.io
  branch: [gh-pages]
  # message: [message]
```

## RSS 订阅

hexo-generator-feed

\_config.yml

```yml
# RSS
# https://github.com/hexojs/hexo-generator-feed
feed:
  type: atom
  path: atom.xml
  limit: 20
  hub:
  content:
  content_limit: 140
  content_limit_delim: ' '
  order_by: -date
  icon: icon.png
  autodiscovery: true
  template:
```

\_config.next.yml

```yml
# Social Links
# Usage: `Key: permalink || icon`
# Key is the link label showing to end users.
# Value before `||` delimiter is the target permalink, value after `||` delimiter is the name of Font Awesome icon.
social:
  GitHub: https://github.com/luoleiself || fab fa-github
  E-Mail: mailto:luoleiself@163.com || fa fa-envelope
  #Weibo: https://weibo.com/yourname || fab fa-weibo
  #Google: https://plus.google.com/yourname || fab fa-google
  #Twitter: https://twitter.com/yourname || fab fa-twitter
  #FB Page: https://www.facebook.com/yourname || fab fa-facebook
  #StackOverflow: https://stackoverflow.com/yourname || fab fa-stack-overflow
  #YouTube: https://youtube.com/yourname || fab fa-youtube
  #Instagram: https://instagram.com/yourname || fab fa-instagram
  #Skype: skype:yourname?call|chat || fab fa-skype
  RSS: /atom.xml || fa fa-rss
```

## 本地文章搜索

hexo-generator-searchdb

\_config.yml

```yml
search:
  path: search.xml
  field: post
  format: html
  limit: 10000
```

\_config.next.yml

```yml
# Local Search
# Dependencies: https://github.com/theme-next/hexo-generator-searchdb
local_search:
  enable: true
  # If auto, trigger search by changing input.
  # If manual, trigger search by pressing enter key or search button.
  trigger: auto
  # Show top n results per article, show all results by setting to -1
  top_n_per_article: 1
  # Unescape html strings to the readable one.
  unescape: false
  # Preload the search data when the page loads.
  preload: false
```

## 相关文章推荐

hexo-related-popular-posts

\_config.next.yml

```yml
# Related popular posts
# Dependencies: https://github.com/tea3/hexo-related-popular-posts
related_posts:
  enable: true
  title: # Custom header, leave empty to use the default one
  display_in_home: false
  params:
    maxCount: 5
    #PPMixingRate: 0.0
    #isDate: false
    #isImage: false
    #isExcerpt: false
```

## 文章字数阅读时间统计

hexo-symbols-count-time

\_config.next.yml

```yml
# Post meta display settings
post_meta:
  item_text: true
  created_at: true
  updated_at:
    enable: true
    another_day: true
  categories: true

# Post wordcount display settings
# Dependencies: https://github.com/theme-next/hexo-symbols-count-time
symbols_count_time:
  symbols: true
  time: true
  total_symbols: true
  total_time: true
  exclude_codeblock: false
  awl: 4
  wpm: 275
  suffix: 'mins.'
```
