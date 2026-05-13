import express from 'express';
import session from 'express-session';
import cookieParser from 'cookie-parser';

const app = express()
let count = 1

app.use(express.static('./')).use(cookieParser())
app.use(session({
  secret: 'helloexpresssession', // 用于签名 session ID 的密钥
  resave: false,             // 强制保存 session 即使未修改
  saveUninitialized: true,   // 是否保存未初始化的 session
  cookie: {
    secure: false,         // 如果使用 HTTPS，则应设为 true
    httpOnly: true,        // 防止通过 JavaScript 访问 cookie
    maxAge: null,          // 不设置 maxAge，让 cookie 成为会话级 cookie
  }
}))

app.get('/evt-source', (req, res) => {
  res.setHeader('Content-Type', 'text/event-stream');
  res.setHeader('Cache-Control', 'no-cache');
  res.setHeader('Connection', 'keep-alive');

  console.log(req.get('host'), req.method, req.params, req.query, req.get('user-agent'), req.cookies)

  // res.cookie('sID', 110, { maxAge: null, httpOnly: true })
  // 定期发送事件给客户端
  const sendEvent = (evtName, data) => {
    /**
     * event: notice
     * id: 1738749223324
     * data: {"message":"Message at 2025-02-05T09:53:43.324Z","count":1}
     */
    // 数据以文本流的形式从服务器发送到客户端, 每一行数据都必须以换行符(\n)结束, 
    // 并且每条消息以两个换行符(\n\n)结束
    res.write(`event: ${evtName}\nid: ${Date.now()}\nretiry: 10\ndata: ${JSON.stringify(data)}\n\n`);
  };

  // 发送初始事件
  sendEvent('connected', { message: 'Connection established' });

  // 模拟每 5 秒发送一次消息
  const timer = setInterval(() => {
    sendEvent(Math.random() > 0.5 ? 'notice' : 'message', { message: `Message at ${new Date().toISOString()}`, count: count });
    count++;
  }, 5000);

  // 当客户端断开连接时清除定时器
  req.on('close', () => {
    clearInterval(timer);
    console.log('req close')
    res.end()
  });
})

app.get('/eventSource.html', (req, res) => {
  res.sendFile('./eventSource.html')
})

app.listen(8080, () => {
  console.log('app listen localhost:8080')
})
