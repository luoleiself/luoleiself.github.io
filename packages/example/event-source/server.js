// server.js
import path from 'path';
import { fileURLToPath } from 'url';
import express from 'express';
const app = express();
const PORT = 3000;

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

let count = 1

app.use(express.static(path.join(__dirname, 'public')))

app.get('/evt-source', (req, res) => {
  res.setHeader('Content-Type', 'text/event-stream');
  res.setHeader('Cache-Control', 'no-cache');
  res.setHeader('Connection', 'keep-alive');

  // 定期发送事件给客户端
  const sendEvent = (evtName, data) => {
    /**
     * event: notice
     * id: 1738749223324
     * retry: 500
     * data: {"message":"Message at 2025-02-05T09:53:43.324Z","count":1}
     */
    // 数据以文本流的形式从服务器发送到客户端, 每一行数据都必须以换行符(\n)结束, 
    // 重连时间间隔 500 毫秒
    // 并且每条消息以两个换行符(\n\n)结束
    res.write(`event: ${evtName}\nid: ${Date.now()}\nretry: 500\ndata: ${JSON.stringify(data)}\n\n`);
  };

  // 发送初始事件
  sendEvent('connected', { message: 'Connection established' });

  // 模拟每 5 秒发送一次消息
  const timer = setInterval(() => {
    sendEvent(Math.random() > 0.5 ? 'notice' : 'message', { message: `Message at ${new Date().toISOString()}`, count: count });
    count++;
  }, 1000);

  // 当客户端断开连接时清除定时器
  req.on('close', () => {
    clearInterval(timer);
    res.end()
  });
})

app.get('/eventSource.html', (req, res) => {
  res.sendFile('./eventSource.html')
})

app.listen(PORT, () => {
  console.log(`app listen localhost:${PORT}`)
})
