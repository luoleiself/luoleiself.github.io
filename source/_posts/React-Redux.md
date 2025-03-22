---
title: React-Redux.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

## Redux

官方推荐使用封装了 Redux 核心的 @reduxjs/toolkit(RTK) 包, 包含了构建 Redux 应用所必须的 API 方法和常用依赖, 简化了大部分 Redux 任务, 阻止了常见错误, 并让编写 Redux 应用程序变得更容易

<!-- 
- dispatch 只能处理同步的 action

- createStore  创建一个 Redux 存储实例
- combineReducers 将多个 reducer 函数合并成为一个更大的 reducer
- applyMiddleware 将多个中间件组合成一个 store 增强器
- compose 将多个 store 增强器合并成一个单一的 store 增强器 
-->

### [@reduxjs/toolkit/query](#RTK-Query)

独立可选的入口, 允许定义端点(REST, GraphQL或任何异步函数)并生成 reducer 和中间件来完整管理数据获取, 加载状态更新和结果缓存, 还可以自动生成 React Hooks, 可用于组件获取数据

### @reduxjs/toolkit(RTK)

- 通过单一清晰的函数调用简化 store 设置, 同时保留完全配置 store 选项的能力
- 消除意外的 mutations
- 消除手写任何 actionCreator 或 actionType 的需求
- 消除编写容易出错的手动不可变更新逻辑的需求, createSlice 使用 `Immer` 库来编写 reducer, 可以直接修改 state 状态而不需要使用解构语法
- 允许将相关的代码放在一个文件中, 而不是分布在多个独立文件中
- 提供优秀的 TypeScript 支持, 其 API 被设计成很好的安全性, 同时减少代码中需要定义的类型数量
- RTK Query 可以消除编写任何 thunk, reducer, actionCreator 或者副作用狗子来管理数据获取和跟踪加载状态的需求

#### configureStore

特点

- slice reducers 自动传递给 combineReducers
- 自动添加了 `redux-thunk` 中间件
- 添加了 Devtools 中间件来捕获更多意外的变更
- 自动设置了 Redux Devtools Extension
- 中间件和 Devtools 增强器被组合在一起添加到了 store 中

<!--more-->

参数

- reducer
  - 如果是一个函数, configureStore 直接使用其作为根 reducer
  - 如果是一个 slice reducers 的对象, configureStore 将使用 combineReducers 合并此对象并自动创建根 reducer
- middleware 函数, 接收 `getDefaultMiddleware` 函数作为参数, 并返回一个中间件数组, 如果未提供, configureStore 将调用 `getDefaultMiddleware` 设置中间件数组
- devTools 是否设置 Redux Devtools, 默认 true
- preloadedState 初始化状态
- enhancers 增强器函数, 和 middleware 参数作用类似, 使用 `getDefaultEnhancers` 函数获取默认的增强器列表

```jsx
import {configureStore} from '@reduxjs/toolkit';
import {offline} from '@redux-offline/redux-offline';
import offlineConfig from '@redux-offline/redux-offline/lib/defaults'

const store = configureStore({
  reducer: {
    counter: counterSlice.reducer
    // ...
  },
  middleware: getDefaultMiddleware => getDefaultMiddleware().concat(thunk),
  enhancers: getDefaultEnhancers => getDefaultEnhancers().concat(offline(offlineConfig)))
});
```

##### middleware

thunk 中间件实现原理

```jsx
// thunk 中间件实现原理
function thunk(store){
  const next = store.dispatch; // 缓存原 dispatch 方法
  function dispatchFn(action){
    if(typeof action === 'function'){
      // 如果 action 是一个函数, 则传入重写的 dispatchFn 方法
      action(store.dispatch, store.getState);
    } else {
      // 否则直接调用原 dispatch 方法派发 action
      next(action); 
    }
  }
  store.dispatch = dispatchFn;
}
```

applyMiddleware 实现原理

```jsx
export default function applyMiddleware(store, ...fns){
  fns.forEach(fn => {
    fn(store)
  });
}
```

#### createAction <em id="createAction"></em> <!--markdownlint-disable-line-->

用于创建 action 的辅助函数

- type 字符串, 标识 action
- prepareAction() 可选, 函数, 接收任意个参数作为 action 的 payload 的值

返回值: actionCreator

- actionCreator.match 函数可以区分 action 是否是同一类型, TypeScript 中可以识别 action 中 payload 的类型

```tsx
import {createAction} from '@reduxjs/toolkit';

const actionCreator = createAction(type, prepareAction?);

// action 类型常量
function increment(amount: number){
  return {
    type: 'INCREMENT',
    payload: amount
  }
}
const action = increment(3);
// {type: 'INCREMENT', payload: 3}

// action 创建函数
const increment = createAction('INCREMENT', (text: string, age: number) => {
  return {
    type: "INCREMENT",
    payload: {
      text: text,
      age: age,
      id: nanoid(),
      createAt: new Date()
    }
  }
});
const action = increment('hello createAction', 18);
// {type: "INCREMENT", payload: {text: "hello createAction", age: 18, id, createAt}}

const increment = createAction<number>('INCREMENT');
function someFn(action: Action){
  if(increment.match(action)){
    // action.payload can be used as `number` here
  }
}
```

#### createReducer <em id="createReducer"></em> <!--markdownlint-disable-line-->

一个简化创建 reducer 函数的工具, 内部使用 `Immer` 库通过在 reducer 中编写可变代码, 大大简化了不可变的更新逻辑, 并支持将特定的操作类型直接映射到 case reducer 函数

- initialState 初始化状态, 可以是一个返回 state 的函数
- builderCallback 回调函数接收一个 `builder` 对象通过 addCase 方法添加 reducer <em id="builderCallback"></em> <!--markdownlint-disable-line-->
  - addCase() 接收两个参数, 调用必须在 `addMatcher` 和 `addDefaultCase` 之前
    - actionCreatorOrType 指定 action.type
    - reducer
  - addMatcher() 匹配传入的 action, 调用必须在 `addCase` 之后和 `addDefaultCase` 之前
    - matcher() 匹配函数, 匹配传入的所有可能的 action.type, 并按定义的顺序调用
    - reducer
  - addDefaultCase() 添加默认的 reducer
    - reducer

返回值: reducer 函数

- 包含 `getInitialState()`, 返回初始状态, 通常用于测试或者配合 React `useReducer` Hook 实现 SSR

```jsx
import {createReducer, createAction} from '@reduxjs/toolkit';
const reducer = createReducer(initialState, builderCallback);

// 普通 reducer
function coutenReducer(state = initialState, action){
  switch(aciton.type){
    case 'increment':
      return {...state, value: state.value++}
    case 'decrement':
      return {...state, value: state.value--}
    case 'incrementByAmount':
      return {...state, value: state.value + action.payload}
    default:
      return {...state}
  }
}

// createReducer
const increment = createAction('counter/increment');
const decrement = createAction('counter/decrement');
const incrementByAmount = createAction('counter/incrementByAmount');

const counterReducer = createReducer(initialState, builder => {
  builder.addCase(increment, (state, action) => {
    state.value++; // immer 创建的 state 副本, 直接修改
  }).addCase(decrement, (state, action) => {
    state.value--;
  }).addCase(incrementByAmount, (state, action) => {
    state.value += action.payload;
  }).addMatcher((action) => isMatchedAction(action.type), (state, action) => {
    // ...
  }).addDefaultCase((state, action) => {
    // ...
  });
})
```

#### createSlice <em id="createSlice"></em> <!--markdownlint-disable-line-->

使用 `Immer` 库编写 reducer, 不需要使用解构语法修改 state, 接受一个初始状态, 对象或者 reducer 函数, 并自动创建一个与 reducer 和 状态对应的 [actionCreator](#createAction), 在内部调用 `createAction` 和 `createReducer`

- name  标识 state, 将作为生成的 [actionCreator](#createAction) 的前缀
- initialState 初始化状态
- reducers
  - 对象方式, 每个 属性方法名 都是一个 reducer
  - 如果需要自定义 case Reducer, 每个 reducer 将是一个具有 prepare 函数 和 reducer 函数的对象
    - prepare()
    - reducer
  - 如果是一个函数, 将接收一个 create 对象, 具有三个方法
    - create.reducer(reducer) 标准的 reducer
    - create.prepareReducer(prepare, reducer) 自定义 actionCreator 的 payload
    - create.asyncThunk(thunk, opts) 创建异步的函数代替 actionCreator
      - pending
      - fulfilled
      - rejected
- extraReducers 函数, 处理自己创建的 [actionCreator](#createAction) 之外的情况, 如处理异步请求的状态, 同 [builderCallback](#builderCallback)
- reducerPath 标识 slice 的位置, 默认 name
- selectors 接收 state 作为第一个参数和剩余的参数并返回指定结果

返回值, 包含上面的部分属性

- reducer
- actions
- caseReducers
- getInitialState() 返回初始状态
- selectSlice 关联自动创建的一个 selector
- getSelectors()
- injectInfo() 注入 slice

```jsx
// counter.js
import {createSlice, configureStore} from '@reduxjs/toolkit';

const counterSlice = createSlice({
  name: 'counter',
  initialState: {value: 0},
  // reducers 为一个对象
  reducers: {
    increment(state, action) {
      state.value++;
    },
    decrement(state, action){
      state.value--;
    }
  },
  // 自定义 case reducer, prepareAction
  reducers: {
    // case reducer, prepareAction
    incrementByAmout: {
      reducer(state, action){
        state.value += action.payload.value;
      },
      prepare(text: string){
        return {payload: {text: text, value: 100}}
      }
    }
  },
  // reducers 为一个函数, 接收一个 create 对象作为参数, 并返回一个包含 reducer 的对象
  // create 包含 3 个函数: reducer, prepareReducer, asyncThunk
  reducers: (create) => ({
      increment: create.reducer(state, action) => {
        state.value++;
      },
      decrement: create.reducer(state, action) => {
        state.value--;
      },
      incrementByAmount: create.prepareReducer(
        (text: string) => {
          return { payload: {text: text, value: 100}}
        }, (state, action) => {
          // 从 prepare 回调推断 action type
          state.value += action.payload.value;
        }
      ),
      fetchTodo: create.asyncThunk(
        async (id: string, thunkApi) => {
          const res = await fetch(thunkApi);
          return (await res.json()) as Item
        }, {
          pending: state => {
            state.loading = true;
          },
          rejected: state =>{
            state.loading = false;
          },
          fulfilled: (state, action) => {
            state.loading = false;
            state.todos.push(action.payload);
          }
        }
      )
  }),
  // 处理自己创建的 actionCreator 之外的情况
  extraReducers(builder){
    builder.addCase('INCREMENT', (state, action) => {
      state.value++;
    })
  }
});
export const { increment, decrement, incrementByAmount, fetchTodo } = counterSlice.actions;
export default counterSlice.reducer;

// store.js
import counterReducer { increment, decrement, incrementByAmount, fetchTodo } from 'counter.js';
const store = configureStore({
  reducer: {
    counter: counterReducer
  }
})
store.dispatch(increment());
sotre.dispatch(decrement());
store.dispatch(incrementByAmount({value: 10}));

store.dispatch({type: 'counter/increment'})
store.dispatch({type: 'counter/decrement'})
```

两种获取 selector 的方式

- selectors

```jsx
const counterSlice = createSlice({
  name: 'counter',
  initialState: { value: 0 } satisfies CounterState as CounterState,
  reducers: {
    // ...
  },
  selectors: {
    selectValue: (sliceState) => sliceState.value,
  },
});
// createSlice 默认创建一个 selectSlice 方法
console.log(counterSlice.selectSlice({ counter: { value: 2 } })) // { value: 2 }

// 通过 slice 实例的 selectors 属性获取所有的 selector
const { selectValue } = counterSlice.selectors
console.log(selectValue({ counter: { value: 2 } })) // 2
```

- getSelectors()

```jsx
const { selectValue } = counterSlice.getSelectors(
  (rootState: RootState) => rootState.aCounter,
)
console.log(selectValue({ aCounter: { value: 2 } })) // 2

const {selectValue} = counterSlice.getSelectors();
console.log(selectValue({value: 2})) //  2
```

dispatch 提交

- dispatch 提交 action 时, 如果参数是一个 action 对象形式, 则会忽略 case reducer 中配置的 prepare 方法

```jsx
const counterSlice = createSlice({
  name: 'counter',
  initialState: {
    count: 0,
  },
  reducers: {
    incrementByAmount:{
      reducer(state, action){
        state.count += action.payload;
      },
      prepare(val){
        return {payload: val + 2};
      }
    }
  }
});
dispatch(incrementByAmount(3));
// action 对象方式提交会忽略 case redcuer 的 prepare 方法
dispatch({type: 'counter/incrementByAmount', payload: 1});
```

#### combineSlices

合并多个 slice 为一个 reducer, 并允许初始化后更多的 reducer 注入

返回值

- withLazyLoadedSlices() 向 state 添加声明的 slice
- inject(slice, options) 添加 slice
  - options.overrideExisting 布尔值, 标识是否替换已存在的 slice
- selector() 将 reducer 包装在代理中以确保在当前状态未定义的情况下都能恢复到其初始状态

```jsx
import {combineSlices} from '@reduxjs/toolkit';

const lazySlice = createSlice({
  name: 'counter',
  initialState: {value: 0}
});

const rootReducer = combineSlices(staticSlice, userSlice);
const injectReducer = rootReducer.inject(lazySlice);
// OR
const injectSlice = lazySlice.injectInfo(rootReducer);

const selectCounterValue = (rootState) => rootState.counter?.value // number | undefined
const wrappedSelectCounterValue = injectReducer.selector((rootState) => rooState.counter.value);
console.log(
  selectCounterValue({}), // undefined
  selectCounterValue({counter: {value: 2}}), // 2
  wrappedSelectCounterValue({}), // 0
  wrappedSelectCounterValue({counter: {value: 2}}), // 2
)
```

#### createAsyncThunk <em id="createAsyncThunk"></em> <!--markdownlint-disable-line-->

接收一个 [actionCreator](#createAction)和一个回调函数并返回一个 Promise, 同时会创建三个 actionCreator 分别对应 pending, fulfilled, rejected 的状态, 不会生成 reducer

- type actionCreator, 如 `users/requestStatus` 将被创建为
  - pending: `users/requestStatus/pending`
  - fulfilled: `users/requestStatus/fulfilled`
  - rejected: `users/requestStatus/rejected`
- payloadCreator 函数, 将返回一个 promise, 接收两个参数
  - arg 包含了 thunk actionCreator 被 dispatch 时传入的参数, `dispatch(fetchUsers({status: 'active'}))`
  - thunkApi 包含了 thunk 函数的所有参数
    - dispatch()  Redux 的 dispatch 方法
    - getState()  Redux 的 getState 方法
    - extra 传递给 thunk 中间件的参数
    - requestId 自动生成的标识当前请求的唯一 id
    - signal 信号, AbortController.signal
    - rejectWithValue(value, [meta]) 修改当前 promise 的状态为 rejected
    - fulfilledWithValue(value, [meta])  修改当前 promise 的状态为 fulfilled
- options
  - condition(arg,{getState, extra}): boolean | Promise\<boolean\> 用来跳过执行 payloadCreator 和 所有的 dispatch
  - dispatchConditionRejection 布尔值, 如果 condition() 返回 false 所有的 action 都不会 dispatch, 如果想要当 thunk 结束 action 的状态标记为 rejected, 则设置为 true
  - idGenerator(arg): string 默认的 requestId 由 nanoid() 生成, 自定义生成 id 逻辑
  - serializeError(error: unknown) => any 替换内部的 `miniSerializeError` 方法
  - getPendingMeta({arg, requestId}, {getState, extra}): any 创建对象和 `pendingAction.meta` 合并

返回值

- thunk 函数, 带有 3 个状态
  - pending
  - fulfilled
  - rejected

```jsx
import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';

const promise = createAsyncThunk(type, payloadCreator, options?);

const fetchUserById = createAsyncThunk(
  'users/fetchUserById', 
  async (userId: number, {dispatch, requestId, getState, fulfilledWithValue, rejectWithValue}) => {
    try{
      const response = await fetch(userId);
      return response.data;
    }catch(err){
      return rejectWithValue(err.response.data);
    }
  }, {
    condition(userId, {getState, extra}){
      const {users} = getState();
      const fetchStatus = users.requests[userId];
      if(fetchStatus === 'fulfilled' || fetchStatus === 'loading'){
        // Already fetched or in progress, don't need to re-fetch
        return false;
      }
    }
  }
);
const usersSlice = createSlice({
  name: 'users',
  initialState: { },
  reducers:{},
  // 处理 asyncThunk 状态的 reducer
  extraReducers(builder) {
    builder.addCase(fetchUserById.pending, (state, action) => {
      state.status = 'loading';
    }).addCase(fetchUserById.fulfilled, (state, action) => {
      state.status = 'fulfilled';
      state.user = action.payload;
    }).addCase(fetchUserById.rejected, (state, action) => {
      state.status = 'rejected';
    });
  }
});

dispatch(fetchUserById(123));
```

#### createEntityAdapter

生成一组预构建的 reducer 和 seletors, 用于对包含特定类型数据对象实例的规范化状态结构执行 CRUD 操作, 这些 reducer 函数可以作为 case reducer 传递给 [createReducer](#createReducer) 和 [createSlice](#createSlice), 也可以作为 `createReducer` 和 `createSlice` 的辅助函数

- selectId 可选, 函数, 接收一个 entity 实例并返回一个唯一 id, 如果未提供则默认为 entity => entity.id
- sortComparer 可选, 函数, 接收两个 entity 实例, 返回一个标准的 `Array.sort()` 排序之后的结果 (1, 0, -1) 以指示它们的排序相对排序, 如果未提供将不会排序, 也不会保证排序

- getInitialState() 如果传入对象参数, 将被合并到 initialState 中并返回
- getSelectors() 生成一组标准的 selector 函数

- addOne/addMany 向 state 添加 items
- setOne/setMany 添加新 items 或替换现有 items
- setAll 替换所有 items
- removeOne/removeMany 根据 ID 删除 items
- removeAll 移除所有 items
- updateOne/updateMany 通过提供部分值更新现有 items
- upsertOne/upsertMany 添加新 items 或更新现有 items

```jsx
import {createSlice, createAsyncThunk, createEntityAdapter} from '@reduxjs/toolkit';

const todosAdapter = createEntityAdapter({
  selectId: todo => todo.id,
  sortComparer: (a, b) => a.id < b.id
});
const initialState = todosAdapter.getInitialState({loading: 'idle'});

// Thunk 函数
const fetchTodos = createAsyncThunk("todos/fetchTodos", async () => {
  const response = await client.get("/fakeApi/todos");
  return response.todos;
});
const saveNewTodo = createAsyncThunk("todos/saveNewTodo",
  async (text) => {
    const initialTodo = { text };
    const response = await client.post("/fakeApi/todos", { todo: initialTodo });
    return response.todo;
  }
);

const todosSlice = createSlice({
  name: 'todos',
  initialState,
  reducers: {
    todoDeleted: todosAdapter.removeOne, // 根据 id 删除 todo
    completeTodosCleard(state, action) {
      const completedIds = Object.values(state.entities)
        .filter(todo => todo.complete)
        .map(todo => todo.id);
      // 删除所有已完成的 todo
      todosAdapter.removeMany(state, completedIds);
    }
  },
  extraReducers(builder){
    builder.addCase(fetchTodos.pending, (state, action) => {
      state.status = 'loading';
    }).addCase(fetchTodos.fulfilled, (state, action) => {
      state.status = 'idle';
    }).addCase(saveNewTodo.fulfilled, todosAdapter.addOne)
  }
})
```

#### createSelector <em id="createSelector"></em> <!--markdownlint-disable-line-->

函数组件每次重新渲染都会重新执行 selector, createSelector 用于创建带有**记忆化**的 selector, 当给定的 inputSelector 没有发生变化时返回已缓存的 selector

- inputSelectors 创建记忆化 selector 的依赖, 可以是一个函数, 也可以是多个函数组成的数组, 返回值依次作为 resultFn 的参数传入
  - selectorFn 接收 state 作为第一个参数和剩余的参数并返回指定结果
- resultFn 在 inputSelectors 之后调用并依次接收来自 inputSelectors 函数的返回值作为参数并返回结果

返回值, 带有记忆化的函数

```jsx
import {createSelector} from 'reselect';
import {useSelector} from 'react-redux';

const selectTodos = state => state.todos;
const selectTodosStatus = (state, completed) => completed;
const memoizedSelectTodoCount = createSelector([selectTodos, selectTodosStatus], (todos, completed) => {
  return todos.filter(todo => todo.completed === completed).length;
});

function CompletedTodosCount({completed}){
  const matchingCount = useSelector((state) => memoizedSelectTodoCount(state, complete));
  return <div>{matchingCount}</div>
}
function App(){
  return (
    <>
      <span>Number of done todos</span>
      <CompletedTodosCount completed={true}/>
    </>
  )
}
```

#### nanoid

生成一个非加密安全的字符串 id, 通常被用作 [createAsyncThunk](#createAsyncThunk) 的 request IDs.

```jsx
import {nanoid} from '@reduxjs/toolkit';
console.log(nanoid()); // 'dgPXxUz_6fWIQBD8XmiSy'
```

#### miniSerializeError

createAsyncThunk 默认的错误序列化函数

#### copyWithStructuralSharing

递归的将两个相似的对象合并在一起, 如果值看起来相同, 则保留现有的引用. 这在内部用于帮助确保重新获取的数据继续使用相同的引用,
除非新数据实际发生了变化, 以避免不必要的重新呈现. 否则每次重新获取都可能导致整个数据集被替换, 所有消费组件总是重新渲染

#### @reduxjs/toolkit/query <em id="RTK-Query"></em> <!--markdownlint-disable-line-->

独立可选的入口, 允许定义端点(REST, GraphQL或任何异步函数)并生成 reducer 和中间件来完整管理数据获取, 加载状态更新和结果缓存, 还可以自动生成 React Hooks, 可用于组件获取数据

## react-redux

### Provider

- store
- serverState 使用 SSR 时传递的 prop
- context
- stabilityCheck
- children

```jsx
import { Provider } from 'react-redux';
import { createRoot } from 'react-dom/client';
createRoot(document.getElementById('root')).render(
  <Provider store={store}>
    {/*  */}
  </Provider>
)
```

```tsx
import { hydrateRoot } from 'react-dom/client';
import { Provider} from 'react-redux';
import { configureStore } from '@reduxjs/toolkit';

const preloadState = window.__PRELOAD_STATE__;

const store = configureStore({
  reducers: {},
  preloadState
});

hydrateRoot(document.getElementById('root'), 
  <Provider store={store} serverState={preloadState}>
    <App/>
  </Provider>
);
```

### shallowEqual

### useSelector

使用 selector 函数从 Redux store 中提取数据用于当前组件

使用 [createSelector](#createSelector) 创建记忆化的 selector

- selector
- equalityFn

```jsx
import {useSelector, shallowEqual} from 'react-redux';

const selectedData = useSelector(selectorReturningObject, shallowEqual);
// OR
const selectedData = useSelector(selectorReturningObject, {equalityFn: shallowEqual});

function TodoListItem(props){
  const todo = useSelector(state => state.todos[props.id]);
  return <div>{todo.text}</div>
}
```

### useDispatch

```jsx
import {useCallback, memo} from 'react';
import {useDispatch} from 'react-redux';

function CounterComponent(){
  const dispatch = useDispatch();
  const incrementCounter = useCallback(() => {
    dispatch({type:'increment-counter'});
  },[dispatch]);
  return (
    <div>
      <span>CounterComponent</span>
      <MyIncrement onIncrement={incrementCounter}/>
    </div>
  )
}
const MyIncrement = memo(({onIncrement}) => {
  return (<button onClick={onIncrement}>increment counter</button>)
})
```

### useStore

大多数情况使用 `useSelector`

```jsx
import {useStore} from 'react-redux';

function MyComponent(){
  const store = useStore();

  return <div>{store.getState().todos.length}</div>;
}
```
