---
title: Grid
date: 2024-05-17 16:00:55
categories:
  - WebAPI
tags:
  - Grid
---

## Grid

CSS Grid 是一个用于 web 的二维布局系统, 网格是由一系列水平及垂直的线构成的一种布局模式, 根据网格, 可以将内容按照行和列的格式进行排版.
CSS Grid 布局 和 Flexible 布局的主要区别在于 CSS Flexible 是为了一维布局服务的(沿横向或纵向), 而 CSS Grid 是为二维布局服务的(同时沿着横向和纵向)

### 辅助属性

- 网格轨道, 网格布局中使用 grid-template-rows 和 grid-template-columns 属性定义的网格上的行和列, 网格轨道是网格上任意两条相邻线之间的空间.
- fr 单位, 新的长度单位, 表示网格容器中占用的轨道
- minmax(min, max), 函数为一个行/列的尺寸设置了取值范围, 例如 minmax(100px, auto) 表示尺寸至少为 100px, 如果内容尺寸大于 100px 则会根据内容自动调整
- repeat(count, value), 表示轨道列表的重复片段, 允许以更紧凑的形式写入大量显示重复模式的列或行

<!-- more -->

### 容器属性

#### display

- grid 指定容器的行为类似块级元素并且采用网格布局
- inline-grid 指定容器的行为类似行内元素并且采用网格布局

#### grid-template-rows

基于**网格行**的维度定义网格线的名称和网格轨道的尺寸大小

```css
/*自动调整网格轨道大小, 等价于 max-content, min-content*/
grid-template-rows: auto;
/*定义 4 行并且指定每行的行高, 第 2 行高度自适应*/
grid-template-rows: 40px auto 40px 50px;
/*定义 3 行, 
第 1 行分配 1fr 可用空间,
第 2 行分配 2fr 可用空间, 
第 3 行分配 3fr 可用空间*/
grid-template-rows: 1fr 2fr 3fr;
/*定义 3 行, 每行分配 33.33% 可用空间*/
grid-template-rows: repeat(3, 33.33%);
```

- 网格线, 使用**方框号**指定每根网格线的名字, 默认从 1 开始

```css
/*定义 3 行并且指定每行的行高, 
第 2 行高度自适应, 
指定每根网格线的名字 r1 开始*/
grid-template-rows: [r1] 100px [r2] auto [r3] 100px [r4];
```

#### grid-template-columns

基于**网格列**的维度定义网格线的名称和网格轨道的尺寸大小

```css
/*自动调整网格轨道大小, 等价于 max-content, min-content*/
grid-template-columns: auto;
/*定义 3 列并指定第 1 列和第 3 列的列宽, 第 2 列宽度自适应*/
grid-template-columns: 40px auto 50px;
/*定义 3 列并且指定每列的列宽, 第 3 列分配 2fr 可用空间*/
grid-template-columns: 40px 40px 2fr;
/*定义 2 列, 第 1 列分配 1fr 可用空间, 第 2 列分配 2fr 可用空间*/
grid-template-columns: 1fr 2fr;
/*定义 3 列, 
第 1 列和第 3 列分配 1fr 可用空间, 
第 2 列分配 2fr 可用空间*/
grid-template-columns: repeat(1, 1fr 2fr 1fr);
```

- 网格线, 使用**方框号**指定每根网格线的名字, 默认从 1 开始

```css
/*定义 3 列并且指定每列的列宽, 第 2 列宽度自适应，指定每根网格线的名字 c1 开始*/
grid-template-columns: [c1] 100px [c2] auto [c3] 100px [c4];
```

#### grid-template-areas

定义网格区域的别名, 网格区域和网格项没有关联, 但是可以和[网格定位](#grid-area)属性关联, 例如 `grid-row-start`, `grid-column-start`, `grid-row-end`, `grid-column-end`

- 如果某些区域不需要, 则使用 `.` 表示
- 区域别名会影响到网格线, 每个区域的起始网格线自动命名为 `区域名-start`, 终止网格线自动命名为 `区域名-end`

```css
/* 定义九宫格每个区域别名 */
grid-template-areas:
  'a b c'
  'd e f'
  'g h i';

/*左侧区域 b, 右上角区域 a, 右下角区域 c*/
grid-template-areas:
  'b b a'
  'b b c'
  'b b c';

/*左上角区域 a, 右下角为 b 和 c, 右上角和左下角未指定区域别名*/
grid-template-areas:
  'a a .'
  'a a .'
  '. b c';

/*顶部页眉区域 header, 底部页脚区域 footer, 中间区域 sidebar 和 main*/
grid-template-areas:
  'header header header'
  'sidebar main main'
  'footer footer footer';
```

#### grid-auto-flow

精确指定在网格中被自动布局的元素如何排列, 默认值为 row

- row 指定自动布局算法按照通过逐行填充来排列元素, 在必要时增加新行
- column 指定自动布局算法按照通过逐列填充来排列元素, 在必要时增加新列
- dense 指定自动布局算法使用一种`稠密`堆积算法, 省略时使用 `稀疏` 算法, 配合 row 或 column 使用
  - `稠密`算法, 如果后面出现稍小的元素, 则会试图去填充网格中前面留下的空白, 这样会填上稍大元素留下的空白, 但同时也可能导致原来出现的次序被打乱
  - `稀疏`算法, 布局算法只会向前移动, 永远不会填充空白, 这样保证了所有自动布局元素按照次序出现, 即使可能会留下被后面元素填充的空白

```css
grid-auto-flow: row;
grid-auto-flow: column;

grid-auto-flow: row dense;
grid-auto-flow: column dense;
```

#### grid-auto-rows

指定**隐式**创建的网格横向轨道的高度, 例如某项超出了当前的网格布局区域, 浏览器会自动生成多余的网格以便放置项目

```css
grid-auto-rows: 20px;
grid-auto-rows: 3fr;
grid-auto-rows: 25%;
grid-auto-rows: minmax(100px, auto);
grid-auto-rows: minmax(max-content, 2fr);
```

#### grid-auto-columns

指定**隐式**创建的网格纵向轨道的宽度, 作用参考 `grid-auto-rows`

```css
grid-auto-columns: 20px;
grid-auto-columns: 3fr;
grid-auto-columns: 25%;
grid-auto-columns: minmax(100px, auto);
grid-auto-columns: minmax(max-content, 2fr);
```

#### grid-gap

`grid-gap: <grid-row-gap> <grid-column-gap>;` 网格行和列间隙的缩写

指定网格行和列之间的间隙, 此属性作为 `gap` 兼容低版本的别名

```css
grid-gap: 10px 20px;
grid-row-gap: 20px;
grid-column-gap: 10px;

gap: 10px 20px;
row-gap: 20px;
column-gap: 10px;
```

##### grid-row-gap

指定**网格行**之间的间隙大小, 此属性作为 `row-gap` 兼容低版本的别名

##### grid-column-gap

指定**网格列**之间的间隙大小

### 项目属性

#### grid-area <em id="grid-area"></em> <!-- markdownlint-disable-line -->

`grid-area: <row-start> / <column-start> / <row-end> / <column-end>;` 合并简写形式

指定项目放在哪个区域, 是一种对 `grid-row-start`, `grid-column-start`, `grid-row-end`, `grid-column-end`的简写, 通过基线、 跨度(span)或没有(自动)的网格线放置在指定一个网格项的大小和位置, 继而确定区域的边界

- 如果指定 4 个 `<grid-line>` 的值时, 第 1 个值为 `grid-row-start`, 第 2 个值为 `grid-column-start`, 第 3 个值为`grid-row-end`, 第 4 个值为 `grid-column-end`
  - 当 `grid-column-end` 被忽略时, 如果 `grid-column-start` 为自定义关键字数据类型, `grid-column-end` 则为该值, 否则为 auto;
  - 当 `grid-row-end` 被忽略时, 如果 `grid-row-start` 为自定义关键字数据类型, `grid-row-end` 则为该值, 否则为 auto;
  - 当 `grid-column-start` 被忽略时, 如果 `grid-row-start` 为自定义关键字数据类型, 则所有四项普通写法的属性值均为该值, 否则为 auto;
- span, 为网格单元定义一个跨度, 使得网格单元的网格区域中的一条边界远离另一条边界 n 条基线. 默认值为 1;

```css
grid-area: auto;
grid-area: 2 / 2 / auto / span 3;

.item-1 {
  /* 放在 3 行 3 列的中心网格位置 */
  /* 网格线编号, 规则见上 */
  grid-area: 2 / 2;
  /* 网格线名 */
  grid-area: r2 / c2 / r3 / c3;
  /* 区域别名 */
  grid-area: e;
  /* 区域别名网格线起始名和终止名 */
  grid-area: e-start / e-start / e-end / e-end;
}
```

#### grid-row

`grid-row: <start-line> / <end-line>;` 合并简写形式

```css
.item-1 {
  /* 上边框是第 2 根水平网格线, 下边框是第 4 根水平网格线 */
  grid-row: 2 / 4;
  /* 左边框是第 1 根垂直网格线, 有边框是第 3 根垂直网格线 */
  grid-column: 1 / 3;

  /* 使用网格线名 */
  grid-row: r2 / r4;
  grid-column: c1 / c3;
}
```

##### grid-row-start

指定上边框所在的水平网格线

##### grid-row-end

指定下边框所在的水平网格线

#### grid-column

`grid-column: <start-line> / <end-line>;` 合并简写形式

```css
.item-1 {
  /* 上边框是第 2 根水平网格线, 下边框是第 3 根水平网格线 */
  grid-row: 2 / 3;
  /* 左边框是第 2 根垂直网格线, 有边框是第 4 根垂直网格线 */
  grid-column: 2 / 4;

  /* 使用网格线名 */
  grid-row: r2 / r3;
  grid-column: c2 / c4;
}
```

##### grid-column-start

指定左边框所在的垂直网格线

##### grid-column-end

指定有边框所在的垂直网格线
