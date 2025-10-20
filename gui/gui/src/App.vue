<script setup lang="ts">
import { ref } from 'vue'

const searchValue = ref('')
const loading = ref(false)
const result = ref<any | null>(null) // 支持对象
const error = ref<string | null>(null)

function validateInput(val: string) {
  // 只允许数字和TJC字母，长度不超过6
  const valid = /^[0-9TJCYtjcy]{0,6}$/.test(val)
  return valid
}

function onInput(e: Event) {
  const val = (e.target as HTMLInputElement).value
  // 只保留数字和TJC字母，最大6位
  const filtered = val.replace(/[^0-9TJCtjc]/g, '').slice(0, 6)
  searchValue.value = filtered
  if (!validateInput(filtered)) {
    error.value = '只能输入最多6位数字或T/J/C字母，不能包含特殊字符'
  } else {
    error.value = null
  }
}

async function handleSearch() {
  if (!searchValue.value.trim()) {
    error.value = '请输入车厢号'
    result.value = null
    return
  }
  error.value = null
  result.value = null
  loading.value = true
  try {
    // 发送请求处
    const res = await fetch('http://127.0.0.1:9987/api/calculate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        // 'cookie': cookies, // 如需cookie可补充
      },
      body: JSON.stringify({ carriage_number: searchValue.value.trim() }),
    })
    if (!res.ok) throw new Error('查询失败')
    const data = await res.json()
    // 支持对象格式渲染
    result.value = data.result || '未查询到相关信息'
  } catch (e: any) {
    error.value = e.message || '查询出错'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container">
    <h1>上海地铁车厢号数据库</h1>
    <div class="search-bar">
      <input
        v-model="searchValue"
        type="text"
        maxlength="6"
        placeholder="请输入车厢号"
        @input="onInput"
        @keyup.enter="handleSearch"
        autocomplete="off"
      />
      <button :disabled="loading" @click="handleSearch">
        Go!
      </button>
    </div>
    <div class="result" v-if="result || loading">
      <span>查询结果：</span>
      <span class="result-text">
        <template v-if="loading">查询中...</template>
        <template v-else>
          <template v-if="typeof result === 'object' && result">
            <div v-if="result.isInputCarriageTypeCorrect === false">
              输入车厢号有误，正确应该为 {{ Array.isArray(result.Carriage_num) ? result.Carriage_num[result.Carriage_index] : result.Carriage_num }}
            </div>
            <div>车号：{{ result.TrainId }}</div>
            <div>车型：{{ result.Train_type }}</div>
            <div>关于该列车：{{ result.Train_detail || '暂无信息' }}</div>
          </template>
          <template v-else>
            {{ result }}
          </template>
        </template>
      </span>
    </div>
    <div class="error" v-if="error">{{ error }}</div>
  </div>
</template>
<style scoped>
.container {
  max-width: 420px;
  margin: 60px auto;
  padding: 32px 24px;
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 4px 24px rgba(0,0,0,0.08);
  text-align: center;
  font-family: 'Segoe UI', 'PingFang SC', Arial, sans-serif;
}
h1 {
  font-size: 1.4rem;
  margin-bottom: 32px;
  color: #222;
  letter-spacing: 1px;
}
.search-bar {
  display: flex;
  gap: 12px;
  justify-content: center;
  margin-bottom: 24px;
}
input[type="text"] {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #d0d0d0;
  border-radius: 6px;
  font-size: 1rem;
  transition: border 0.2s;
}
input[type="text"]:focus {
  border-color: #409eff;
  outline: none;
}
button {
  padding: 8px 20px;
  background: #409eff;
  color: #fff;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: 
    background 0.2s,
    box-shadow 0.2s,
    transform 0.28s cubic-bezier(.34,1.56,.64,1),
    filter 0.2s;
}
button:disabled {
  background: #bcdcff;
  cursor: not-allowed;
  box-shadow: none;
  transform: none;
}
button:not(:disabled):hover {
  box-shadow: 0 4px 18px rgba(64,158,255,0.18);
  transform: scale(1.1);
  filter: brightness(1.05);
}
button:not(:disabled):active {
  animation: bounce-press 0.6s cubic-bezier(.34,1.56,.64,1);
}
@keyframes bounce-press {
  0% {
    transform: scale(1.1);
  }
  25% {
    transform: scale(0.9);
  }
  60% {
    transform: scale(1.15);
  }
  100% {
    transform: scale(1.0);
  }
}
.result {
  margin-top: 18px;
  color: #222;
  font-size: 1.05rem;
}
.result-text {
  font-weight: bold;
  color: #409eff;
  margin-left: 6px;
}
.error {
  margin-top: 18px;
  color: #e74c3c;
  font-size: 0.98rem;
}

/* 响应式：手机屏幕下按钮和输入框纵向排列，按钮最大宽度不超过屏幕20% */
@media (max-width: 600px) {
  .container {
    padding: 18px 4vw;
    margin: 24px auto;
  }
  .search-bar {
    flex-direction: column;
    gap: 10px;
    align-items: stretch;
  }
  input[type="text"] {
    width: 100%;
    box-sizing: border-box;
  }
  button {
    width: 20vw;
    max-width: 100px;
    min-width: 60px;
    align-self: flex-end;
    box-sizing: border-box;
    margin-left: auto;
    margin-right: auto;
  }
}
</style>