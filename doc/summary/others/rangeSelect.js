process.env.NODE_ENV === 'development' && console.log(
  `/**
  * @directive v-range-select.constant.stop="DOMSelector"
  * @modifiers constant          // 非必传, 是否持久高亮选中项, 默认选择结束时移除所有选中项的高亮
  * @modifiers stop              // 非必传, 是否阻止事件冒泡
  * @value DOMSelector           // 非必传, 支持的 DOM 选择器, 作为过滤选中项的条件, 如果没有则表示所有 Element 都支持选中
  * @attr range-active-style=""  // 非必传, 选中项的样式, 默认红色描边,黑色阴影
  * @attr range-style=""         // 非必传, 选区框的样式, 默认蓝色虚线边框
  * @event range-select-start    // 非必传, 选择开始时触发
  * @event range-select-end      // 非必传, 选择结束时触发
  * @event range-select          // 选择改变时触发并返回选中的项
  */`
)
const MOUSEDOWN = 'mousedown'
const MOUSEUP = 'mouseup'
const MOUSEMOVE = 'mousemove'
const genUUID = () => Math.random().toString(36).substring(2)
// 切换状态
const changeItemStatus = function(el, selectedElements = []) {
  const { modifiers, selectedElementsLen, uuid, evtType } = el.conf
  const len = selectedElements.length
  if ((evtType === MOUSEUP || selectedElementsLen !== len) && !modifiers?.constant) {
    el.querySelectorAll('[data-range-active="' + uuid + '"]').forEach(item => item.removeAttribute('data-range-active'))
  }
  if (len !== 0 && selectedElementsLen !== len) {
    selectedElements.forEach(item => item.element.setAttribute('data-range-active', uuid))
  }
}
// 更新选区
const updateSelectionAndHighlight = function(x, y) {
  // console.log('updateSelectionAndHighlight', this)
  const { selectionBox, startX, startY, el, value, uuid } = this
  const width = Math.abs(x - startX)
  const height = Math.abs(y - startY)

  selectionBox.style.left = Math.min(startX, x) + 'px'
  selectionBox.style.top = Math.min(startY, y) + 'px'
  selectionBox.style.width = width + 'px'
  selectionBox.style.height = height + 'px'

  const selRect = selectionBox.getBoundingClientRect()
  const selectedElements = Array.from(el.querySelectorAll(value || '*')).map(element => ({ element, rect: element.getBoundingClientRect() })).filter(({ element, rect }) => {
    return element.nodeName !== 'STYLE' && element.dataset.range !== uuid && rect.right <= selRect.right && rect.left >= selRect.left && rect.bottom <= selRect.bottom && rect.top >= selRect.top
  })
  // 切换状态
  changeItemStatus(el, selectedElements)

  return selectedElements
}
// handle mousedown mouseup mousemove
const handleElMouseEvt = function(evt) {
  // console.log('hanledElMouseEvt', this, evt)
  const { selectionBox, isSelecting, modifiers, selectedElementsLen, el } = this
  const that = this
  if (modifiers?.stop) {
    evt.stopPropagation()
  }
  const containerRect = el.getBoundingClientRect()
  const x = evt.clientX - containerRect.left
  const y = evt.clientY - containerRect.top
  that.evtType = evt.type

  if (evt.type === MOUSEDOWN && evt.button === 0) {
    that.startX = x
    that.startY = y
    that.isSelecting = true
    selectionBox.style.display = 'block'
    updateSelectionAndHighlight.call(that, x, y)
    el.dispatchEvent(new CustomEvent('range-select-start'))
  } else if (evt.type === MOUSEUP && isSelecting && evt.button === 0) {
    that.isSelecting = false
    that.selectedElementsLen = 0
    selectionBox.style.display = 'none'
    // 切换状态
    changeItemStatus(el, [])
    el.dispatchEvent(new CustomEvent('range-select-end'))
  } else if (evt.type === MOUSEMOVE && isSelecting) {
    const selectedElements = updateSelectionAndHighlight.call(that, x, y)
    const len = selectedElements.length
    if (selectedElementsLen !== len) {
      that.selectedElementsLen = len
      el.dispatchEvent(new CustomEvent('range-select', { detail: selectedElements.map(item => item.element) }))
    }
  }
}
// 初始化配置
const initialRangeSelect = function() {
  const uuid = genUUID()
  const div = document.createElement('div')
  div.dataset.range = uuid
  div.style.display = 'none'

  return { uuid, selectionBox: div, isSelecting: false, startX: 0, startY: 0 }
}

const bindHook = function(el, binding, vnode, prevVnode) {
  console.log('range-select bind', el, binding, vnode)
  if (binding.value && typeof binding.value !== 'string') {
    throw new Error('v-range-select accepts a dom selector as its value.')
  }

  const conf = initialRangeSelect()
  let cls = el.classList.value.trim().replace(/\s/g, '.')
  if (cls.length !== 0) {
    cls = '.' + cls
  } else {
    cls = el.nodeName.toLowerCase() + '[data-range-container="' + conf.uuid + '"]'
    el.dataset.rangeContainer = conf.uuid
  }

  const handler = handleElMouseEvt.bind(conf)
  const rangeActiveStyle = el.getAttribute('range-active-style') || 'box-shadow:0 0 10px rgba(0, 0, 0, 0.5);outline:2px solid red;'
  let rangeStyle = el.getAttribute('range-style')

  rangeStyle = rangeStyle ? rangeStyle.trim().endsWith(';')
    ? rangeStyle + 'position:absolute;' : rangeStyle + ';position:absolute;'
    : 'border:1px dashed #1a44e9;pointer-events:none;position:absolute;'
  const style = `<style data-range-style="${conf.uuid}">
    ${cls} {position:relative;}
    ${cls} *::selection {color:initial;background:initial;box-shadow:initial;outline:initial;}
    ${cls} [data-range="${conf.uuid}"] {${rangeStyle}}
    ${cls} [data-range-active="${conf.uuid}"] {${rangeActiveStyle}}
  </style>`
  conf.value = binding.value
  conf.modifiers = binding.modifiers
  conf.el = el
  el.conf = conf
  el.__handleElMouseEvt__ = handler
  el.insertAdjacentHTML('afterbegin', style)
  el.appendChild(conf.selectionBox)

  el.addEventListener('mousedown', handler)
  el.addEventListener('mouseup', handler)
  el.addEventListener('mousemove', handler)
}
const unbindHook = function(el, binding, vnode, prevVnode) {
  // console.log('range-select unbind')
  el.querySelectorAll('style[data-range-style="' + el.conf.uuid + '"],[data-range="' + el.conf.uuid + '"]').forEach(item => item.remove())
  el.querySelectorAll('[data-range-active="' + el.conf.uuid + '"]').forEach(item => item.removeAttribute('data-range-active'))
  el.removeAttribute('data-range-container')
  const handler = el.__handleElMouseEvt__
  if (handler) {
    el.removeEventListener('mousedown', handler)
    el.removeEventListener('mouseup', handler)
    el.removeEventListener('mousemove', handler)
    delete el.conf.el
    delete el.conf.modifiers
    delete el.__handleElMouseEvt__
    delete el.conf
  }
}

export default {
  bind(el, binding, vnode, prevVnode) {
    bindHook(el, binding, vnode, prevVnode)
  },
  unbind(el, binding, vnode, prevVnode) {
    unbindHook(el, binding, vnode, prevVnode)
  },
  created(el, binding, vnode, prevVnode) {
    bindHook(el, binding, vnode, prevVnode)
  },
  beforeUnmount(el, binding, vnode, prevVnode) {
    unbindHook(el, binding, vnode, prevVnode)
  }
  /**
   * 自定义指令钩子:
   * created 在绑定元素的 attribute 或事件监听器被应用之前调用
   * beforeMount 在绑定元素的父组件挂载之前调用
   * mounted 绑定元素的父组件被挂载时调用
   * beforeUpdate 在包含组件的 VNode 更新之前调用
   * updated 在包含组件的 VNode 及其子组件的 VNode 更新之后调用
   * beforeUnmount 在绑定元素的父组件卸载之前调用
   * unmounted 卸载绑定元素的父组件时调用
   */
}
