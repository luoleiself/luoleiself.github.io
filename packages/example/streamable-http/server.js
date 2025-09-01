// server.js
import path from 'path';
import { fileURLToPath } from 'url';
import express from 'express';
const app = express();
const PORT = 3000;

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// 模拟数据生成函数
function* generateData() {
  for (let i = 1; i <= 10; i++) {
    yield { id: i, message: `Data chunk ${i}`, timestamp: new Date().toISOString() };
  }
}

app.use(express.static(path.join(__dirname, 'public')));

// Streamable HTTP endpoint
app.get('/stream', (req, res) => {
  // 设置响应头以支持流式传输
  res.setHeader('Content-Type', 'application/json');
  res.setHeader('Transfer-Encoding', 'chunked');

  // 发送初始响应
  res.write('[\n');

  let isFirst = true;
  const dataGenerator = generateData();

  // 模拟流式数据发送
  const interval = setInterval(() => {
    const result = dataGenerator.next();

    if (result.done) {
      // 结束响应
      res.write('\n]');
      res.end();
      clearInterval(interval);
    } else {
      // 发送数据块
      const chunk = result.value;
      const separator = isFirst ? '' : ',\n';
      res.write(`${separator}  ${JSON.stringify(chunk)}`);
      isFirst = false;
    }
  }, 1000); // 每秒发送一个数据块

  // 处理客户端断开连接
  req.on('close', () => {
    clearInterval(interval);
  });
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
